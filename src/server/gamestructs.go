package main

import "github.com/veandco/go-sdl2/sdl"

//GameStatus : struct representing status of the game at any given instant. This will be communicated in json format
type GameStatus struct {
	Teams    []team        `json:"player"`
	Ball     ball          `json:"ball"`
	Score    []int         `json:"score"`
	Renderer *sdl.Renderer `json:"renderer"`
}

//KeyboardInput : struct representing keyboard input.
//This will be sent to the server in json format
type KeyboardInput struct {
	TeamID      int8 //team1/team2
	SelectStick int8 //A/S/D/F
	KeyPressed  int8 //up/down
}

//struct for the teams containing some arrays of players
type team struct {
	GoalKeeper [1]player
	Defence    [2]player
	Mid        [5]player
	Attack     [3]player
	LastMotion int32
	LastStick  []player
}

//struct for player it contains an imaginary circle for detecting collision which also defines players position
//and also texture pointer for drawing the player
type player struct {
	Circle
	Tex *sdl.Texture
}

//circle struct is inherited
//ball has its texture
//xv  : pixels moved by ball in one frame in X-direction
//yv  : pixels moved by ball in one frame in Y-direction
type ball struct {
	Circle
	Tex    *sdl.Texture
	Xv, Yv float64
}

//Circle : struct representing center co-ordinates and the radius
type Circle struct {
	X, Y, Radius float64
}
