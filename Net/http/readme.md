# An Introduction http
Processing HTTP requests with Go is primarily about two things: handlers and servemuxes.

If you’re coming from an MVC-background, you can think of handlers as being a bit like controllers. Generally speaking, they're responsible for carrying out your application logic and writing response headers and bodies.

Whereas a servemux (also known as a router) stores a mapping between the predefined URL paths for your application and the corresponding handlers. Usually you have one servemux for your application containing all your routes.

Go's `net/http` package ships with the simple but effective `http.ServeMux` servemux, plus a few functions to generate common handlers including `http.FileServer()`, `http.NotFoundHandler()` and `http.RedirectHandler()`.