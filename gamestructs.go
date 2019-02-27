package main

//Position : struct representing position
type Position struct {
	X float64
	Y float64
}

//GameStatus : struct representing status of the game at any
//given instant. This will be communicated in json format
type GameStatus struct {
	Teams []team
	Ball  Position
	Score []int
}
