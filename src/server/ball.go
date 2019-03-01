package serve

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//flag for checking if ball is in the goal or not
//insideGoal : false means ball in not inside the goal
var insideGoal = false

//draw function for drawing the ball
func (ball *ball) draw(renderer *sdl.Renderer) *sdl.Renderer {
	renderer.Copy(ball.Tex,
		&sdl.Rect{X: 0, Y: 0, W: 2 * radius, H: 2 * radius},
		&sdl.Rect{X: int32(ball.X - radius), Y: int32(ball.Y - radius), W: 2 * radius, H: 2 * radius})
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
	bal.Tex, err = renderer.CreateTextureFromSurface(BallImg)
	if err != nil {
		fmt.Println(err)
		return ball{}, fmt.Errorf("%v", err)
	}

	bal.X = float64(x)
	bal.Y = float64(y)
	bal.Radius = radius
	bal.Xv = BallSpeedX
	bal.Yv = BallSpeedY
	return bal, nil
}

//update function for ball
func (ball *ball) update() {

	goalID, index := ball.CollidesWall()
	if index != -1 {
		OnCollisionWithWall(ball, index)
	}
	if insideGoal == true {
		ball.movementInsidePost()
		if ball.X+radius < 0 || ball.X > boxWidth-1+radius {
			ball.reset(goalID)
			fmt.Println("GoalID is: ", goalID)
			insideGoal = false
			gameStatus.Score[goalID-1]++
			fmt.Println(gameStatus.Score[0], ":", gameStatus.Score[1])
		}
	}
	ball.X += ball.Xv
	ball.Y += ball.Yv
}

//when goal is scored this function restores the initial position of the ball
func (ball *ball) reset(goal int) {
	if goal == 1 || goal == 2 {
		sdl.Delay(2000)
		ball.X = float64(boxWidth / 2)
		ball.Y = float64(boxHeight / 2)
		ball.Yv = BallSpeedY
		if goal == 2 {
			ball.Xv = -BallSpeedX
		} else if goal == 1 {
			ball.Xv = BallSpeedX
		}
	}
}
