package main

import "fmt"

func solo(lo chan string) {
	for v := 0; v < 4; v++ {
		lo <- "soos mas"

	}
	close(lo)
}
func main() {
	c := make(chan string)
	go solo(c)
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("channel close", ok) //When the value of ok is set to false means the channel is closed
			break
		}
		fmt.Println("channel open", res, ok)
	}
}
