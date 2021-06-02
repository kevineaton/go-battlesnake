package api

import (
	"fmt"
	"net/http"
)

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

// for now, keep the last move for each game id in memory
var lastMoves = map[string]string{}

func DecideNextMove(request *MoveRequest) (string, error) {
	// set up the logic
	move := ""

	// for this initial implementation, we are just going to brute force and guess so
	// I can see what kind of extra helpers I may need
	// this is NOT an efficient algorithm, to be 100% clear

	// set up some state and place holders
	height := request.Board.Height
	width := request.Board.Width
	snakeHeadX := request.You.Head.X
	snakeHeadY := request.You.Head.Y
	laidOutBoard := laidOutBoard{}
	closestEnemyHead := Point{}
	closestEnemyHeadDistance := 99999.9

	for y := height; y >= 0; y-- {
		row := make([]string, width)
		for x := 0; x < width; x++ {
			row[x] = "|   |"
			// let's look at whether there is a snake, food, hazard, or wall here
			// first, figure out the snake

			// The player first, body, then head
			for _, s := range request.You.Body {
				if s.X == x && s.Y == y {
					row[x] = BOARD_BODY
				}
			}
			if snakeHeadX == x && snakeHeadY == y {
				row[x] = BOARD_HEAD
			}

			// now snakes
			for _, snake := range request.Board.Snakes {
				if snake.ID != request.You.ID {
					for _, s := range snake.Body {
						if s.X == x && s.Y == y {
							row[x] = BOARD_ENEMY_BODY
						}
					}
					if snake.Head.X == x && snake.Head.Y == y {
						row[x] = BOARD_ENEMY_HEAD
					}
					// figure out if this is the closest snake head
					distanceFromSnake := snake.Head.getDistanceFrom(request.You.Head)
					if distanceFromSnake < closestEnemyHeadDistance {
						closestEnemyHeadDistance = distanceFromSnake
						closestEnemyHead = snake.Head
					}
				}
			}

			// food
			for _, f := range request.Board.Food {
				if f.X == x && f.Y == y {
					row[x] = BOARD_FOOD
				}
			}

			// hazards
			for _, h := range request.Board.Hazards {
				if h.X == x && h.Y == y {
					row[x] = BOARD_HAZARD
				}
			}
		}
		laidOutBoard.Board = append(laidOutBoard.Board, row)
	}
	laidOutBoard.print()

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
		rightPiece := laidOutBoard.Board[height-snakeHeadY][snakeHeadX+1]
		if rightPiece == BOARD_BODY || rightPiece == BOARD_ENEMY_BODY || rightPiece == BOARD_ENEMY_HEAD || rightPiece == BOARD_HAZARD {
			rightValid = false
		}
	}

	if leftValid {
		leftPiece := laidOutBoard.Board[height-snakeHeadY][snakeHeadX-1]
		if leftPiece == BOARD_BODY || leftPiece == BOARD_ENEMY_BODY || leftPiece == BOARD_ENEMY_HEAD || leftPiece == BOARD_HAZARD {
			leftValid = false
		}
	}

	if upValid {
		upPiece := laidOutBoard.Board[height-snakeHeadY-1][snakeHeadX]
		if upPiece == BOARD_BODY || upPiece == BOARD_ENEMY_BODY || upPiece == BOARD_ENEMY_HEAD || upPiece == BOARD_HAZARD {
			upValid = false
		}
	}

	if downValid {
		downPiece := laidOutBoard.Board[height-snakeHeadY+1][snakeHeadX]
		if downPiece == BOARD_BODY || downPiece == BOARD_ENEMY_BODY || downPiece == BOARD_ENEMY_HEAD || downPiece == BOARD_HAZARD {
			downValid = false
		}
	}

	// now let's try to predict the direction the next snake is and try to move away from it
	// for this first version, only consider the actual location of the head and try to increase the gap regardless of likely position
	// start with a distance of 6
	if closestEnemyHeadDistance > 0 && closestEnemyHeadDistance < 6 {
		if leftValid && closestEnemyHead.X > snakeHeadX {
			// try to make x smaller
			move = "left"
		} else if rightValid && closestEnemyHead.X < snakeHeadX {
			// try to make x larger
			move = "right"
		} else if downValid && closestEnemyHead.Y > snakeHeadY {
			move = "down"
		} else if upValid {
			move = "up"
		}
	}

	// potential decision point: if we only have one valid direction, we could exit
	// we could probably clean up this layout
	if downValid && !upValid && !leftValid && !rightValid {
		move = "down"
	}
	if !downValid && upValid && !leftValid && !rightValid {
		move = "up"
	}
	if !downValid && !upValid && leftValid && !rightValid {
		move = "left"
	}
	if !downValid && !upValid && !leftValid && rightValid {
		move = "right"
	}
	if move != "" {
		return move, nil
	}

	// figure out where the nearest food is, important for staying alive; we do this if health < 50%
	// for now, a simple distance formula will be used, BUT we will likely want to "sim out" getting to the food
	if request.You.Health < 30 {
		// and eventually consider what other snakes on the board might be going for
		closestFoodIndex := -1
		closestFoodDistance := 9999999.9
		for i := range request.Board.Food {
			f := request.Board.Food[i]
			distance := f.getDistanceFrom(Point{X: snakeHeadX, Y: snakeHeadY})
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
	// TODO: we should probably introduce weighting throughout the above, such as determining snakes
	fmt.Println("Move is empty")
	if move == "" {
		validMovesMap := map[string]int{
			"right": 0,
			"left":  0,
			"up":    0,
			"down":  0,
		}
		if rightValid {
			// if this is 1 square away from a wall, it gets a score of 50
			// otherwise, give it a 100
			if snakeHeadX == width-2 {
				validMovesMap["right"] = 50
			} else {
				validMovesMap["right"] = 100
			}
		}
		if upValid {
			if snakeHeadY == height-2 {
				validMovesMap["up"] = 50
			} else {
				validMovesMap["up"] = 100
			}
		}
		if leftValid {
			if snakeHeadX == 2 {
				validMovesMap["left"] = 50
			} else {
				validMovesMap["left"] = 100
			}
		}
		if downValid {
			if snakeHeadY == 1 {
				validMovesMap["down"] = 50
			} else {
				validMovesMap["down"] = 100
			}
		}
		fmt.Printf("\nWeights\n%+v\nRight valid: %v - Left valid: %v - Up valid: %v - Down Valid: %v", validMovesMap, rightValid, leftValid, upValid, downValid)

		// ok, so simple; loop through all of the directions; if it's a 100, there we go
		// we want to prioritize moving right, then up, then left, then down
		// note that in Go, maps do not guarantee an order, so we don't range over them
		// again, this is just a rough algo and one we will clean up later
		if validMovesMap["right"] == 100 {
			move = "right"
		} else if validMovesMap["up"] == 100 {
			move = "up"
		} else if validMovesMap["left"] == 100 {
			move = "left"
		} else if validMovesMap["down"] == 100 {
			move = "down"
		}
		if move == "" {
			// move is still empty, so repeat for 50
			if validMovesMap["right"] == 50 {
				move = "right"
			} else if validMovesMap["up"] == 50 {
				move = "up"
			} else if validMovesMap["left"] == 50 {
				move = "left"
			} else if validMovesMap["down"] == 50 {
				move = "down"
			}
		}

	}
	fmt.Printf("\nMoving: %s\n", move)
	lastMoves[request.Game.ID] = move
	return move, nil
}
