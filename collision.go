package main

import (
	"math"
	"sort"
)

func (ball *ball) nearestObstacle() int {
	distance := [10]int32{29, 61, 136, 211, 286, 361, 436, 511, 586, 619}

	i := sort.Search(len(distance), func(i int) bool { return distance[i] >= int32(ball.x) })
	if i < len(distance) && distance[i] == int32(ball.x) {
		return i
	} else {
		if distance[i]-int32(ball.x) < distance[i-1]-int32(ball.x) {
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
