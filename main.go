package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

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
		team1.update()
		team2.update()
		renderer.Present()
		sdl.Delay(16)
	}
}
