package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//BallSpeedX : Speed of ball in X-direcrion
//BallSpeedY : Speed of ball in Y-direcrion
//radius     : radius of ball
const (
	radius     = 8
	BallSpeedX = 1.5
	BallSpeedY = 0.5
)

//flag for checking if ball is in the goal or not
//insideGoal : false means ball in not inside the goal
var insideGoal = false

//circle struct is inherited
//ball has its texture
//xv  : pixels moved by ball in one frame in X-direction
//yv  : pixels moved by ball in one frame in Y-direction
type ball struct {
	circle
	tex    *sdl.Texture
	xv, yv float64
}

//draw function for drawing the ball
func (ball *ball) draw(renderer *sdl.Renderer) *sdl.Renderer {
	renderer.Copy(ball.tex,
		&sdl.Rect{0, 0, 2 * radius, 2 * radius},
		&sdl.Rect{int32(ball.x - radius), int32(ball.y - radius), 2 * radius, 2 * radius})
	return renderer
}

//it imports a texture to the ball and defines it initial coordinates
func newBall(renderer *sdl.Renderer, x, y int32) (bal ball, err error) {
	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	BallImg, err := img.Load("Ball.png")
	if err != nil {
		fmt.Println(err)
		return ball{}, fmt.Errorf("%v", err)
	}
	defer BallImg.Free()
	bal.tex, err = renderer.CreateTextureFromSurface(BallImg)
	if err != nil {
		fmt.Println(err)
		return ball{}, fmt.Errorf("%v", err)
	}
	bal.x = float64(x)
	bal.y = float64(y)
	bal.radius = radius
	bal.xv = BallSpeedX
	bal.yv = BallSpeedY
	return bal, nil
}

//update function for ball
func (ball *ball) update() {
	goalId, index := ball.collidesWall()
	if index != -1 {
		onCollisionWithWall(ball, index)
	}
	if insideGoal == true {
		ball.movementInsidePost()
		if ball.x+radius < 0 || ball.x > boxWidth-1+radius {
			ball.reset(goalId)
			insideGoal = false
			gameStatus.Score[goalId-1]++
		}
	}
	ball.x += ball.xv
	ball.y += ball.yv
}

//when goal is scored this function restores the initial position of the ball
func (ball *ball) reset(goal int) {
	if goal == 1 || goal == 2 {
		sdl.Delay(2000)
		ball.x = float64(boxWidth / 2)
		ball.y = float64(boxHeight / 2)
		ball.yv = BallSpeedY
		if goal == 2 {
			ball.xv = -BallSpeedX
		} else if goal == 1 {
			ball.xv = BallSpeedX
		}
	}
}
