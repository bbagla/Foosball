package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// const (
// 	playerSpeed   = 5
// 	playerWidth   = 26
// 	playerHeight  = 30
// 	boundarywidth = 29
// )

// type team struct {
// 	GoalKeeper [1]player
// 	defence    [2]player
// 	mid        [5]player
// 	attack     [3]player
// }

// type player struct {
// 	circle
// 	tex *sdl.Texture
// }

func newteam(renderer *sdl.Renderer, teamid int32) (t team, err error) {
	offset := int32(0)
	if teamid == 2 {
		offset = boxWidth - 1 - 2*playerWidth
	}
	t.GoalKeeper[0], err = newplayer(renderer, int32(math.Abs(float64(61-playerWidth-offset))), boxHeight/2-playerHeight, teamid)
	if err != nil {
		fmt.Println(err)
		return team{}, fmt.Errorf("%v", err)
	}
	for i := range t.Defence {
		t.Defence[i], err = newplayer(renderer, int32(math.Abs(float64(136-playerWidth-offset))), int32(boxHeight*(i+1))/3-playerHeight, teamid)
		if err != nil {
			fmt.Println(err)
			return team{}, fmt.Errorf("%v", err)
		}
	}
	for i := range t.Mid {
		t.Mid[i], err = newplayer(renderer, int32(math.Abs(float64(286-playerWidth-offset))), int32(boxHeight*(i+1))/6-playerHeight, teamid)
		if err != nil {
			fmt.Println(err)
			return team{}, fmt.Errorf("%v", err)
		}
	}
	for i := range t.Attack {
		t.Attack[i], err = newplayer(renderer, int32(math.Abs(float64(436-playerWidth-offset))), int32(boxHeight*(i+1))/4-playerHeight, teamid)
		if err != nil {
			fmt.Println(err)
			return team{}, fmt.Errorf("%v", err)
		}
	}
	return t, nil
}

func newplayer(renderer *sdl.Renderer, x, y, teamid int32) (p player, err error) {

	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	playerImg, err := img.Load("Player_Red.png")
	if teamid == 2 {
		playerImg, err = img.Load("Player_Blue.png")
	}
	if err != nil {
		fmt.Println(err)
		return player{}, fmt.Errorf("%v", err)
	}
	defer playerImg.Free()
	p.Tex, err = renderer.CreateTextureFromSurface(playerImg)
	if err != nil {
		fmt.Println(err)
		return player{}, fmt.Errorf("%v", err)
	}

	p.X = float64(x + playerWidth/2)
	p.Y = float64(y + playerHeight/2)
	p.Radius = 13
	return p, nil
}

func playerDraw(p *player, renderer *sdl.Renderer) *sdl.Renderer {
	renderer.Copy(p.Tex,
		&sdl.Rect{0, 0, playerWidth, playerHeight},
		&sdl.Rect{int32(p.X), int32(p.Y), playerWidth, playerHeight})
	return renderer
}

func (t *team) draw(renderer *sdl.Renderer) {
	playerDraw(&t.GoalKeeper[0], renderer)
	for i := range t.Defence {
		playerDraw(&t.Defence[i], renderer)
	}
	for i := range t.Mid {
		playerDraw(&t.Mid[i], renderer)
	}
	for i := range t.Attack {
		playerDraw(&t.Attack[i], renderer)
	}
}

func (t *team) update() () {
	t.LastMotion = 0
	if keyboardInput.KeyPressed == 1 {
		t.LastMotion = 1
		if t.LastStick[0].Y > boundarywidth {
			for i := range t.LastStick {
				if t.LastStick[i].Y > boundarywidth {
					t.LastStick[i].Y -= playerSpeed
				}
			}
		}
	} else if keyboardInput.KeyPressed == 2 {
		t.LastMotion = -1
		if t.LastStick[len(t.LastStick)-1].Y < boxHeight-boundarywidth-playerHeight-1 {
			for i := range t.LastStick {
				if t.LastStick[i].Y < boxHeight-playerHeight-boundarywidth-1 {
					t.LastStick[i].Y += playerSpeed
				}
			}
		}
	}
}