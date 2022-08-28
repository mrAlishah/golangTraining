package main

import "fmt"

func main() {
	bie := make(chan string, 5)
	bie <- "advertising"
	bie <- "attend"
	bie <- "action"
	bie <- "active"
	fmt.Println("cap is", cap(bie))
	fmt.Println("len is", len(bie))
	fmt.Println("read value", <-bie, <-bie)
	fmt.Println("new len is ", len(bie))

}
