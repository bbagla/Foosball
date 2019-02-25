package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const(
	boxWidth = 800
	boxHeight = 600
)
func drawBackground(tableTex *sdl.Texture,renderer *sdl.Renderer) *sdl.Texture{
	backImg,err:= sdl.LoadBMP("/home/bhavya/go/src/sdl/table.bmp")
	if err!=nil{
		fmt.Println(err)
	}
	tableTex,err = renderer.CreateTextureFromSurface(backImg)
	if err!=nil{
		fmt.Println(err)
	}
	return  tableTex
}
