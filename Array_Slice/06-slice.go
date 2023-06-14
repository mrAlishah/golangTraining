package main

import (
	"fmt"
	"unsafe"
)

// Show growth slice capecity
func main() {
	// an integer slice
	slice := make([]int, 0, 0)
	fmt.Printf("slice:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice, unsafe.Sizeof(slice), cap(slice), len(slice), slice)

	slice = append(slice, 1)
	fmt.Printf("slice:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice, unsafe.Sizeof(slice), cap(slice), len(slice), slice)

	for i := 2; i <= 10; i++ {
		slice = append(slice, i)
		fmt.Printf("slice%d:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", i, slice, unsafe.Sizeof(slice), cap(slice), len(slice), slice)
	}

	//! Zero slice-----------------------------------------------------------------
	//The zero value of a slice type is nil.A nil slice has length and capacity 0. It is possible to append values to a nil slice using the append function.
	fmt.Println("----------- Zero slice")
	var slice1 []int
	fmt.Printf("slice1:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice1, unsafe.Sizeof(slice1), cap(slice1), len(slice1), slice1)

	if slice1 == nil {
		slice1 = append(slice1, 0, 1, 2, 3, 4)
		fmt.Println("append into a slice1:", slice1)
	}

	s1 := []int{5, 6, 7}
	slice1 = append(slice1, s1...)
	fmt.Println("append as variadic into a slice1:", slice1)
}
