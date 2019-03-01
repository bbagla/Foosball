package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed   = 5
	playerWidth   = 26
	playerHeight  = 30
	boundarywidth = 29
)

// type team struct {
// 	goalKeeper [1]player
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

func (t *team) update(last_stick []player, last_motion int32) ([]player, int32) {
	keys := sdl.GetKeyboardState()
	var stick1 = t.GoalKeeper[0:1]
	var stick2 = t.Defence[0:2]
	var stick3 = t.Mid[0:5]
	var stick4 = t.Attack[0:3]
	if keys[sdl.SCANCODE_A] == 1 {
		last_stick = stick1
	} else if keys[sdl.SCANCODE_S] == 1 {
		last_stick = stick2
	} else if keys[sdl.SCANCODE_D] == 1 {
		last_stick = stick3
	} else if keys[sdl.SCANCODE_F] == 1 {
		last_stick = stick4
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		last_motion = 1
		if last_stick[0].Y > boundarywidth {
			for i := range last_stick {
				if last_stick[i].Y > boundarywidth {
					last_stick[i].Y -= playerSpeed
				}
			}
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		last_motion = -1
		if last_stick[len(last_stick)-1].Y < boxHeight-boundarywidth-playerHeight-1 {
			for i := range last_stick {
				if last_stick[i].Y < boxHeight-playerHeight-boundarywidth-1 {
					last_stick[i].Y += playerSpeed
				}
			}
		}
	}
	return last_stick, last_motion
}