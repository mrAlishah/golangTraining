package main

import (
	"fmt"
)

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}
func main() {
	var ssd string
	fmt.Scanln(&ssd)
	findType(ssd)
	var bno int
	fmt.Scanln(&bno)
	findType(bno)

}
