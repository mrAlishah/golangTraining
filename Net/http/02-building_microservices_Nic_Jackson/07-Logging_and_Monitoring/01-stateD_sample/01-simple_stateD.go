package main

import (
	"fmt"
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
	"net/http"
	"runtime"
	"time"
)

// https://github.com/alexcesaro/statsd
func main() {
	c, err := statsd.New() // Connect to the UDP port 8125 by default.
	if err != nil {
		// If nothing is listening on the target port, an error is returned and
		// the returned client does nothing but is still usable. So we can
		// just log the error and go on.
		log.Print(err)
	}
	defer c.Close()

	// Increment a counter.
	c.Increment("foo.counter")

	// Gauge something. goroutine counts
	c.Gauge("num_goroutine", runtime.NumGoroutine())

	// Time something.estimate time.
	t := c.NewTiming()
	p := Ping("http://example.com/")
	fmt.Println(p)
	t.Send("homepage.response_time")

	// It can also be used as a one-liner to easily time a function.
	pingHomepage := func() {
		defer c.NewTiming().Send("homepage.response_time")

		Ping("http://example.com/")
	}
	pingHomepage()

	// Cloning a Client allows using different parameters while still using the
	// same connection.
	// This is way cheaper and more efficient than using New().
	stat := c.Clone(statsd.Prefix("http"), statsd.SampleRate(0.2))
	stat.Increment("view") // Increments http.view

}

func Ping(url string) time.Duration {
	start := time.Now()
	//we make a get request to the URL and calculate the time it took for the request to be completed
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err == nil {
		t := time.Since(start).Round(time.Millisecond)
		return t
	}

	return -1
}
