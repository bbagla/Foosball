package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

//boxWidth is the width of the window.
//boxHeight is the height of the window.
const (
	boxWidth  = 648
	boxHeight = 498
)

// drawBackground makes the texture for the background.

func drawBackground(tableTex *sdl.Texture, renderer *sdl.Renderer) *sdl.Texture {
	backImg, err := sdl.LoadBMP("table.bmp")
	if err != nil {
		fmt.Println(err)
	}
	defer backImg.Free()
	tableTex, err = renderer.CreateTextureFromSurface(backImg)
	if err != nil {
		fmt.Println(err)
	}
	return tableTex
}
