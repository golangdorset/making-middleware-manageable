# Making Middleware Manageable - Part Three
Ok so for part three we're adding some logging, because what good web application
doesn't create some logs.

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
### Middleware/Handler Order
Our new logging middlware is defined as follows:
```go
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("request start")

		// Run the original handler.
		h.ServeHTTP(w, r)

		log.Println("request end")
	})
}
```
Notice that we're running code before _and_ after the original handler. This is
perfectly valid as `ServeHTTP` ensures the original handler is executed
correctly. Having the option to run code before and after the original handler
gives great flexibility as to what our middleware can do.

### Multiple Middlewares
Now that we have multiple middleware in our application you'll notice that we
wrapped the function calls like so when passing them to `http.Handle`
```go
logger(checkAPIKey(helloHandler()))
```
Because our middleware share the same signature we can easily wrap the functions
like this.

But can you spot the problem with this pattern? We've got 2 middlewares now; what
happens when we have 5 or 10?