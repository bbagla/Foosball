package main

import "github.com/veandco/go-sdl2/sdl"


//GameStatus : struct representing status of the game at any
//given instant. This will be communicated in json format
type GameStatus struct {
	Teams      []team        `json:"player"`
	Ball       ball          `json:"ball"`
	Score      []int         `json:"score"`
	// LastMotion int32         `json:"lastmotion"`
	// LastStick  []player      `json:"laststick"`
	Renderer   *sdl.Renderer `json:"renderer"`
}

//KeyboardInput : struct representing keyboard input.
//This will be sent to the server in json format
type KeyboardInput struct {
	TeamID      int8 //team1/team2
	SelectStick int8 //A/S/D/F
	KeyPressed  int8 //up/down
}

type team struct {
	GoalKeeper [1]player
	Defence    [2]player
	Mid        [5]player
	Attack     [3]player
	LastMotion int32
	LastStick  []player
}

type player struct {
	Circle
	Tex *sdl.Texture
}

type ball struct {
	Circle
	Tex    *sdl.Texture
	Xv, Yv float64
}

//Circle : struct representing center co-ordinates and the radius
type Circle struct {
	X, Y, Radius float64
}
