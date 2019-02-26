package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 5
	playerWidth = 20
	playerHeight = 30
)

type team struct {
	goalKeeper player
	leftBack player
	rightBack player
	centerMid player
	lmid player
	llmid player
	rmid player
	rrmid player
	striker player
	leftWing player
	rightWing player
}

type player struct{
	tex *sdl.Texture
	x,y float64
}

func newteam(renderer *sdl.Renderer) (p team, err error){
	p.goalKeeper,err = newplayer(renderer,61-playerWidth/2,boxHeight/2 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}

	p.leftBack,err = newplayer(renderer,136 - playerWidth/2,boxHeight/3 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}

	p.rightBack,err = newplayer(renderer,136 - playerWidth/2,boxHeight/1.5 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}


	p.llmid,err = newplayer(renderer,286 - playerWidth/2,boxHeight/6 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}

	p.lmid,err = newplayer(renderer,286 - playerWidth/2,boxHeight/3 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}

	p.centerMid,err = newplayer(renderer,286 - playerWidth/2,boxHeight/2 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}

	p.rmid,err = newplayer(renderer,286 - playerWidth/2,(boxHeight*2)/3 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}

	p.rrmid,err = newplayer(renderer,286 - playerWidth/2,(boxHeight*5)/6 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}

	p.leftWing,err = newplayer(renderer,436 - playerWidth/2,boxHeight/4 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}

	p.striker,err = newplayer(renderer,436 - playerWidth/2,boxHeight/2 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}


	p.rightWing,err = newplayer(renderer,436 - playerWidth/2,(boxHeight*3)/4 - playerHeight/2)
	if err!=nil{
		fmt.Println(err)
		return team{},fmt.Errorf("%v",err)
	}
	return p,nil
}


func newplayer(renderer *sdl.Renderer,x,y int32) (p player, err error){

	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
		playerImg,err:= img.Load("/home/bhavya/go/src/sdl/Player_Red.png")
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

		p.x=float64(x)
		p.y=float64(y)

		return  p, nil
}

func  playerDraw (p *player,renderer *sdl.Renderer,x,y int32) (*sdl.Renderer){
	renderer.Copy(p.tex,
		&sdl.Rect{0,0,25,30},
		&sdl.Rect{int32(p.x),int32(p.y),playerWidth,playerHeight})
	return renderer
}

func (p *team) draw (renderer *sdl.Renderer){
	playerDraw(&p.goalKeeper,renderer,int32(p.goalKeeper.x),int32(p.goalKeeper.y))
	playerDraw(&p.leftBack,renderer,int32(p.leftBack.x),int32(p.leftBack.y))
	playerDraw(&p.rightBack,renderer,int32(p.rightBack.x),int32(p.rightBack.y))
	playerDraw(&p.centerMid,renderer,int32(p.centerMid.x),int32(p.centerMid.y))
	playerDraw(&p.rmid,renderer,int32(p.rmid.x),int32(p.rmid.y))
	playerDraw(&p.rrmid,renderer,int32(p.rrmid.x),int32(p.rrmid.y))
	playerDraw(&p.lmid,renderer,int32(p.lmid.x),int32(p.lmid.y))
	playerDraw(&p.llmid,renderer,int32(p.llmid.x),int32(p.llmid.y))
	playerDraw(&p.striker,renderer,int32(p.striker.x),int32(p.striker.y))
	playerDraw(&p.leftWing,renderer,int32(p.leftWing.x),int32(p.leftWing.y))
	playerDraw(&p.rightWing,renderer,int32(p.rightWing.x),int32(p.rightWing.y))
}

func (p *team) update(){
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_UP]==1{
		if p.rmid.y>0 {
			p.rmid.y -= playerSpeed
		}else{
			p.rmid.y=0
		}
	}else if keys[sdl.SCANCODE_DOWN] == 1{
		if p.rmid.y <boxHeight - playerHeight -1{
			p.rmid.y += playerSpeed
		}else{
			p.rmid.y = boxHeight - playerHeight - 1
		}
	}
}










































