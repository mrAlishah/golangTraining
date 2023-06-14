package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//just for test printf info
	var a1 int = 3
	fmt.Printf("int: %v | size: %d | Type: %T \n", a1, unsafe.Sizeof(a1), a1)

	// an integer array
	array := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("array: %v | size: %d | Len:%d | Cap: %d | Type: %T \n", array, unsafe.Sizeof(array), len(array), cap(array), array)

	// creating slice from an array
	//[Index element,End element]
	slice := array[1:2]
	fmt.Printf("slice[1 : 2]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n", slice, unsafe.Sizeof(slice), len(slice), cap(slice), slice)

	slice1 := array[2:2]
	fmt.Printf("slice1[2 : 2]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n", slice1, unsafe.Sizeof(slice1), len(slice1), cap(slice1), slice1)

	slice2 := array[3:7]
	fmt.Printf("slice2[3 : 7]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n", slice2, unsafe.Sizeof(slice2), len(slice2), cap(slice2), slice2)

	slice3 := array[:3]
	fmt.Printf("slice3[ : 3]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n", slice3, unsafe.Sizeof(slice3), len(slice3), cap(slice3), slice3)
	//fmt.Println("slice3[5]",slice3[5]) //ERROR: panic: runtime error: index out of range [5] with length 3

	slice4 := array[6:]
	fmt.Printf("slice4[6 : ]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n", slice4, unsafe.Sizeof(slice4), len(slice4), cap(slice4), slice4)

	slice4[0] = 9
	fmt.Printf("Pointer slice4[0]=9: %v | size: %d | Len:%d | Cap: %d | Type: %T \n", slice4, unsafe.Sizeof(slice4), len(slice4), cap(slice4), slice4)
	fmt.Printf("Pointer array[6]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n", array, unsafe.Sizeof(array), len(array), cap(array), array)

	//! about references-----------------------------------------------------------------
	// array: var slice=array is pointer | slice(array):=array is clone
	// slice: var slice=slice is pointer | slice:=slice is pointer

	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //this is array type int
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}    // this is a slice

	fmt.Println("\n-----------")
	var s1 []int = a[3:6]                       //made slice 's1'
	fmt.Println("s1:made slice from array", s1) //=>[3 4 5]

	fmt.Println("\n-----------")
	var s2 []int = s[3:6]                       //made slice 's2', this is pointer
	fmt.Println("s2:made slice from slice", s2) //=>[3 4 5]

	fmt.Println("\n-----------")
	s1[1] = 44                                                //slice is by reference or pointer then if you need to not change so clone it
	fmt.Println("s1[1]=44:we wanted to change slice s1 ", s1) //=>[3 44 5]
	fmt.Println("a:array a changed", a)                       //=>[0 1 2 3 44 5 6 7 8 9]

	fmt.Println("\n-----------")
	s2[1] = 44                                                //slice is by reference or pointer then if you need to not change so clone it
	fmt.Println("s2[1]=44:we wanted to change slice s2 ", s2) //=>[3 44 5]
	fmt.Println("s:slice a changed", s)                       //=>[0 1 2 3 44 5 6 7 8 9]

	fmt.Println("\n-----------")
	s3 := a //this is copy of array a
	s3[1] = 11
	fmt.Println("s3[1]=11:slice s3 was clone a", s3) //=>[0 11 2 3 44 5 6 7 8 9]
	fmt.Println("a:array a not changed", a)          //=>[0 1 2 3 44 5 6 7 8 9]

	fmt.Println("\n-----------")
	s4 := s //this is pointer of slice s
	s4[1] = 11
	fmt.Println("s4[1]=11:slice s4 is pointer s", s4) //=>[0 11 2 3 44 5 6 7 8 9]
	fmt.Println("s:slice s was changed", s)           //=>[0 11 2 3 44 5 6 7 8 9]

}
