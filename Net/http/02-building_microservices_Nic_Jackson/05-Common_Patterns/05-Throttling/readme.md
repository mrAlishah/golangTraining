# Throttling Pattern
Throttling is a pattern where you restrict the number of connections that a service can handle, returning an HTTP error code when this threshold has been exceeded. The full source code for this example can be found in the file throttling/limit_handler.go. The middleware pattern for Go is incredibly useful here: what we are going to do is to wrap the handler we would like to call, but before we call the handler itself, we are going to check to see if the server can honor the request. In this example, for simplicity, we are going only to limit the number of concurrent requests that the handler can serve, and we can do this with a simple buffered channel.
### Getting Started with Throttling
Declare a LimitHandler handler. We have two private fields:
-  one holds the number of connections as a buffered channel.
-  the second is the handler we are going to call after we have checked that the system is healthy

```go
type LimitHandler struct {
	connections chan struct{}
	handler     http.Handler
}
```

To create an instance of this object we are going to use the **NewLimitHandler** function. This takes the parameters connection, which is the number of connections we allow to process at any one time and the handler which would be called if successful:
```go
func NewLimitHandler(connections int, next http.Handler) *LimitHandler {
	cons := make(chan struct{}, connections)
	for i := 0; i < connections; i++ {
		cons <- struct{}{}
	}

	return &LimitHandler{
		connections: cons,
		handler:     next,
	}
}
```

This is quite straightforward: we create a buffered channel with the size equal to the number of concurrent connections, and then we fill that ready for use:
```go
func (l *LimitHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	select {
	case <-l.connections:
		l.handler.ServeHTTP(rw, r)
		l.connections <- struct{}{} // release the lock
	default:
		http.Error(rw, "Busy", http.StatusTooManyRequests)
	}
}

Note:If we look at the ServeHTTP method starting at line 29, we have a select statement. The beauty of channel is that we can write a statement like this: if we cannot retrieve an item from the channel then we should return a busy error message to the client.
```

### Testing
Let's take a closer look at the flow through this test:
1. Block at line 109.2. Call handler.ServeHTTP twice concurrently.
3. One ServeHTTP method returns immediately with http.TooManyRequests and decrements the wait group.
4. Call cancel context allowing the one blocking ServeHTTP call to return and decrement the wait group.
5. Perform assertion.

### Running
To build this project.
```
go test
```

Output
```
PASS
ok      throttling      0.490s
```

read this package:<br/>
[rk-boot](https://github.com/rookie-ninja/rk-boot)
[gohalt](https://github.com/1pkg/gohalt)<br/>
[limiter](https://github.com/ulule/limiter)
