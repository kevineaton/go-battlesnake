package api

// Game is the meta data about a Game: https://docs.battlesnake.com/references/api#game
type Game struct {
	ID      string `json:"id"`
	Ruleset struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"ruleset"`
	Timeout int `json:"timeout"` // in milliseconds
}

// Point is an x:y coordinate on the board and is used for board layout and snake placement
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Board is the primary view of the state of a game: https://docs.battlesnake.com/references/api#board
type Board struct {
	Height  int           `json:"height"`
	Width   int           `json:"width"`
	Food    []Point       `json:"food"`
	Hazards []Point       `json:"hazards"`
	Snakes  []Battlesnake `json:"snakes"`
}
