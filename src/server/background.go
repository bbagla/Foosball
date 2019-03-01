package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

// drawBackground makes the texture
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
