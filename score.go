package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"strconv"
)

//updates scores on the top bar using sdl_ttf binding
func scoreUpdate(renderer *sdl.Renderer) *sdl.Texture {
	font, err := ttf.OpenFont("ZCOOLQingKeHuangYou-Regular.ttf", 32)
	if err != nil {
		fmt.Println(err)
	}
	s := "PLAYER 1 = " + strconv.Itoa(gameStatus.Score[0]) + "                                  PLAYER 2 = " + strconv.Itoa(gameStatus.Score[1])
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
