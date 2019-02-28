package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type circle struct {
	x, y, radius float64
}

var gameStatus = GameStatus{
	Teams: make([]team, 2),
	Ball:  Position{},
	Score: make([]int, 2),
}

var last_motion int32

func main() {
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
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

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
	var last_stick = team1.mid[0:5]
	last_motion = 0
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
		last_stick, last_motion = team1.update(last_stick, 0)
		if insideGoal == true {
			ball.movementInsidePost()
		}
		last_stick = team1.update(last_stick)
		renderer.Present()
		sdl.Delay(16)
	}
}
