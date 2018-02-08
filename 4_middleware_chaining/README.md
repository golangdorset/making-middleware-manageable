# Making Middleware Manageable - Part Four
For the final part we're going to tidy up how we use multiple middlewares.

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
### The Middleware Type
Since part two we've been using the same signature for our middlewares. You'll
notice here that we've created a custom type for this signature:
```go
type middleware func(http.Handler) http.Handler
```
This allows us to create variables and function arguments that will always match
the signature we need for middleware.

### Middleware Chaining
You'll notice that we are now using the `chain` function to wrap multiple
middlewares around the http handler:
```go
chain(helloHandler(), logger, checkAPIKey)
```
This pattern is much more readable from the get go and will continue to be so
when adding more and more middleware.

What is also worth noting here is that the `chain` function takes the middlewares
as a variadic argument:
```go
func chain(h http.Handler, m ...middleware) http.Handler {}
```
> The ... syntax indicates it is variadic

The variadic argument is automatically expanded into a slice by Go allowing the
`chain` function to iterate over the middlewares. This can also work in the
reverse direction, if we had say 5 middlewares we could define them as a slice
and then expand the slice into a variadic argument like so:
```go
middlewares := []middleware{
	logger,
	checkAPIKey,
	foo,
	bar,
	baz
}

chain(helloHandler(), middlewares...)
```
> Pretty clean right?