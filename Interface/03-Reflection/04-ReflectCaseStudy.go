package main

import (
	"fmt"
	"reflect"
)

type FFoo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

// https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
func main() {
	greeting := "hello"

	gVal := reflect.ValueOf(greeting)
	// not a pointer so all we can do is read it
	fmt.Println(gVal.Interface())

	gpVal := reflect.ValueOf(&greeting)
	// it’s a pointer, so we can change it, and it changes the underlying variable
	gpVal.Elem().SetString("goodbye")
	fmt.Println(greeting)

	//------------------------------------------------------
	f := FFoo{A: 10, B: "Salutations"}

	fType := reflect.TypeOf(f) //reflect.Type
	fVal := reflect.New(fType) //reflect.Value
	fVal.Elem().Field(0).SetInt(20)
	fVal.Elem().Field(1).SetString("Greetings")
	f2 := fVal.Elem().Interface().(FFoo)
	fmt.Printf("f2: %+v, %d, %s\n", f2, f2.A, f2.B)
	fmt.Printf("f: %+v, %d, %s\n", f, f.A, f.B)
}
