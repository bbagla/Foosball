package serve

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

//BallSpeedX : Speed of ball in X-direcrion
//BallSpeedY : Speed of ball in Y-direcrion
const (
	boxWidth  = 648
	boxHeight = 498

	radius     = 8
	BallSpeedX = 1.5
	BallSpeedY = 0.5

	playerSpeed   = 5
	playerWidth   = 26
	playerHeight  = 30
	boundarywidth = 29
)

var (
	gameLogStd = log.New(os.Stdout, "[game] ", log.Ldate|log.Ltime)
	gameLogErr = log.New(os.Stderr, "ERROR [game] ", log.Ldate|log.Ltime)

	gameStatus = GameStatus{
		Teams:    make([]team, 2),
		Ball:     ball{},
		Score:    make([]int, 2),
		Renderer: nil,
	}
	ticking = false
)

func resetGame() {
	gameStatus.Score[0] = 0
	gameStatus.Score[1] = 0
	PrepareGame()
}
func PrepareGame() {
	ticking = false
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println(err)
		return
	}
	window, err := sdl.CreateWindow("sdl2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(boxWidth), int32(boxHeight), sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()
	gameStatus.Renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gameStatus.Renderer.Destroy()

	var tableTex *sdl.Texture
	tableTex = drawBackground(tableTex, gameStatus.Renderer)
	defer tableTex.Destroy()

	gameStatus.Teams[0], err = newteam(gameStatus.Renderer, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	gameStatus.Teams[1], err = newteam(gameStatus.Renderer, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	gameStatus.Ball, err = newBall(gameStatus.Renderer, boxWidth/2, boxHeight/2)
	if err != nil {
		fmt.Println(err)
		return
	}
	gameStatus.Teams[0].LastStick = gameStatus.Teams[0].Mid[0:5]
	gameStatus.Teams[1].LastStick = gameStatus.Teams[1].Mid[0:5]

	sendGameStatus()
	startGame()
}

func startGame() {
	if !ticking {
		ticking = true
		go gameLoop()
	}
}

func gameLoop() {
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		if !ticking {
			return
		}
		gameStatus.Ball.CheckCollision(gameStatus.Teams[0], 1)
		gameStatus.Ball.CheckCollision(gameStatus.Teams[1], 2)
		gameStatus.Ball.update()
		if channel == 0 {
			gameStatus.Teams[0].update()
		}
		if channel == 1 {
			gameStatus.Teams[1].update()
		}
		sendGameStatus()
		if gameStatus.Score[0] == 5 || gameStatus.Score[1] == 5 {
			resetGame()
		}
	}
}
