package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// port is the network port to expose the API on.
	port = 8000

	// apiKeyHeader is the name of the HTTP header which stores the API key.
	apiKeyHeader = "x-api-key"

	// apiKey is the expected api key.
	apiKey = "foo-bar-baz"
)

// helloHandler returns a HTTP handler which writes a friendly message.
func helloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Golang Dorset!"))
	})
}

// checkAPIKey takes a HTTP handler and returns a wrapper around this handler.
// The wrapper checks the validity of an API key.
func checkAPIKey(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(apiKeyHeader) != apiKey {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised"))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	// Setup our handler for the /hello uri. Notice we are wrapping the original
	// handler with our API key checking middleware.
	http.Handle("/hello", checkAPIKey(helloHandler()))

	fmt.Printf("Serving on port %d. Press CTRL+C to cancel.\n", port)

	// Start the server.
	//
	// Note: ListenAndServe should not be used in production as it does not
	// provide any sensible values for read, write or header timeouts.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
