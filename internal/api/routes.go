package api

import (
	"encoding/json"
	"net/http"
)

// test run
func TestHandPingPong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m := map[string]string{
		"status": "ok",
	}
	json.NewEncoder(w).Encode(m)
}
