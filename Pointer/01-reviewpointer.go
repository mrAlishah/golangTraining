package main

import "fmt"

func main() {
	b := 34
	var a *int = &b //& is used to get the address of var and * assigning the address of b to a
	fmt.Println("address of b is \n", a)
	fmt.Printf("type of a is %T\n", a)
	fmt.Printf("type of b is %T\n", b)
	//Creating pointers using the new function
	c := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *c, c, c)
	*c = 78
	fmt.Println("new size:", *c)

}
