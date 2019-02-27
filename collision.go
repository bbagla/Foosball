package main

import (
	"math"
	"sort"
)

func (ball *ball) nearestObstacle() int {
	distance := [8]int32{61, 136, 211, 286, 361, 436, 511, 586}

	i := sort.Search(len(distance), func(i int) bool { return distance[i] >= int32(ball.x) })
	if i < len(distance) && distance[i] == int32(ball.x) {
		return i
	} else {
		if math.Abs(float64(distance[i]-int32(ball.x))) < math.Abs(float64(distance[i-1]-int32(ball.x))) {
			return i
		} else {
			return i - 1
		}
	}
}

func (c1 *circle) collides(c2 circle) bool {
	distance := math.Sqrt(math.Pow(c2.x-c1.x, 2) + math.Pow(c2.y-c1.y, 2))
	return distance <= c1.radius+c2.radius
}

func onCollision(collides bool, ball *ball, teamid int32) {
	if collides && ((ball.xv < 0 && teamid == 1) || (ball.xv > 0 && teamid == 2)) {
		ball.xv = -ball.xv
	}
}

func (c1 *circle) collidesWall() (goal int, index int) {
	if c1.x <= boundarywidth-radius {
		if c1.y <= 297 && c1.y >= 201 {
			return 1, -1
		} else {
			return 0, 1
		}
	} else if c1.x >= boxWidth-boundarywidth-radius {
		if c1.y <= 297 && c1.y >= 201 {
			return 2, -1
		} else {
			return 0, 2
		}
	} else if c1.y <= boundarywidth-radius {
		return 0, 3
	} else if c1.y >= boxHeight-boundarywidth-radius {
		return 0, 4
	}
	return -1, -1
}

func onCollisionWithWall(ball *ball, index int) {
	if index == 1 || index == 2 {
		ball.xv = -ball.xv
	} else if index == 3 || index == 4 {
		ball.yv = -ball.yv
	}
}
