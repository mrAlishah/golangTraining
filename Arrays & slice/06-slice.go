package main

import (
	"fmt"
	"unsafe"
)

//Show growth slice capecity
func main() {
	  // an integer slice
	  slice := make([]int,0,0)
	  fmt.Printf("slice:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice,unsafe.Sizeof(slice),cap(slice),len(slice),slice)
	
	  slice = append(slice, 1)
	  fmt.Printf("slice:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice,unsafe.Sizeof(slice),cap(slice),len(slice),slice)

	for i:= 2; i<=10 ; i++{
		slice = append(slice, i)
		fmt.Printf("slice%d:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",i,slice,unsafe.Sizeof(slice),cap(slice),len(slice),slice)
	}	
}