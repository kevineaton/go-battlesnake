package api

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

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

type laidOutBoard struct {
	Board [][]string
}

const (
	BOARD_HEAD       = "| H |"
	BOARD_BODY       = "| B |"
	BOARD_FOOD       = "| F |"
	BOARD_HAZARD     = "| Z |"
	BOARD_ENEMY_BODY = "| X |"
	BOARD_ENEMY_HEAD = "| E |"
)

func (p1 Point) getDistanceFrom(p2 Point) float64 {
	return math.Sqrt((math.Pow((float64(p2.X-p1.X)), 2) + math.Pow((float64(p2.Y-p1.Y)), 2)))
}

func (board *laidOutBoard) print() error {
	var writer io.Writer
	if Config.PrintBoardBeforeTurn == "" || Config.PrintBoardBeforeTurn == "no" {
		return nil
	}
	// TODO: if the user provides a file, we want to write to it in append mode
	if Config.PrintBoardBeforeTurn == "stdout" || Config.PrintBoardBeforeTurn == "out" || Config.PrintBoardBeforeTurn == "yes" {
		writer = os.Stdout
	} else if Config.PrintBoardBeforeTurn == "stdout" || Config.PrintBoardBeforeTurn == "out" {
		writer = os.Stderr
	} else {
		return errors.New("invalid writer specified")
	}
	fmt.Fprintf(writer, "-------------------------------------------------------------\n")
	for _, col := range board.Board {
		fmt.Fprintf(writer, "%v\n", col)
	}
	fmt.Fprintf(writer, "-------------------------------------------------------------\n")
	return nil
}
