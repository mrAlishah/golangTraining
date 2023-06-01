# Health checks
## Health checks Pattern
* Health checks should be an essential part of your microservices setup.
* Every service should expose a health check endpoint which can be accessed by the consul or another server monitor. Health checks are important as they allow the process responsible for running the application to restart or kill it when it starts to misbehave or fail.
* Of course, you must be incredibly careful with this and not set this too aggressively.

### Usage of Health check
What you record in your health check is entirely your choice. However, I recommend you look at implementing these features:
* Data store connection status (general connection state, connection pool status)
* Current response time (rolling average)
* Current connections
* Bad requests (running average)

### Getting Started with Health check
We are defining two handlers one which deals with our main request at the path / and one used for checking the health at the path /health.<br/>
Exponentially Weighted Moving Average algorithms [link](https://github.com/VividCortex/ewma)
```go
func main() {
	ma = ewma.NewMovingAverage()

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}
```

Implement a handler function with a health check
```go
func mainHandler(rw http.ResponseWriter, r *http.Request) {
	
	startTime := time.Now()

	if !isHealthy() {
		respondServiceUnhealthy(rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Average request time: %f (ms)\n", ma.Value()/1000000)

	duration := time.Now().Sub(startTime)
	ma.Add(float64(duration))
}
```

Create a mutex Lock for setting global varaible.
```go
func respondServiceUnhealthy(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusServiceUnavailable)

	resetMutex.RLock()
	defer resetMutex.RUnlock()

	if !resetting {
		go sleepAndResetAverage()
	}
}
```

Sleep to wait timeout, and reset values
```go
func sleepAndResetAverage() {

	resetMutex.Lock()
	resetting = true
	resetMutex.Unlock()

	time.Sleep(timeout)
	ma = ewma.NewMovingAverage()

	resetMutex.Lock()
	resetting = false
	resetMutex.Unlock()
}

```
