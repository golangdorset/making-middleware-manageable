# Making Middleware Manageable - Part One
To kick things off we've created a simple HTTP server with a single endpoint of
`/hello`. A request to this endpoint returns a friendly greeting.

## Usage
To run the application:
```bash
go run main.go
```

To make a request:
```bash
curl localhost:8000/hello
```

## Things to Note
### HTTP Handlers
The function which handles requests to `/hello` is defined as:
```go
func helloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Golang Dorset!"))
	})
}
```
This is a very simple Go function which returns `http.Handler`. The return type
here is an interface provided by the net/http package, it is simply defined as
something that response to a HTTP request. The interface has one method:
`ServeHTTP(ResponseWriter, *Request)`. See the [godoc](https://godoc.org/net/http#Handler)
for more information.

Note that the body of `helloHandler` actually returns a `http.HandlerFunc`. This
is because `http.Handler` is just an interface, it cannot do anything by itself.
Whereas `http.Handlerfunc` is an implementation of that interface that allows the
returning function to respond to a HTTP request. See the definition of how it
is implementation `ServeHTTP`:
```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```
It simply takes the response (`w`) and request (`r`) parameters and runs the
original function (`f`).