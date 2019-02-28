package main

import (
	"math"
	"sort"
)

//nearestObstacle() returns a integer corresponding to the stick that is nearest to the ball at that moment.
func (ball *ball) nearestObstacle() int {
	distance := [8]int32{61, 136, 211, 286, 361, 436, 511, 586}

	i := sort.Search(len(distance), func(i int) bool { return distance[i] >= int32(ball.x) })
	if i < len(distance) && distance[i] == int32(ball.x) {
		return i
	} else if i == 0 {
		return i
	} else if i == 8 {
		return i - 1
	} else {
		if math.Abs(float64(distance[i]-int32(ball.x))) < math.Abs(float64(distance[i-1]-int32(ball.x))) {
			return i
		} else {
			return i - 1
		}
	}
}

//collides() returns true if the ball collides with a player.
func (c1 *ball) collides(c2 player) bool {
	distance := math.Sqrt(math.Pow(c2.x-c1.x+c2.radius, 2) + math.Pow(c2.y-c1.y+c2.radius, 2))
	return distance <= c1.radius+c2.radius
}

//CheckCollision() checks at every frame if a collision has happened with a player or not.
//If a collision has happened, it invokes onCollisionwithPlayer().
func (ball *ball) CheckCollision(t team, teamid int32) {
	index := ball.nearestObstacle()
	arr := [2][]int{{0, 1, 3, 5}, {7, 6, 4, 2}}
	var stick [4][]player
	stick[0] = t.goalKeeper[0:1]
	stick[1] = t.defence[0:2]
	stick[2] = t.mid[0:5]
	stick[3] = t.attack[0:3]
	for i, j := range arr[teamid-1] {
		if j == index {
			for k := range stick[i] {
				if ball.collides(stick[i][k]) {
					onCollisionwithPlayer(ball, teamid, t.lastMotion)
					break
				}
			}
		}
	}
}

//onCollisionwithPlayer() changes the direction of the ball.
//It also changes the speed of the ball if it has not been increased by a collision with another player.
func onCollisionwithPlayer(ball *ball, teamid int32, lastMotion int32) {
	if (ball.xv < 0 && teamid == 1) || (ball.xv > 0 && teamid == 2) {
		ball.xv = -ball.xv
	}
	if math.Abs(ball.xv) <= BallSpeedX {
		ball.xv *= 2
		ball.yv *= 2
	}
	ball.yv += float64(lastMotion) * 0.2
}

//collidesWall() checks if a collision has happened with a wall or not.
//It also checks and returns if a goal has happened.
func (c1 *ball) collidesWall() (goal int, index int) {
	//index -1 means no collision
	// 1 means collision with left wall
	// 2 right wall
	// 3 upper wall
	// 4 means lower wall
	if c1.x < boundaryWidth+radius && c1.xv < 0 {
		if c1.y <= 297-radius && c1.y >= 201+radius {
			insideGoal = true
			return 2, -1
		} else {
			return 0, 1
		}
	} else if c1.x > boxWidth-boundaryWidth-radius-1 && c1.xv > 0 {
		if c1.y <= 297-radius && c1.y >= 201+radius {
			insideGoal = true
			return 1, -1
		} else {
			return 0, 2
		}
	} else if c1.y < boundaryWidth+radius && c1.yv < 0 {
		return 0, 3
	} else if c1.y > boxHeight-boundaryWidth-radius-1 && c1.yv > 0 {
		return 0, 4
	}
	return 0, -1
}

//onCollisionWithWall() changes the direction and speed of the ball.
func onCollisionWithWall(ball *ball, index int) {
	if index == 1 || index == 2 {
		ball.xv = -ball.xv
	} else if index == 3 || index == 4 {
		ball.yv = -ball.yv
	}
	if ball.xv > BallSpeedX {
		ball.xv /= 2
		ball.yv = BallSpeedY * (ball.yv / math.Abs(ball.yv))
	}
}

//movementInsidePost() decides the movement of the ball inside the goal post.
func (bal *ball) movementInsidePost() {
	bal.yv = 0
}
