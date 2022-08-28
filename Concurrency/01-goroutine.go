package main

import "fmt"

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello() //start new goroutine
	fmt.Println("main function")
}
