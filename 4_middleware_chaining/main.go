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

// middleware is a wrapper that takes in a HTTP handler, does something, and
// then calls the original.
type middleware func(http.Handler) http.Handler

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
		// Check the API key.
		if r.Header.Get(apiKeyHeader) != apiKey {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised"))
			return
		}

		// Run the original handler.
		h.ServeHTTP(w, r)
	})
}

// logger takes a HTTP handler and returns a wrapper around this handler.
// The wrapper simply logs the request starting and the request ending.
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("request start")

		// Run the original handler.
		h.ServeHTTP(w, r)

		log.Println("request end")
	})
}

// chain iterates over a list of middlewares, executes them, and returns the
// originally intended handler function.
func chain(h http.Handler, m ...middleware) http.Handler {
	for i := range m {
		h = m[len(m)-1-i](h)
	}

	return h
}

func main() {
	// Set log flags.
	log.SetFlags(log.LstdFlags)

	// Setup our handler for the /hello uri. We now are utilising the chain
	// function which is much more readable and allows us to easily add more
	// middleware wrappers.
	http.Handle("/hello", chain(helloHandler(), logger, checkAPIKey))

	log.Printf("Serving on port %d. Press CTRL+C to cancel.\n", port)

	// Start the server.
	//
	// Note: ListenAndServe should not be used in production as it does not
	// provide any sensible values for read, write or header timeouts.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
