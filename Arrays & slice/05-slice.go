package main

import (
	"fmt"
	"unsafe"
)

func main() {
  //just for test printf info
  var a int=3
  fmt.Printf("int: %v | size: %d | Type: %T \n",a,unsafe.Sizeof(a),a)

  // an integer array
  array := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
  fmt.Printf("array: %v | size: %d | Len:%d | Cap: %d | Type: %T \n",array,unsafe.Sizeof(array),len(array),cap(array),array)

  // creating slice from an array
  //[Index element,End element]
  slice := array[1 : 2]
  fmt.Printf("slice[1 : 2]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n",slice,unsafe.Sizeof(slice),len(slice),cap(slice),slice)


  slice1 := array[2 : 2]
  fmt.Printf("slice1[2 : 2]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n",slice1,unsafe.Sizeof(slice1),len(slice1),cap(slice1),slice1)

  slice2 := array[3 : 7]
  fmt.Printf("slice2[3 : 7]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n",slice2,unsafe.Sizeof(slice2),len(slice2),cap(slice2),slice2)

  slice3 := array[ : 3]
  fmt.Printf("slice3[ : 3]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n",slice3,unsafe.Sizeof(slice3),len(slice3),cap(slice3),slice3)
  //fmt.Println("slice3[5]",slice3[5]) //ERROR: panic: runtime error: index out of range [5] with length 3

  slice4 := array[6 : ]
  fmt.Printf("slice4[6 : ]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n",slice4,unsafe.Sizeof(slice4),len(slice4),cap(slice4),slice4)

  slice4[0]=9
  fmt.Printf("Pointer slice4[0]=9: %v | size: %d | Len:%d | Cap: %d | Type: %T \n",slice4,unsafe.Sizeof(slice4),len(slice4),cap(slice4),slice4)
  fmt.Printf("Pointer array[6]: %v | size: %d | Len:%d | Cap: %d | Type: %T \n",array,unsafe.Sizeof(array),len(array),cap(array),array)


}