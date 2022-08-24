package main

import "fmt"

func assert(i interface{}) {
	s := i.(int)
	fmt.Println(s)

}
func main() {
	var s interface{} = 56
	assert(s)
}