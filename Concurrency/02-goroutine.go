package main

import (
	"fmt"
	"time"
)

func helllo() {
	fmt.Println("Hello world goroutine")

}
func main() {
	go helllo()
	time.Sleep(1 * time.Second) //main goroutine is put to sleep for 1 second.
	fmt.Println("main function")
}
