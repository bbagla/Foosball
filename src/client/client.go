package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var server = "localhost:80"

var keyboardInput = KeyboardInput{
	// teamID:     0,
	SelectStick: 0,
	KeyPressed:  0,
}
var gameStatus = GameStatus{
	Teams: make([]team, 2),
	Ball:  ball{},
	Score: make([]int, 2),
	// LastStick:  make([]player, 4),
	// LastMotion: 0,
	// Renderer:   nil,
}

func setPlayerImage(renderer *sdl.Renderer, teamId int32) *sdl.Texture {
	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	playerImg, err := img.Load("Player_Red.png")
	if teamId == 2 {
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

func setTeam(t *team, teamId int32, renderer *sdl.Renderer) {
	t.GoalKeeper[0].Tex = setPlayerImage(renderer, teamId)
	for i := range t.Defence {
		t.Defence[i].Tex = setPlayerImage(renderer, teamId)
	}
	for i := range t.Mid {
		t.Mid[i].Tex = setPlayerImage(renderer, teamId)
	}
	for i := range t.Attack {
		t.Attack[i].Tex = setPlayerImage(renderer, teamId)
	}
}

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

		window, err := sdl.CreateWindow("sdl2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(boxWidth), int32(boxHeight), sdl.WINDOW_OPENGL)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer window.Destroy()
		Renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer Renderer.Destroy()

		var tableTex *sdl.Texture
		tableTex = drawBackground(tableTex, Renderer)
		defer tableTex.Destroy()
		Renderer.Present()

		setTeam(&gameStatus.Teams[0], 1, Renderer)
		setTeam(&gameStatus.Teams[1], 2, Renderer)
		gameStatus.Ball.Tex = setBallImage(Renderer)
		for {

			err := conn.ReadJSON(&gameStatus)
			if err != nil {
				log.Println("Read Error: ", err)
				break
			}

			// log.Println("Recieved :  ", gameStatus.Ball.X, gameStatus.Ball.Y)
			//RENDER STUFF HERE
			Renderer.Copy(tableTex, nil, nil)
			Renderer.Present()
			gameStatus.Ball.draw(Renderer)
			gameStatus.Teams[0].draw(Renderer)
			gameStatus.Teams[1].draw(Renderer)
			//Renderer=gameStatus.Renderer
			Renderer.Present()

			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					return

				case *sdl.KeyboardEvent:
					{
						switch t.Keysym.Sym {
						case sdl.K_a:
							fmt.Println("selected a")
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
					// fmt.Println("selected Down")
					// case *sdl.KEYUP:
					// 	fmt.Println("selected Up")
				}

			}
			//READ KEYBOARD INPUT
			// keys := sdl.GetKeyboardState()
			// fmt.Println("got keys")
			// if keys[sdl.SCANCODE_A] == 1 {
			// 	fmt.Println("selected A")
			// 	keyboardInput.SelectStick = 1
			// } else if keys[sdl.SCANCODE_S] == 1 {
			// 	fmt.Println("selected S")
			// 	keyboardInput.SelectStick = 2
			// } else if keys[sdl.SCANCODE_D] == 1 {
			// 	fmt.Println("selected D")
			// 	keyboardInput.SelectStick = 3
			// } else if keys[sdl.SCANCODE_F] == 1 {
			// 	fmt.Println("selected F")
			// 	keyboardInput.SelectStick = 4
			// }

			// if keys[sdl.SCANCODE_UP] == 1 {
			// 	keyboardInput.KeyPressed = 1
			// } else if keys[sdl.SCANCODE_DOWN] == 1 {
			// 	keyboardInput.KeyPressed = 2
			// } else if keys[sdl.SCANCODE_SPACE] == 1 {
			// 	keyboardInput.KeyPressed = 3
			// }

			// log.Println("Sending: input", keyboardInput)
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
		// Block until interrupted. Then send the close message to the server and wait for our other read/write Goroutine
		// to signal 'done'
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
