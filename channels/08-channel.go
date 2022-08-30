package main

import "fmt"

func main() {
	who := make(chan string, 2)
	who <- "naveen"
	who <- "gina"
	who <- "steve"
	fmt.Println(<-who)
	fmt.Println(<-who)
}
