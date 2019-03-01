package main

import (
	client "concurrency-17/src/client"
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
	"strconv"
	"time"
)

var server = "localhost:8000"

var keyboardInput = client.KeyboardInput{
	SelectStick: 0,
	KeyPressed:  0,
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
func setTeam(t *client.Team, teamID int32, renderer *sdl.Renderer) {
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

func ScoreUpdate(renderer *sdl.Renderer) *sdl.Texture {
	font, err := ttf.OpenFont("ZCOOLQingKeHuangYou-Regular.ttf", 32)
	if err != nil {
		fmt.Println(err)
	}
	s := "PLAYER 1 = " + strconv.Itoa(client.GamesStatus.Score[0]) + "                                  PLAYER 2 = " + strconv.Itoa(client.GamesStatus.Score[1])
	fontsurface, err := font.RenderUTF8Solid(s, sdl.Color{0, 0, 0, 0})
	if err != nil {
		fmt.Println(err)
	}
	texture, err := renderer.CreateTextureFromSurface(fontsurface)
	if err != nil {
		fmt.Println(err)
	}
	return texture
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
	serverptr:=flag.String("server","localhost:8000","server string")
	chanptr:=flag.String("channel","0","server string")
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
		Host:   *serverptr,
		Path:   "/socket/"+*chanptr,
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
		window, err := sdl.CreateWindow("sdl2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(client.BoxWidth), int32(client.BoxHeight), sdl.WINDOW_OPENGL)
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
		tableTex = client.DrawBackground(tableTex, Renderer)
		defer tableTex.Destroy()
		Renderer.Present()
		//Initiate teams
		setTeam(&client.GamesStatus.Teams[0], 1, Renderer)
		setTeam(&client.GamesStatus.Teams[1], 2, Renderer)
		//Associates the  ball to its corresponding ball from GameStatus.
		client.GamesStatus.Ball.Tex = setBallImage(Renderer)
		for {

			err := conn.ReadJSON(&client.GamesStatus)
			if err != nil {
				log.Println("Read Error: ", err)
				break
			}
			Renderer.Copy(tableTex, nil, nil)
			Renderer.Copy(ScoreUpdate(Renderer), nil, &sdl.Rect{120, 0, 400, 25})
			Renderer.Present()
			client.GamesStatus.Ball.Draw(Renderer)
			client.GamesStatus.Teams[0].Draw(Renderer)
			client.GamesStatus.Teams[1].Draw(Renderer)
			Renderer.Present()
			//fmt.Println(client.GamesStatus.Score)
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
						case sdl.K_SPACE:
							keyboardInput.KeyPressed = 3	
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
