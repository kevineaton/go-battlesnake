package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecideMoveDefault(t *testing.T) {
	ConfigSetup()

	request := simpleMoveRequestMove1
	request2 := simpleMoveRequestMove2

	move, err := DecideNextMove(&request)
	assert.Nil(t, err)
	assert.Equal(t, "up", move)
	move, err = DecideNextMove(&request2)
	assert.Nil(t, err)
	assert.Equal(t, "right", move)
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
		Length:  2,
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
				Length:  2,
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
			{X: 0, Y: 1},
			{X: 0, Y: 0},
			{X: 1, Y: 0},
		},
		Latency: "10",
		Head:    Point{X: 0, Y: 1},
		Length:  2,
		Shout:   "Yo yo yo",
	},
}
