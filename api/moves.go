package api

import "net/http"

// MoveRequest is the incoming HTTP request that contains the state of the game and requires a movement response: https://docs.battlesnake.com/references/api/sample-move-request
type MoveRequest struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

// MoveResponse is the response to the MoveRequest: https://docs.battlesnake.com/references/api/sample-move-request
type MoveResponse struct {
	Move  string `json:"move"`  // one of "up" "down" "left" "right"
	Shout string `json:"shout"` // what to yell
}

// Bind is called after render binds the data from the body into the struct
func (data *MoveRequest) Bind(r *http.Request) error {
	return nil
}

var moves = []string{"left", "right", "up", "down"}
