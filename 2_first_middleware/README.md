# Making Middleware Manageable - Part Two
Ok so for part two we've realised that we need to secure the `/hello` endpoint
of our application. We'll do this with an API key.

## Usage
To run the application:
```bash
go run main.go
```

To make a request:
```bash
curl localhost:8000/hello
```
> This will fail because you have no api key.

To make an authenticated request:
```bash
curl -H "x-api-key: foo-bar-baz" localhost:8000/hello
```

## Things to Note
### Middleware
We defined our first middleware, to check API keys, as follows:
```go
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
```
You'll notice that the signature for the function is very similar to that of
`helloHandler` in that it returns a `http.Handler` interface. The key difference
here is that we're also taking in a `http.Handler` interface as a parameter.

This parameter is the original handler function, it is in effect being wrapped
by the middleware. In this case we check the API key header and then run the
original handler function by calling `h.ServeHTTP(w, r)`