package api

import (
	"net/http"

	"github.com/go-chi/render"
)

// ChangeSnakeRoute allows the caller to change the snake details. Note that this is
// current unauthenticated so it could be used for bad intentions. You may want to comment this out in the
// http.go file
func ChangeSnakeRoute(w http.ResponseWriter, r *http.Request) {
	isAuthed, key := checkAuthKey(r)
	if !isAuthed {
		SendError(&w, r, http.StatusUnauthorized, "unauthorized", "invalid authorization", &map[string]interface{}{
			"route":        "ChangeSnakeRoute",
			"keyIsMissing": key == "",
		})
		return
	}

	input := &SnakeOptions{}
	render.Bind(r, input)
	if input.Randomize {
		input.Color = getRandomColorHex()
		input.Head = getRandomHead()
		input.Tail = getRandomTail()
	}
	if input.Color != "" {
		Config.SnakeColor = input.Color
	}
	if input.Head != "" {
		Config.SnakeHead = input.Head
	}
	if input.Tail != "" {
		Config.SnakeTail = input.Tail
	}
	Send(w, http.StatusOK, input)
}
