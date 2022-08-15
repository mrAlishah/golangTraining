package main

import (
	"fmt"
)

func main() {

	var srt string
	var age int
	_, err := fmt.Scan(&srt, &age)
	if err != nil {
		return
	}
	fmt.Printf("My name is %s & i am %d years old\n", srt, age)
	//so we can scan for string & int input and printf for output

}
