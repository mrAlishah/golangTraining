package main

import (
	"fmt"
	"reflect"
)

/*
I Want a New Struct
There’s one more thing that you can make using reflection in Go. You can make brand-new structs at runtime
by passing a slice of reflect.StructField instances to the reflect.StructOf function.
This one is a bit weird; we are making a new type, but we don’t have a name for it,
so you can’t really turn it back into a “normal” variable. You can create a new instance and use Interface()
to put the value into a variable of type interface{}, but if you want to set any values on it, you need to use reflection.
*/
func MakeStruct(vals ...interface{}) interface{} {
	var sfs []reflect.StructField
	for k, v := range vals {
		t := reflect.TypeOf(v)
		sf := reflect.StructField{
			Name: fmt.Sprintf("F%d", (k + 1)),
			Type: t,
		}
		sfs = append(sfs, sf)
	}
	st := reflect.StructOf(sfs)
	so := reflect.New(st)
	return so.Interface()
}

func main() {
	s := MakeStruct(0, "hello", []int{}) //not changed in initial
	// this returned a pointer to a struct with 3 fields:
	// an int, a string, and a slice of ints
	// but you can’t actually use any of these fields
	// directly in the code; you have to reflect them
	sr := reflect.ValueOf(s)
	fmt.Println("struct:", sr)

	// getting and setting the int field
	fmt.Println("\nStructField[0]: ", sr.Elem().Field(0).Interface())
	sr.Elem().Field(0).SetInt(20)
	fmt.Println("After changed: StructField[0]: ", sr.Elem().Field(0).Interface())

	// getting and setting the string field
	fmt.Println("\nStructField[1]: ", sr.Elem().Field(1).Interface())
	sr.Elem().Field(1).SetString("reflect me")
	fmt.Println("After changed: StructField[1]: ", sr.Elem().Field(1).Interface())

	// getting and setting the []int field
	fmt.Println("\nStructField[2]: ", sr.Elem().Field(2).Interface())
	v := []int{1, 2, 3}
	rv := reflect.ValueOf(v)
	sr.Elem().Field(2).Set(rv)
	fmt.Println("After changed: StructField[2]: ", sr.Elem().Field(2).Interface())
	fmt.Println("\nstruct:", sr)
}
