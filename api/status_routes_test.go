package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestStatusAndHealthRoutes(t *testing.T) {
	ConfigSetup()
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc.Encode(map[string]string{})

	// right now, these are very simple as there's no external caching to set up
	code, body, err := TestAPICall(http.MethodGet, "/status", b, StatusRequestRoute)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code, fmt.Sprintf("body response was %+v", body))
	m, _ := UnmarshalMap(body)
	ret := map[string]string{}
	err = mapstructure.Decode(m, &ret)
	assert.Nil(t, err)
	status := ret["status"]
	assert.Equal(t, "up", status)

	// now the health
	code, body, err = TestAPICall(http.MethodGet, "/health", b, StatusRequestRoute)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, code, fmt.Sprintf("body response was %+v", body))
	m, _ = UnmarshalMap(body)
	ret = map[string]string{}
	err = mapstructure.Decode(m, &ret)
	assert.Nil(t, err)
	status = ret["health_status"]
	assert.Equal(t, "up", status)

}
