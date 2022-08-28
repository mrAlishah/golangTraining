package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello goroutine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("main going to call hello go goroutine")
	go hello(done)
	<-done
	fmt.Println("main received data")

}
