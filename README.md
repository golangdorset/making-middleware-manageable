# Making Middleware Manageable
This is a repository of companion code examples for the Golang Dorset lightning
talk on HTTP middleware.

## Summary
The repository is comprise of 4 parts:

1. [Part One](1_no_middleware) - A simple HTTP server with an API whose /hello endpoint returns a friendly message.
2. [Part Two](2_first_middleware) - Extending the HTTP server with a middleware to secure /hello with an API key.
3. [Part Three](3_second_middleware) - Further extending the HTTP server with another middleware to add request logging.
4. [Part Four](4_middleware_chaining) - Tidying up the middleware implementation by chaining.