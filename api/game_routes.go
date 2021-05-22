package api

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/go-chi/render"
)

// GameRequest is what the API will send when a snake is entered into a new game
type GameRequest struct {
	Game  Game        `json:"game"`
	Turn  int64       `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

// Bind is called after render binds the data from the body into the struct
func (data *GameRequest) Bind(r *http.Request) error {
	return nil
}

// GetSnakeRoute gets the attributes about a snake. In the current UI, this route is not easily seen (missing the HTTP method in the menu)
func GetSnakeRoute(w http.ResponseWriter, r *http.Request) {

	// this is simple and returns just the configured variables
	Send(w, http.StatusOK, map[string]string{
		"apiversion": "1",
		"author":     Config.Author,
		"color":      Config.SnakeColor,
		"head":       Config.SnakeHead,
		"tail":       Config.SnakeTail,
		"version":    Config.Version,
	})
}

// GameStartRoute is called when a new game is starting: https://docs.battlesnake.com/references/api#start
func GameStartRoute(w http.ResponseWriter, r *http.Request) {

	request := &GameRequest{}
	render.Bind(r, request)

	if request.Turn > 0 || request.Game.ID == "" {
		SendError(&w, r, http.StatusBadRequest, "game_start_bad_data", "the data for the game to begin is invalid", &map[string]interface{}{
			"request": request,
		})
		return
	}

	fmt.Printf("\n-----Received Start request-----\n%+v\n---------\n", request)
	// responses are ignored
	Send(w, http.StatusOK, map[string]string{})
}

// MoveRequestRoute prompts the server for a new move
func MoveRequestRoute(w http.ResponseWriter, r *http.Request) {

	request := &MoveRequest{}
	render.Bind(r, request)

	fmt.Printf("\n-----Received Move request-----\n%+v\n---------\n", request)

	// for initial testing, just choose a random direction
	// TODO: break off into actual logic
	direction := moves[rand.Intn(4)]
	response := MoveResponse{
		Move:  direction,
		Shout: "Yo yo yo",
	}
	Send(w, http.StatusOK, response)
}

// GameEndRoute is called at the end of the game and represents the final state of the game: https://docs.battlesnake.com/references/api#end
func GameEndRoute(w http.ResponseWriter, r *http.Request) {
	request := &GameRequest{}
	render.Bind(r, request)

	if request.Turn <= 0 || request.Game.ID == "" {
		SendError(&w, r, http.StatusBadRequest, "game_end_bad_data", "the data for the game to begin is invalid", &map[string]interface{}{
			"request": request,
		})
		return
	}

	fmt.Printf("\n-----Received End request-----\n%+v\n---------\n", request)

	// TODO: we can log the results for future analysis

	// responses are ignored
	Send(w, http.StatusOK, map[string]string{})
}
