package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
