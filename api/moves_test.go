package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecideMoveDefault(t *testing.T) {
	ConfigSetup()
	original := Config.PrintBoardBeforeTurn
	Config.PrintBoardBeforeTurn = "stdout"
	request := simpleMoveRequestMove1
	request2 := simpleMoveRequestMove2
	request3 := simpleMoveRequestMove3

	move, err := DecideNextMove(&request)
	assert.Nil(t, err)
	assert.Equal(t, "up", move)
	move, err = DecideNextMove(&request2)
	assert.Nil(t, err)
	assert.Equal(t, "right", move)
	move, err = DecideNextMove(&request3)
	assert.Nil(t, err)
	assert.Equal(t, "right", move)
	Config.PrintBoardBeforeTurn = original
}

var simpleMoveRequestMove1 = MoveRequest{
	Game: Game{
		ID: "Test",
	},
	Board: Board{
		Height: 7,
		Width:  10,
		Food: []Point{
			{
				X: 2,
				Y: 2,
			},
			{
				X: 3,
				Y: 4,
			},
		},
		Snakes: []Battlesnake{
			{
				ID:     "cevvyn",
				Name:   "Cevvyn",
				Health: 98,
				Body: []Point{
					{X: 0, Y: 0},
					{X: 1, Y: 0},
				},
				Latency: "10",
				Head:    Point{X: 0, Y: 0},
				Length:  2,
				Shout:   "Yo yo yo",
			},
			{
				ID:     "enemy",
				Name:   "Enemy",
				Health: 98,
				Body: []Point{
					{X: 1, Y: 2},
				},
				Latency: "10",
				Head:    Point{X: 1, Y: 2},
				Length:  3,
				Shout:   "Yo yo yo",
			},
		},
	},
	Turn: 1,
	You: Battlesnake{
		ID:     "cevvyn",
		Name:   "Cevvyn",
		Health: 98,
		Body: []Point{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
		},
		Latency: "10",
		Head:    Point{X: 0, Y: 0},
		Length:  3,
		Shout:   "Yo yo yo",
	},
}

var simpleMoveRequestMove2 = MoveRequest{
	Game: Game{
		ID: "Test",
	},
	Board: Board{
		Height: 7,
		Width:  10,
		Food: []Point{
			{
				X: 2,
				Y: 2,
			},
			{
				X: 3,
				Y: 4,
			},
		},
		Snakes: []Battlesnake{
			{
				ID:     "cevvyn",
				Name:   "Cevvyn",
				Health: 98,
				Body: []Point{
					{X: 0, Y: 1},
					{X: 0, Y: 0},
					{X: 1, Y: 0},
				},
				Latency: "10",
				Head:    Point{X: 0, Y: 1},
				Length:  3,
				Shout:   "Yo yo yo",
			},
			{
				ID:     "enemy",
				Name:   "Enemy",
				Health: 98,
				Body: []Point{
					{X: 1, Y: 2},
					{X: 0, Y: 2},
				},
				Latency: "10",
				Head:    Point{X: 0, Y: 2},
				Length:  3,
				Shout:   "Yo yo yo",
			},
		},
	},
	Turn: 2,
	You: Battlesnake{
		ID:     "cevvyn",
		Name:   "Cevvyn",
		Health: 98,
		Body: []Point{
			{X: 0, Y: 1},
			{X: 0, Y: 0},
			{X: 1, Y: 0},
		},
		Latency: "10",
		Head:    Point{X: 0, Y: 1},
		Length:  3,
		Shout:   "Yo yo yo",
	},
}

var simpleMoveRequestMove3 = MoveRequest{
	Game: Game{
		ID: "Test",
	},
	Board: Board{
		Height: 7,
		Width:  10,
		Food: []Point{
			{
				X: 2,
				Y: 2,
			},
			{
				X: 3,
				Y: 4,
			},
		},
		Snakes: []Battlesnake{
			{
				ID:     "cevvyn",
				Name:   "Cevvyn",
				Health: 98,
				Body: []Point{
					{X: 1, Y: 1},
					{X: 1, Y: 1},
					{X: 0, Y: 1},
				},
				Latency: "10",
				Head:    Point{X: 1, Y: 1},
				Length:  3,
				Shout:   "Yo yo yo",
			},
			{
				ID:     "enemy",
				Name:   "Enemy",
				Health: 98,
				Body: []Point{
					{X: 1, Y: 3},
					{X: 1, Y: 2},
					{X: 0, Y: 2},
				},
				Latency: "10",
				Head:    Point{X: 0, Y: 2},
				Length:  3,
				Shout:   "Yo yo yo",
			},
		},
	},
	Turn: 2,
	You: Battlesnake{
		ID:     "cevvyn",
		Name:   "Cevvyn",
		Health: 98,
		Body: []Point{
			{X: 1, Y: 1},
			{X: 0, Y: 1},
			{X: 0, Y: 0},
		},
		Latency: "10",
		Head:    Point{X: 1, Y: 1},
		Length:  3,
		Shout:   "Yo yo yo",
	},
}
