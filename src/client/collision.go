package main

import (
	"math"
	"sort"
)

//nearestObstacle() returns a integer corresponding to the stick that is nearest to the ball at that moment.
func (ball *ball) nearestObstacle(c1 chan int) {
	distance := [8]int32{61, 136, 211, 286, 361, 436, 511, 586}

	i := sort.Search(len(distance), func(i int) bool { return distance[i] >= int32(ball.X) })
	if i < len(distance) && distance[i] == int32(ball.X) {
		c1 <- i
	} else if i == 0 {
		c1 <- i
	} else if i == 8 {
		c1 <- i - 1
	} else {
		if math.Abs(float64(distance[i]-int32(ball.X))) < math.Abs(float64(distance[i-1]-int32(ball.X))) {
			c1 <- i
		} else {
			c1 <- i - 1
		}
	}
}

//collides() returns true if the ball collides with a player.
func (ball *ball) collides(c2 player) bool {
	distance := math.Sqrt(math.Pow(c2.X-ball.X+c2.Radius, 2) + math.Pow(c2.Y-ball.Y+c2.Radius, 2))
	return distance <= ball.Radius+c2.Radius
}

//CheckCollision() checks at every frame if a collision has happened with a player or not.
//If a collision has happened, it invokes onCollisionwithPlayer().
func (ball *ball) CheckCollision(t team, teamid int32) {
	c1 := make(chan int)
	go ball.nearestObstacle(c1)
	arr := [2][]int{{0, 1, 3, 5}, {7, 6, 4, 2}}
	var stick [4][]player
	stick[0] = t.GoalKeeper[0:1]
	stick[1] = t.Defence[0:2]
	stick[2] = t.Mid[0:5]
	stick[3] = t.Attack[0:3]
	index := <-c1
	for i, j := range arr[teamid-1] {
		if j == index {
			for k := range stick[i] {
				go ball.collision(t, teamid, stick[i][k])
			}
		}
	}
}

func (ball *ball) collision(t team, teamid int32, p player) {
	if ball.collides(p) {
		onCollisionwithPlayer(ball, teamid, t.LastMotion)
	}
}

//onCollisionwithPlayer() changes the direction of the ball.
//It also changes the speed of the ball if it has not been increased by a collision with another player.
func onCollisionwithPlayer(ball *ball, teamid int32, lastMotion int32) {
	if (ball.Xv < 0 && teamid == 1) || (ball.Xv > 0 && teamid == 2) {
		ball.Xv = -ball.Xv
	}
	if math.Abs(ball.Xv) <= BallSpeedX {
		ball.Xv *= 2
		ball.Yv *= 2
	}
	ball.Yv += float64(lastMotion) * 0.2
}

//collidesWall() checks if a collision has happened with a wall or not.
//It also checks and returns if a goal has happened.
//index -1 means no collision
// 1 means collision with left wall
// 2 right wall
// 3 upper wall
// 4 means lower wall
func (ball *ball) collidesWall() (goal int, index int) {
	if ball.X < boundarywidth+radius && ball.Xv < 0 {
		if ball.Y <= 297-radius && ball.Y >= 201+radius {
			insideGoal = true
			return 2, -1
		}
		return 0, 1

	} else if ball.X > boxWidth-boundarywidth-radius-1 && ball.Xv > 0 {
		if ball.Y <= 297-radius && ball.Y >= 201+radius {
			insideGoal = true
			return 1, -1
		}
		return 0, 2
	} else if ball.Y < boundarywidth+radius && ball.Yv < 0 {
		return 0, 3
	} else if ball.Y > boxHeight-boundarywidth-radius-1 && ball.Yv > 0 {
		return 0, 4
	}
	return 0, -1
}

//onCollisionWithWall() changes the direction and speed of the ball.
func onCollisionWithWall(ball *ball, index int) {
	if index == 1 || index == 2 {
		ball.Xv = -ball.Xv
	} else if index == 3 || index == 4 {
		ball.Yv = -ball.Yv
	}
	if math.Abs(ball.Xv) > BallSpeedX {
		ball.Xv /= 2
		ball.Yv = BallSpeedY * (ball.Yv / math.Abs(ball.Yv))
	}
}

//movementInsidePost() decides the movement of the ball inside the goal post.
func (ball *ball) movementInsidePost() {
	ball.Yv = 0
}
