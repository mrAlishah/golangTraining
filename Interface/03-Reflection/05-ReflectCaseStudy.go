package main

import (
	"fmt"
	"reflect"
)

/*
Making Without Make
you can also use reflection to make instances that normally require the make function.
You can make a slice, map, or channel using the reflect.MakeSlice, reflect.MakeMap, and reflect.MakeChan functions.
In all cases, you supply a reflect.Type and get back a reflect.Value that you can manipulate with reflection,
or that you can assign back to a standard variable.
*/
func main() {
	// declaring these vars, so I can make a reflect.Type
	//this is "Type" as defined
	intSlice := make([]int, 0)
	mapStringInt := make(map[string]int)
	fmt.Println("1-make([]int, 0): ", intSlice)
	fmt.Println("1-make(map[string]int): ", mapStringInt)

	// here are the reflect.Types
	// this is "reflect.Type" as defined
	sliceType := reflect.TypeOf(intSlice)
	mapType := reflect.TypeOf(mapStringInt)
	fmt.Println("2-reflect.TypeOf(intSlice): ", sliceType)
	fmt.Println("2-reflect.TypeOf(mapStringInt): ", mapType)

	// and here are the new values that we are making
	// this is "reflect.Value" as defined
	intSliceReflect := reflect.MakeSlice(sliceType, 0, 0) // argument needs reflect.Type
	mapReflect := reflect.MakeMap(mapType)                // argument needs reflect.Type
	fmt.Println("3-reflect.MakeSlice(sliceType, 0, 0): ", intSliceReflect)
	fmt.Println("3-reflect.MakeMap(mapType): ", mapReflect)

	// and here we are using them
	v := 10
	rv := reflect.ValueOf(v)
	intSliceReflect = reflect.Append(intSliceReflect, rv) //we can not use v as int instead of "reflect.Value"
	intSlice2 := intSliceReflect.Interface().([]int)      //assertion needs to be Interface().(Type)
	fmt.Println(intSlice2)

	k := "hello"
	rk := reflect.ValueOf(k)
	mapReflect.SetMapIndex(rk, rv)
	mapStringInt2 := mapReflect.Interface().(map[string]int)
	fmt.Println(mapStringInt2)
}
