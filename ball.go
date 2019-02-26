package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	radius = 8
)

type ball struct {
	tex                   *sdl.Texture
	x, y, velocity, theta float64
}

func (ball *ball) draw(renderer *sdl.Renderer) *sdl.Renderer {
	renderer.Copy(ball.tex,
		&sdl.Rect{0, 0, 2 * radius, 2 * radius},
		&sdl.Rect{int32(ball.x), int32(ball.y), 2 * radius, 2 * radius})
	return renderer
}

func newBall(renderer *sdl.Renderer, x, y int32) (bal ball, err error) {

	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	playerImg, err := img.Load("Ball.png")
	if err != nil {
		fmt.Println(err)
		return ball{}, fmt.Errorf("%v", err)
	}
	defer playerImg.Free()
	bal.tex, err = renderer.CreateTextureFromSurface(playerImg)
	if err != nil {
		fmt.Println(err)
		return ball{}, fmt.Errorf("%v", err)
	}

	bal.x = float64(x)
	bal.y = float64(y)

	return bal, nil
}
