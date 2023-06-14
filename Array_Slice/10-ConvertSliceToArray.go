package main

import "fmt"

// it is possible to covert a slice to an array **pointer**
func main() {
	s := []int{1, 2}
	a := (*[2]int)(s)
	fmt.Printf("%T : %+v\n", s, s)
	fmt.Printf("%T : %+v\n", a, a)

	s = append(s, 3)
	/*
		Compiler Error: first argument to append must be a slice; have a (variable of type *[2]int)
		a = append(a, 3)
	*/

	fmt.Printf("%T : %+v\n", s, s)
	fmt.Printf("%T : %+v\n", a, a)

}

/*
[]int : [1 2]
*[2]int : &[1 2]
[]int : [1 2 3]
*[2]int : &[1 2]
*/
