package main

import "fmt"

func main() {

	//Type switches
	//https://go.dev/tour/methods/16
	var x interface{} = "foo"

	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil") // here v has type interface{}
	case int:
		fmt.Println("x is", v) // here v has type int
	case bool, string:
		fmt.Println("x is bool or string") // here v has type interface{}
	default:
		fmt.Println("type unknown") // here v has type interface{}
	}
	//--------------------------------
	//https://go.dev/tour/methods/15
	//This statement asserts that the interface value i holds the concrete type T and assigns
	//the underlying T value to the variable t.
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	//--------------------------------
	//If i holds a T, then t will be the underlying value and ok will be true.
	s, ok := i.(string)
	fmt.Println(s, ok)

	//If not, ok will be false and t will be the zero value of type T, and no panic occurs.
	f, ok := i.(float64)
	fmt.Println(f, ok)

	//-------------------------------
	//If i does not hold a T, the statement will trigger a panic.
	f = i.(float64) // panic
	fmt.Println(f)

}
