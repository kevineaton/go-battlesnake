package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestSnakeInfoRoute(t *testing.T) {
	ConfigSetup()
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc.Encode(map[string]string{})

	// begin with a bad call
	code, body, err := TestAPICall(http.MethodGet, "/", b, GetSnakeRoute)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code)

	// check that the fields line up
	m, err := UnmarshalMap(body)
	assert.Nil(t, err)
	result := map[string]string{}
	err = mapstructure.Decode(m, &result)
	assert.Nil(t, err)
	assert.Equal(t, "1", result["apiversion"])
	assert.Equal(t, Config.Author, result["author"])
	assert.Equal(t, Config.SnakeColor, result["color"])
	assert.Equal(t, Config.SnakeHead, result["head"])
	assert.Equal(t, Config.SnakeTail, result["tail"])
	assert.Equal(t, Config.Version, result["version"])

}

func TestGameStartRoute(t *testing.T) {
	ConfigSetup()
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc.Encode(map[string]string{})

	// begin with a bad call
	code, _, err := TestAPICall(http.MethodPost, "/start", b, GameStartRoute)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, code, "bad body sent")

	// now pass in blank but valid fields; should still error
	input := &GameRequest{}
	b.Reset()
	enc.Encode(input)
	code, _, err = TestAPICall(http.MethodPost, "/start", b, GameStartRoute)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, code, "empty body sent")

}
