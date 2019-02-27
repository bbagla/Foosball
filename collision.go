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

func (c1 *circle) collides(c2 player) bool {
	distance := math.Sqrt(math.Pow(c2.x-c1.x, 2) + math.Pow(c2.y-c1.y, 2))
	return distance <= c1.radius+c2.radius
}

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
					onCollisionwithPlayer(ball, teamid)
					break
				}
			}
		}
	}
}

func onCollisionwithPlayer(ball *ball, teamid int32) {
	if (ball.xv < 0 && teamid == 1) || (ball.xv > 0 && teamid == 2) {
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
