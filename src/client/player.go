package client

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

//creates a new team and their coordinates are passed into  it manually
func newteam(renderer *sdl.Renderer, teamid int32) (t Team, err error) {
	offset := int32(0)
	if teamid == 2 {
		offset = BoxWidth - 1 - 2*playerWidth
	}
	t.GoalKeeper[0], err = newplayer(renderer, int32(math.Abs(float64(61-playerWidth-offset))), BoxHeight/2-playerHeight, teamid)
	if err != nil {
		fmt.Println(err)
		return Team{}, fmt.Errorf("%v", err)
	}
	for i := range t.Defence {
		t.Defence[i], err = newplayer(renderer, int32(math.Abs(float64(136-playerWidth-offset))), int32(BoxHeight*(i+1))/3-playerHeight, teamid)
		if err != nil {
			fmt.Println(err)
			return Team{}, fmt.Errorf("%v", err)
		}
	}
	for i := range t.Mid {
		t.Mid[i], err = newplayer(renderer, int32(math.Abs(float64(286-playerWidth-offset))), int32(BoxHeight*(i+1))/6-playerHeight, teamid)
		if err != nil {
			fmt.Println(err)
			return Team{}, fmt.Errorf("%v", err)
		}
	}
	for i := range t.Attack {
		t.Attack[i], err = newplayer(renderer, int32(math.Abs(float64(436-playerWidth-offset))), int32(BoxHeight*(i+1))/4-playerHeight, teamid)
		if err != nil {
			fmt.Println(err)
			return Team{}, fmt.Errorf("%v", err)
		}
	}
	return t, nil
}

//imports image and attaches it with the player , coordinates are passed into this function for initialization
func newplayer(renderer *sdl.Renderer, x, y, teamid int32) (p Player, err error) {

	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	playerImg, err := img.Load("Player_Red.png")
	if teamid == 2 {
		playerImg, err = img.Load("Player_Blue.png")
	}
	if err != nil {
		fmt.Println(err)
		return Player{}, fmt.Errorf("%v", err)
	}
	defer playerImg.Free()
	p.Tex, err = renderer.CreateTextureFromSurface(playerImg)
	if err != nil {
		fmt.Println(err)
		return Player{}, fmt.Errorf("%v", err)
	}

	p.X = float64(x + playerWidth/2)
	p.Y = float64(y + playerHeight/2)
	p.Radius = 13
	return p, nil
}

//copies texture from player image to the renderer
func playerDraw(p *Player, renderer *sdl.Renderer) *sdl.Renderer {
	renderer.Copy(p.Tex,
		&sdl.Rect{X: 0, Y: 0, W: playerWidth, H: playerHeight},
		&sdl.Rect{X: int32(p.X), Y: int32(p.Y), W: playerWidth, H: playerHeight})
	return renderer
}

//calls playerDraw for each and every team member of the team which is called
func (t *Team) Draw(renderer *sdl.Renderer) {
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
