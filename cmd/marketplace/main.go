package main

import (
	"net/http"

	"github.com/cdlinkin/marketplace/internal/api"
)

// test run
func main() {
	http.HandleFunc("/ping", api.TestHandPingPong)

	http.ListenAndServe(":8080", nil)
}
