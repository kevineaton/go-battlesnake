package api

import "net/http"

// StatusRequestRoute is the http end point for checking the server's status
func StatusRequestRoute(w http.ResponseWriter, r *http.Request) {
	// we don't have any connections to set up, so we can just return a 200

	Send(w, http.StatusOK, map[string]string{
		"status": "up",
	})
}

// HealthRequestRoute is the http end point for checking the server's health
func HealthRequestRoute(w http.ResponseWriter, r *http.Request) {
	// we don't have any connections to set up, so we can just return a 200
	Send(w, http.StatusOK, map[string]string{
		"health_status": "up",
	})
}
