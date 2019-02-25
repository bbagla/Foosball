package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const(
 	boxWidth = 800
 	boxHeight = 600
)
type color struct{
	redColor,greenColor,blueColor byte
}

type position struct{
	x,y float32
}

type player struct {
	position
	radius int
	color
}


type ball struct {
	position
	radius int
	color
	xVelocity,yVelocity float32
}

func(player *player) draw(pixels []byte) {
	for y := -player.radius; y < player.radius; y++ {
		for x := -player.radius; x < player.radius; x++ {
			if x*x+y*y < player.radius*player.radius{
				setPixel(int(player.x)+x, int(player.y)+y, player.color, pixels)
			}
		}
	}
}

func(ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius{
				setPixel(int(ball.x)+x, int(ball.y)+y, ball.color, pixels)
			}
		}
	}
}

func setPixel(x,y int, colorsss color, pixels []byte){
	index:= (y*(boxWidth)+x)*4
	if index < len(pixels) - 4 && index >=0 {
		pixels[index] = colorsss.redColor
		pixels[index+1] = colorsss.greenColor
		pixels[index+2] = colorsss.blueColor
	}
}

func main() {
	window,err := sdl.CreateWindow("sdl2",sdl.WINDOWPOS_UNDEFINED,sdl.WINDOWPOS_UNDEFINED,int32(boxWidth),int32(boxHeight),sdl.WINDOW_SHOWN)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer window.Destroy()
	renderer,err := sdl.CreateRenderer(window,-1,sdl.RENDERER_ACCELERATED)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex,err :=renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888,sdl.TEXTUREACCESS_STREAMING,800,600)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	pixels :=make([]byte,800*600*4)

	player1:= player{position{100,100},20,color{255,0,0}}
	player2:= player{position{600,600},20,color{0,0,255}}
	ball:= ball{position{600,300},40,color{255,255,255},0,0}

	for{
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent(){
			switch event.(type){
			case *sdl.QuitEvent:
				return
			}
		}
		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)
		tex.Update(nil,pixels,3200)
		renderer.Copy(tex,nil,nil)
		renderer.Present()
		sdl.Delay(16)
	}
}
