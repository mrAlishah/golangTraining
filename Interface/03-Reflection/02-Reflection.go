package main

import (
	"fmt"
	"reflect"
)

// https://www.golangprograms.com/reflection-in-golang.html
func main() {
	//reflectTypeOf()
	//reflectValueOf()
	//reflectField()
	//reflectFieldByIndex()
	reflectFieldByName()
}

// The reflect.TypeOf function returns a value of type reflect.
// Type, which represents the type of the variable passed into the TypeOf function.
func reflectTypeOf() {
	fmt.Println("reflect.TypeOf()------------------------------------------------------")
	v1 := []int{1, 2, 3, 4, 5}
	fmt.Println("v1: ", reflect.TypeOf(v1))

	v2 := "Hello World"
	fmt.Println("v2: ", reflect.TypeOf(v2))

	v3 := 1000
	fmt.Println("v3: ", reflect.TypeOf(v3))

	v4 := map[string]int{"mobile": 10, "laptop": 5}
	fmt.Println("v4: ", reflect.TypeOf(v4))

	v5 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("v5: ", reflect.TypeOf(v5))

	v6 := true
	fmt.Println("v6: ", reflect.TypeOf(v6))
}

// The reflect.ValueOf function to create a reflect.Value instance that represents the value of a variable.
// reflect.Value has methods for finding out information about the value of a variable.
func reflectValueOf() {
	fmt.Println("reflect.ValueOf()------------------------------------------------------")
	v1 := []int{1, 2, 3, 4, 5}
	fmt.Println("v1: ", reflect.ValueOf(v1))

	v2 := "Hello World"
	fmt.Println("v2: ", reflect.ValueOf(v2))

	v3 := 1000
	fmt.Println("v3: ", reflect.ValueOf(v3))
	fmt.Println("v3: ", reflect.ValueOf(&v3))

	v4 := map[string]int{"mobile": 10, "laptop": 5}
	fmt.Println("v4: ", reflect.ValueOf(v4))

	v5 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("v5: ", reflect.ValueOf(v5))

	v6 := true
	fmt.Println("v6: ", reflect.ValueOf(v6))
}

type T struct {
	A int
	B string
	C float64
	D bool
}

// The reflect.Field() Function is used to access the name and type of struct fields.
func reflectField() {
	fmt.Println("reflect.Field()------------------------------------------------------")
	t := T{10, "ABCD", 15.20, true}

	valueT := reflect.ValueOf(t)
	typeT := reflect.TypeOf(t)
	for i := 0; i < typeT.NumField(); i++ {
		field := typeT.Field(i)
		value := valueT.Field(i)
		fmt.Println(field.Name, field.Type, value)
	}
}

type First struct {
	A int
	B string
	C float64
}

type Second struct {
	First
	D bool
}

// The reflect.FieldByIndex() Function is used to get the nested field corresponding to index.
func reflectFieldByIndex() {
	fmt.Println("reflect.FieldByIndex()------------------------------------------------------")
	s := Second{First: First{10, "ABCD", 15.20}, D: true}
	t := reflect.TypeOf(s)

	fmt.Printf("%v\n", t.FieldByIndex([]int{0}))
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%+v\n", t.FieldByIndex([]int{0, 1}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 2}))
	fmt.Printf("%v\n", t.FieldByIndex([]int{1}))
}

type TT struct {
	A int
	B string
	C float64
}

func reflectFieldByName() {
	fmt.Println("reflect.FieldByName()------------------------------------------------------")
	s := TT{10, "ABCD", 15.20}
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("A"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("B"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("C"))

	reflect.ValueOf(&s).Elem().FieldByName("A").SetInt(50)
	reflect.ValueOf(&s).Elem().FieldByName("B").SetString("Test")
	reflect.ValueOf(&s).Elem().FieldByName("C").SetFloat(5.5)

	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("A"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("B"))
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("C"))
}
