package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 5
	playerWidth = 20
	playerHeight = 30
)
type player struct{
	tex *sdl.Texture
	x,y float64
}

func newplayer(renderer *sdl.Renderer) (p player, err error){
		playerImg,err:= sdl.LoadBMP("/home/bhavya/go/src/sdl/player.bmp")
		if err!=nil{
			fmt.Println(err)
			return player{},fmt.Errorf("%v",err)
		}
		defer playerImg.Free()
		p.tex,err = renderer.CreateTextureFromSurface(playerImg)
		if err!=nil{
			fmt.Println(err)
			return player{},fmt.Errorf("%v",err)
		}

		return  p, nil
}

func (p *player) draw (renderer *sdl.Renderer){
	renderer.Copy(p.tex,
		&sdl.Rect{0,0,100,150},
		&sdl.Rect{int32(p.x),int32(p.y),playerWidth,playerHeight})
}

func (p *player) update(){
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_UP]==1{
		if p.y>0 {
			p.y -= playerSpeed
		}else{
			p.y=0
		}
	}else if keys[sdl.SCANCODE_DOWN] == 1{
		if p.y <boxHeight - playerHeight -1{
			p.y += playerSpeed
		}else{
			p.y = boxHeight - playerHeight - 1
		}
	}
}










































