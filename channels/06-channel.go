package main

import "fmt"

func main() {
	valaha := make(chan string, 3)
	go func() {
		valaha <- "F"
		valaha <- "A"
		valaha <- "T"
		valaha <- "E"
		valaha <- "M"
		valaha <- "E"
		close(valaha)
	}()

	/* here we use go routine so we do not have deadlock
	but in ex number 8 we have deadlock because we have cap 2 and not had not goroutine and concurrency */

	for va := range valaha {
		fmt.Println(va, len(valaha), cap(valaha))
	}
}
