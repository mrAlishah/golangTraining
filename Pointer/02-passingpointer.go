package main

import "fmt"

func changevalue(cal *int) {
	*cal = 42

}
func main() {
	a := 52
	fmt.Println("value before:", a)
	b := &a
	changevalue(b)
	fmt.Println("value after:", a)
}
