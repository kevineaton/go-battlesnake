package api

import (
	"fmt"
	"math"
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

const (
	BOARD_HEAD   = "| H |"
	BOARD_BODY   = "| B |"
	BOARD_FOOD   = "| F |"
	BOARD_HAZARD = "| Z |"
	BOARD_ENEMY  = "| E |"
)

func DecideNextMove(request *MoveRequest) (string, error) {
	// set up the logic
	move := ""

	// for this initial implementation, we are just going to brute force and guess so
	// I can see what kind of extra helpers I may need
	// this is NOT an efficient algorithm, to be 100% clear

	// first, see if there is even a valid option by checking the immediate surroundings
	height := request.Board.Height
	width := request.Board.Width
	snakeHeadX := request.You.Head.X
	snakeHeadY := request.You.Head.Y
	laidOutBoard := [][]string{}
	for y := height; y >= 0; y-- {
		row := make([]string, width)
		for x := 0; x < width; x++ {
			row[x] = "|   |"
			// let's look at whether there is a snake, food, hazard, or wall here
			// first, figure out the snake

			// The player first, body, then head
			for _, s := range request.You.Body {
				if s.X == x && s.Y == y {
					row[x] = "| B |"
				}
			}
			if snakeHeadX == x && snakeHeadY == y {
				row[x] = "| H |"
			}

			// food
			for _, f := range request.Board.Food {
				if f.X == x && f.Y == y {
					row[x] = "| F |"
				}
			}

			// hazards
			for _, h := range request.Board.Hazards {
				if h.X == x && h.Y == y {
					row[x] = "| Z |"
				}
			}
		}
		fmt.Printf("%v\n", row)
		laidOutBoard = append(laidOutBoard, row)
	}

	// we now have the board laid out for eay visualization and debugging, so
	// now we need to find out which moves would be valid
	leftValid := true
	rightValid := true
	upValid := true
	downValid := true

	// if the head is next to any wall, we can't go in that direction
	if snakeHeadX == 0 {
		leftValid = false
	}
	if snakeHeadX == width-1 {
		rightValid = false
	}

	if snakeHeadY == 0 {
		downValid = false
	}
	if snakeHeadY == height-1 {
		upValid = false
	}

	// we can't go on top of ourselves, hazards, or enemies; the walls should have accounted for out of bounds errors
	// we should also make sure that this wouldn't be a fatal choice for the snake by going where he cannot go
	// start by checking to the immediate right
	if rightValid {
		rightPiece := laidOutBoard[height-snakeHeadY][snakeHeadX+1]
		if rightPiece == BOARD_BODY || rightPiece == BOARD_ENEMY || rightPiece == BOARD_HAZARD {
			rightValid = false
		}
	}

	if leftValid {
		leftPiece := laidOutBoard[height-snakeHeadY][snakeHeadX-1]
		if leftPiece == BOARD_BODY || leftPiece == BOARD_ENEMY || leftPiece == BOARD_HAZARD {
			leftValid = false
		}
	}

	if upValid {
		upPiece := laidOutBoard[height-snakeHeadY-1][snakeHeadX]
		if upPiece == BOARD_BODY || upPiece == BOARD_ENEMY || upPiece == BOARD_HAZARD {
			upValid = false
		}
	}

	if downValid {
		downPiece := laidOutBoard[height-snakeHeadY+1][snakeHeadX]
		if downPiece == BOARD_BODY || downPiece == BOARD_ENEMY || downPiece == BOARD_HAZARD {
			downValid = false
		}
	}

	// potential decision point: if we only have one valid direction, so we could
	// probably just return from here

	// figure out where the nearest food is, important for staying alive; we do this if health < 50%
	if request.You.Health < 50 { // for now, a simple distance formula will be used, BUT we will likely want to "sim out" getting to the food
		// and eventually consider what other snakes on the board might be going for
		closestFoodIndex := -1
		closestFoodDistance := 9999999.9
		for i := range request.Board.Food {
			f := request.Board.Food[i]
			distance := math.Sqrt((math.Pow((float64(f.X-snakeHeadX)), 2) + math.Pow((float64(f.Y-snakeHeadY)), 2)))
			if distance < closestFoodDistance {
				closestFoodIndex = i
			}
		}
		// figure out how to get closer to the closest food
		// begin trying left and right and if they won't work or it's on the same axis, try y
		if closestFoodIndex >= 0 {
			f := request.Board.Food[closestFoodIndex]
			if f.X < snakeHeadX && leftValid {
				move = "left"
			} else if f.X > snakeHeadX && rightValid {
				move = "right"
			} else if f.Y > snakeHeadY && upValid {
				move = "up"
			} else if f.Y < snakeHeadY && downValid {
				move = "down"
			}

		}
	}

	// make a decision

	// move hasn't been set yet (such as by the food algorithm)
	if move == "" {
		validMoves := []string{}
		if rightValid {
			validMoves = append(validMoves, "right")
		} else if upValid {
			validMoves = append(validMoves, "up")
		} else if leftValid {
			validMoves = append(validMoves, "left")
		} else if downValid {
			validMoves = append(validMoves, "down")
		}
		// use the first one, but when we add in persistence, we will want to
		// look at what the last move was and try to go in a circle rather than in a straight line
		if len(validMoves) == 0 {
			// snake's gonna die anyway, so go up
			return "up", nil
		}
		move = validMoves[0]
	}

	return move, nil
}
