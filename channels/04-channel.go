package main

import (
	"fmt"
	"time"
)

func routine(vim chan int) {
	fmt.Println(234 + <-vim) //here we do not send or received 234 first program wait for main after 23 received to channel then received 234

}
func main() {
	vim := make(chan int)
	go routine(vim)
	time.Sleep(5 * time.Second)
	vim <- 23
}
