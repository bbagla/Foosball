package main

import "math"

func (c1 *circle) collides(c2 circle) bool {
	distance := math.Sqrt(math.Pow(c2.x-c1.x, 2) + math.Pow(c2.y-c1.y, 2))
	return distance <= c1.radius+c2.radius
}

func onCollision(collides bool, ball *ball, teamid int32) {
	if collides && ((ball.xv < 0 && teamid == 1) || (ball.xv > 0 && teamid == 2)) {
		ball.xv = -ball.xv
	}

}
