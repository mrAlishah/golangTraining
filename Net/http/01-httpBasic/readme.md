# An Introduction http
Processing HTTP requests with Go is primarily about two things: handlers and servemuxes.

If you’re coming from an MVC-background, you can think of handlers as being a bit like controllers. Generally speaking, they're responsible for carrying out your application logic and writing response headers and bodies.

Whereas a servemux (also known as a router) stores a mapping between the predefined URL paths for your application and the corresponding handlers. Usually you have one servemux for your application containing all your routes.

Go's `net/http` package ships with the simple but effective `http.ServeMux` servemux, plus a few functions to generate common handlers including `http.FileServer()`, `http.NotFoundHandler()` and `http.RedirectHandler()`.

The eagle-eyed of you might have also noticed something interesting: the signature for the `http.ListenAndServe()` function is `ListenAndServe(addr string, handler Handler)`, but we passed a `servemux` as the second parameter.

We were able to do this because the `http.ServeMux` type has a `ServeHTTP() method`, meaning that it too satisfies the `http.Handler interface`.

## Middleware
When you're building a web application there's probably some shared functionality that you want to run for many (or even all) HTTP requests. You might want to log every request, gzip every response, or check a cache before doing some expensive processing.

One way of organising this shared functionality is to set it up as middleware — self-contained code which independently acts on a request before or after your normal application handlers. In Go a common place to use middleware is between a router (such as http.ServeMux) and your application handlers, so that the flow of control for a HTTP request looks 
- `Router → Middleware Handler → Application Handler
`

Making and using middleware in Go is fundamentally simple. We want to:

- Implement our middleware so that it satisfies the `http.Handler` interface.
- Build up a chain of handlers containing both our middleware handler and our normal application handler, which we can register with a router.

Hopefully you're already familiar with the following pattern for constructing a handler:

```go
func messageHandler(message string) http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte(message)
})
}
```

In this code we put our handler logic (a simple w.Write()) in an anonymous function which closes-over the message variable to form a closure. We then convert the closure to a handler with the http.HandlerFunc() adapter, and then return it.
We can use this same pattern to help us create a chain of handlers. Instead of passing a string into the closure (like above) we could pass the next handler in the chain as a variable, and then transfer control to this next handler by calling it's ServeHTTP() method.

This gives us a complete pattern for constructing middleware:

```go
func exampleMiddleware(next http.Handler) http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// Our middleware logic goes here...
next.ServeHTTP(w, r)
})
}
```

You'll notice that this middleware function has a func(http.Handler) http.Handler signature. It accepts a handler as a parameter and returns a handler. This is useful for two reasons:

- Because it returns a handler we can register the middleware function directly with the standard http.ServeMux router in Go's net/http package.
- We can create an arbitrarily long handler chain by nesting middleware functions inside each other.
```go
mux := http.NewServeMux()
mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
```