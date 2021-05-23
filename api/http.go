package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var r *chi.Mux

func SetupApp() *chi.Mux {
	ConfigSetup()
	if r != nil {
		return r
	}
	// set up the main routes and middlewares
	r = chi.NewRouter()
	r.Use(middleware.StripSlashes)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(30 * time.Second))

	// set the routes
	r.Get("/", GetSnakeRoute)            // the main index to identify your battlesnake
	r.Get("/status", StatusRequestRoute) // a status route identifying the server
	r.Get("/health", HealthRequestRoute) // used to check if the server is healthy

	// game routes
	r.Post("/start", GameStartRoute)
	r.Post("/move", MoveRequestRoute)
	r.Post("/end", GameEndRoute)

	// admin routes
	r.Patch("/admin/snake", ChangeSnakeRoute)

	return r
}

func checkAuthKey(r *http.Request) (isAuth bool, key string) {
	foundKey := r.Header.Get("Authorization")
	if foundKey == "" {
		return
	}
	parts := strings.Split(foundKey, ":")
	if len(parts) != 2 {
		return
	}
	key = strings.Trim(parts[1], " ")
	isAuth = key == Config.AuthKey
	return
}

// Send is a helper function to normalize what gets sent back and optionally
// log the request
func Send(w http.ResponseWriter, code int, payload interface{}) {
	// TODO: Add logging, potentially based on a flag
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// SendError allows for hooking into an error return so we can do things like trigger an error log
func SendError(w *http.ResponseWriter, r *http.Request, status int, systemCode string, message string, data *map[string]interface{}) {
	if data == nil {
		data = &map[string]interface{}{}
	}
	// TODO: Add logging, potentially based on a flag
	Send(*w, status, map[string]interface{}{
		"code":    systemCode,
		"message": message,
		"data":    data,
	})
}

// TestAPICall allows an easy way to test HTTP end points in unit testing
func TestAPICall(method string, endpoint string, data io.Reader, handler http.HandlerFunc, authorizationKey string) (code int, body *bytes.Buffer, err error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}
	req, err := http.NewRequest(method, endpoint, data)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer: %s", authorizationKey))
	rr := httptest.NewRecorder()

	chi := SetupApp()
	chi.ServeHTTP(rr, req)

	return rr.Code, rr.Body, nil
}

// UnmarshalMap helps to unmarshal the request for the testing calls
func UnmarshalMap(body *bytes.Buffer) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	retBuf := new(bytes.Buffer)
	retBuf.ReadFrom(body)
	err := json.Unmarshal(retBuf.Bytes(), &ret)
	return ret, err
}

// UnmarshalSlice unmarshals a response that is an array in the data
func UnmarshalSlice(body *bytes.Buffer) ([]interface{}, error) {
	ret := []interface{}{}
	retBuf := new(bytes.Buffer)
	retBuf.ReadFrom(body)
	err := json.Unmarshal(retBuf.Bytes(), &ret)
	return ret, err
}
