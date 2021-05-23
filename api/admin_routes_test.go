package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestSnakeUpdateRoute(t *testing.T) {
	ConfigSetup()
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc.Encode(map[string]string{})

	// begin with a bad call
	code, _, err := TestAPICall(http.MethodPatch, "/admin/snake", b, ChangeSnakeRoute, "")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, code)

	code, _, err = TestAPICall(http.MethodPatch, "/admin/snake", b, ChangeSnakeRoute, Config.AuthKey)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)

	// we are going to change everything and then back again
	currentSnake := &SnakeOptions{
		Color: Config.SnakeColor,
		Head:  Config.SnakeHead,
		Tail:  Config.SnakeTail,
	}

	// make sure they are new
	newHead := SNAKE_HEAD_CODE_2020_CAFFEINE
	if currentSnake.Head == SNAKE_HEAD_CODE_2020_CAFFEINE {
		newHead = SNAKE_HEAD_CODE_2020_TIGER_KING
	}
	newTail := SNAKE_TAIL_CODE_2020_COFFEE
	if currentSnake.Tail == SNAKE_TAIL_CODE_2020_COFFEE {
		newTail = SNAKE_TAIL_CODE_2020_MOUSE
	}

	input := &SnakeOptions{
		Color: "#000000",
		Head:  newHead,
		Tail:  newTail,
	}
	b.Reset()
	enc.Encode(input)

	code, body, err := TestAPICall(http.MethodPatch, "/admin/snake", b, ChangeSnakeRoute, Config.AuthKey)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)
	m, _ := UnmarshalMap(body)
	ret := &SnakeOptions{}
	err = mapstructure.Decode(m, ret)
	assert.Nil(t, err)
	assert.Equal(t, input.Color, ret.Color)
	assert.Equal(t, input.Head, ret.Head)
	assert.Equal(t, input.Tail, ret.Tail)

	code, body, err = TestAPICall(http.MethodGet, "/", b, GetSnakeRoute, Config.AuthKey)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)
	m, _ = UnmarshalMap(body)
	ret = &SnakeOptions{}
	err = mapstructure.Decode(m, ret)
	assert.Nil(t, err)
	assert.Equal(t, input.Color, ret.Color)
	assert.Equal(t, input.Head, ret.Head)
	assert.Equal(t, input.Tail, ret.Tail)

	// randomize it
	randomInput := map[string]bool{
		"randomize": true,
	}
	b.Reset()
	enc.Encode(randomInput)
	code, _, err = TestAPICall(http.MethodPatch, "/admin/snake", b, ChangeSnakeRoute, Config.AuthKey)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)

	code, body, err = TestAPICall(http.MethodGet, "/", b, GetSnakeRoute, Config.AuthKey)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)
	m, _ = UnmarshalMap(body)
	ret = &SnakeOptions{}
	err = mapstructure.Decode(m, ret)
	assert.Nil(t, err)
	// there is a small chance these end up the same, but a very rare chance that they will ALL be the same
	assert.False(t, ret.Color == input.Color && ret.Head == input.Head && ret.Tail == input.Tail)

	// set it back
	b.Reset()
	enc.Encode(currentSnake)

	code, _, err = TestAPICall(http.MethodPatch, "/admin/snake", b, ChangeSnakeRoute, Config.AuthKey)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)

	code, body, err = TestAPICall(http.MethodGet, "/", b, GetSnakeRoute, Config.AuthKey)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)
	m, _ = UnmarshalMap(body)
	ret = &SnakeOptions{}
	err = mapstructure.Decode(m, ret)
	assert.Nil(t, err)
	assert.Equal(t, currentSnake.Color, ret.Color)
	assert.Equal(t, currentSnake.Head, ret.Head)
	assert.Equal(t, currentSnake.Tail, ret.Tail)
}
