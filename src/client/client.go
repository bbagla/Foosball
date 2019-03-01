package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var server = "localhost:80"

var keyboardInput = KeyboardInput{
	SelectStick: 0,
	KeyPressed:  0,
}
var gameStatus = GameStatus{
	Teams: make([]team, 2),
	Ball:  ball{},
	Score: make([]int, 2),
}

//Returns a Texture for the player image based on teamId.
func setPlayerImage(renderer *sdl.Renderer, teamID int32) *sdl.Texture {
	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	playerImg, err := img.Load("Player_Red.png")
	if teamID == 2 {
		playerImg, err = img.Load("Player_Blue.png")
	}
	if err != nil {
		fmt.Println(err)
	}
	defer playerImg.Free()
	Tex, err := renderer.CreateTextureFromSurface(playerImg)
	if err != nil {
		fmt.Println(err)
	}
	return Tex
}

//Associates the textures to their corresponding players from GameStatus.
func setTeam(t *team, teamID int32, renderer *sdl.Renderer) {
	t.GoalKeeper[0].Tex = setPlayerImage(renderer, teamID)
	for i := range t.Defence {
		t.Defence[i].Tex = setPlayerImage(renderer, teamID)
	}
	for i := range t.Mid {
		t.Mid[i].Tex = setPlayerImage(renderer, teamID)
	}
	for i := range t.Attack {
		t.Attack[i].Tex = setPlayerImage(renderer, teamID)
	}
}

//Returns a Texture for the ball image.
func setBallImage(renderer *sdl.Renderer) *sdl.Texture {

	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	BallImg, err := img.Load("Ball.png")
	if err != nil {
		fmt.Println(err)
	}
	defer BallImg.Free()
	Tex, err := renderer.CreateTextureFromSurface(BallImg)
	if err != nil {
		fmt.Println(err)
	}
	return Tex
}

func main() {
	flag.Parse()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Create log file and logger.
	logFile, err := os.Create("client.log")
	if err != nil {
		fmt.Println("Failed to create client.log")
		return
	}
	defer logFile.Close()
	log := log.New(logFile, "", log.Lmicroseconds)

	url := url.URL{
		Scheme: "ws",
		Host:   server,
		Path:   "/socket/0",
	}
	log.Printf("Making connection to: %s", url.String())

	conn, resp, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Println("Dial Error: ", err)
		log.Printf("handshake failed with status %d", resp.StatusCode)

	}
	defer conn.Close()

	done := make(chan struct{})

	// This Goroutine is our read/write loop. It keeps going until it cannot use the WebSocket anymore.
	go func() {

		defer conn.Close()
		defer close(done)
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			fmt.Println(err)
			return
		}
		err = ttf.Init()
		if err != nil {
			fmt.Println(err)
			return
		}
		//sdl.CreateWindow creates a window for running the application.
		window, err := sdl.CreateWindow("sdl2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(boxWidth), int32(boxHeight), sdl.WINDOW_OPENGL)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer window.Destroy()
		//sdl.CreateRenderer creates a renderer for drawing on the window.
		Renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer Renderer.Destroy()

		//tableTex is the variable containing the texture for the background(foosball table).
		var tableTex *sdl.Texture
		tableTex = drawBackground(tableTex, Renderer)
		defer tableTex.Destroy()
		Renderer.Present()
		//Initiate teams
		setTeam(&gameStatus.Teams[0], 1, Renderer)
		setTeam(&gameStatus.Teams[1], 2, Renderer)
		//Associates the  ball to its corresponding ball from GameStatus.
		gameStatus.Ball.Tex = setBallImage(Renderer)
		for {

			err := conn.ReadJSON(&gameStatus)
			if err != nil {
				log.Println("Read Error: ", err)
				break
			}
			Renderer.Copy(tableTex, nil, nil)
			Renderer.Present()
			gameStatus.Ball.draw(Renderer)
			gameStatus.Teams[0].draw(Renderer)
			gameStatus.Teams[1].draw(Renderer)
			Renderer.Present()

			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					return
				//Check which key is pressed and update keyboard input
				case *sdl.KeyboardEvent:
					{
						switch t.Keysym.Sym {
						case sdl.K_a:
							keyboardInput.SelectStick = 1
						case sdl.K_s:
							keyboardInput.SelectStick = 2
						case sdl.K_d:
							keyboardInput.SelectStick = 3
						case sdl.K_f:
							keyboardInput.SelectStick = 4

						case sdl.K_UP:
							keyboardInput.KeyPressed = 1
						case sdl.K_DOWN:
							keyboardInput.KeyPressed = 2
						}
					}
				}

			}
			err = conn.WriteJSON(keyboardInput)
			keyboardInput.KeyPressed = 0
			if err != nil {
				log.Println("Write Error: ", err)
				break
			}

		}
	}()

	for {
		select {
		// Block until interrupted. Then send the close message to the server and wait for our other read/write Goroutine to signal 'done'
		case <-interrupt:
			log.Println("Client interrupted.")
			err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("WebSocket Close Error: ", err)
			}
			// Wait for 'done' or one second to pass.
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		// WebSocket has terminated before interrupt.
		case <-done:
			log.Println("WebSocket connection terminated.")
			return
		}
	}
}
