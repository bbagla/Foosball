package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//circle is a struct containing centre co-ordinates of a circle and radius.
type circle struct {
	x, y, radius float64
}

var gameStatus = GameStatus{
	Teams: make([]team, 2),
	Ball:  Position{},
	Score: make([]int, 2),
}

//last_motion is a variable indicating the key that was pressed last(up/down)
//It is updated at each frame of the application.
var lastMotion int32

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
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
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()
	//tableTex is the variable containing the texture for the background(foosball table).
	var tableTex *sdl.Texture
	tableTex = drawBackground(tableTex, renderer)
	defer tableTex.Destroy()
	team1, err := newteam(renderer, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	team2, err := newteam(renderer, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	ball, err := newBall(renderer, boxWidth/2, boxHeight/2)
	if err != nil {
		fmt.Println(err)
		return
	}
	var lastStick = team1.mid[0:5]
	lastMotion = 0
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.Copy(tableTex, nil, nil)

		team1.draw(renderer)
		team2.draw(renderer)
		ball.draw(renderer)
		ball.update()
		ball.CheckCollision(team1, 1)
		ball.CheckCollision(team2, 2)
		ball.update()
		lastStick, lastMotion = team1.update(lastStick, 0)
		renderer.Copy(scoreUpdate(renderer), nil,
			&sdl.Rect{120, 0, 400, 25})
		renderer.Present()
		sdl.Delay(16)
	}
}
