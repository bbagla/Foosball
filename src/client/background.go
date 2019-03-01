package client

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

//boxWidth is the width of the window.
//boxHeight is the height of the window.
const (
	BoxWidth  = 648
	BoxHeight = 498
)

// drawBackground makes the texture for the background.
func DrawBackground(tableTex *sdl.Texture, renderer *sdl.Renderer) *sdl.Texture {
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
