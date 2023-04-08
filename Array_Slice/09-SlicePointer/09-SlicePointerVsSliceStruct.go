package main

import (
	"fmt"
)

func main() {
	n := 10
	perP := ReturnSliceWithPointers(n)
	perS := ReturnSliceWithStructs(n)

	fmt.Println("----------- ReturnSliceWithPointers")
	for i := 0; i < n; i++ {
		fmt.Printf("perP%d = %+v \n", i, *perP[i])
	}

	fmt.Println("----------- ReturnSliceWithPointers")
	for i := 0; i < n; i++ {
		fmt.Printf("perS%d = %+v \n", i, perS[i])
	}
}

type Person struct {
	Age int
}

func ReturnSliceWithPointers(size int) []*Person {
	res := make([]*Person, size)

	for i := 0; i < size; i++ {
		res[i] = &Person{Age: i}
	}

	return res
}

func ReturnSliceWithStructs(size int) []Person {
	res := make([]Person, size)

	for i := 0; i < size; i++ {
		res[i] = Person{Age: i}
	}

	return res
}
