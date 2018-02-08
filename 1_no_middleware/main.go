package main

import (
	"fmt"
	"log"
	"net/http"
)

// port is the network port to expose the API on.
const port = 8000

// helloHandler returns a HTTP handler which writes a friendly message.
func helloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Golang Dorset!"))
	})
}

func main() {
	// Setup our handler for the /hello uri.
	http.Handle("/hello", helloHandler())

	fmt.Printf("Serving on port %d. Press CTRL+C to cancel.\n", port)

	// Start the server.
	//
	// Note: ListenAndServe should not be used in production as it does not
	// provide any sensible values for read, write or header timeouts.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
