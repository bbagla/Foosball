package main

import (
	"math"
	"sort"
)

func (ball *ball) nearestObstacle() int {
	distance := [8]int32{61, 136, 211, 286, 361, 436, 511, 586}

	i := sort.Search(len(distance), func(i int) bool { return distance[i] >= int32(ball.X) })
	if i < len(distance) && distance[i] == int32(ball.X) {
		return i
	} else if i == 0 {
		return i
	} else if i == 8 {
		return i - 1
	} else {
		if math.Abs(float64(distance[i]-int32(ball.X))) < math.Abs(float64(distance[i-1]-int32(ball.X))) {
			return i
		} else {
			return i - 1
		}
	}
}

func (c1 *ball) collides(c2 player) bool {
	distance := math.Sqrt(math.Pow(c2.X-c1.X+c2.Radius, 2) + math.Pow(c2.Y-c1.Y+c2.Radius, 2))
	return distance <= c1.Radius+c2.Radius
}

func (ball *ball) CheckCollision(t team, teamid int32) {
	index := ball.nearestObstacle()
	arr := [2][]int{{0, 1, 3, 5}, {7, 6, 4, 2}}
	var stick [4][]player
	stick[0] = t.GoalKeeper[0:1]
	stick[1] = t.Defence[0:2]
	stick[2] = t.Mid[0:5]
	stick[3] = t.Attack[0:3]
	for i, j := range arr[teamid-1] {
		if j == index {
			for k := range stick[i] {
				if ball.collides(stick[i][k]) {
					onCollisionwithPlayer(ball, teamid)
					break
				}
			}
		}
	}
}

func onCollisionwithPlayer(ball *ball, teamid int32) {
	if (ball.Xv < 0 && teamid == 1) || (ball.Xv > 0 && teamid == 2) {
		ball.Xv = -ball.Xv
	}
	if math.Abs(ball.Xv) <= BallSpeedX {
		ball.Xv *= 2
		ball.Yv *= 2
	}
	// 	ball.Yv += float64(gameStatus.LastMotion) * 0.2
}

func (c1 *ball) collidesWall() (goal int, index int) {
	//index -1 means no collision
	// 1 means collision with left wall
	// 2 right wall
	// 3 upper wall
	// 4 means lower wall
	if c1.X < boundarywidth+radius && c1.Xv < 0 {
		if c1.X <= 297-radius && c1.Y >= 201+radius {
			insideGoal = true
			return 2, -1
		} else {
			return 0, 1
		}
	} else if c1.X > boxWidth-boundarywidth-radius-1 && c1.Xv > 0 {
		if c1.Y <= 297-radius && c1.Y >= 201+radius {
			insideGoal = true
			return 1, -1
		} else {
			return 0, 2
		}
	} else if c1.Y < boundarywidth+radius && c1.Yv < 0 {
		return 0, 3
	} else if c1.Y > boxHeight-boundarywidth-radius-1 && c1.Yv > 0 {
		return 0, 4
	}
	return 0, -1
}

func onCollisionWithWall(ball *ball, index int) {
	if index == 1 || index == 2 {
		ball.Xv = -ball.Xv
	} else if index == 3 || index == 4 {
		ball.Yv = -ball.Yv
	}
	if ball.Xv > BallSpeedX {
		ball.Xv /= 2
		ball.Yv = BallSpeedY * (ball.Yv / math.Abs(ball.Yv))
	}
}

func (bal *ball) movementInsidePost() {
	bal.Yv = 0
}
