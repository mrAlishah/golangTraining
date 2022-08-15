package main

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/fatih/color"
)

var (
	Pred   = color.New(color.FgRed).PrintfFunc()
	Pgreen = color.New(color.FgGreen).PrintfFunc()
	Pblue  = color.New(color.FgBlue).PrintfFunc()
	Pcyan  = color.New(color.FgCyan).PrintfFunc()
	Pmagneta  = color.New(color.FgMagenta).PrintfFunc()
	Pyellow  = color.New(color.FgYellow).PrintfFunc()
)

func splitter(title string){
	// Pgreen := color.New(color.FgGreen).PrintfFunc()
	//color.Green(fmt.Sprintf("\n## %s %s \n",title,str))
	l := len(title)
	str := strings.Repeat("-",130-l)
	Pyellow("\n## %s %s \n",title,str);	
}


//Slice Tricks, look at:
//https://ueokande.github.io/go-slice-tricks/
//https://github.com/golang/go/wiki/SliceTricks#expand
func main(){
//!-----------------------------------------------------------------
	  // an integer slice
	  slice := make([]int,0,0)
	  slice=append(slice, 1,2,3,4,5,6,7,8,9,10)
	  fmt.Printf("origin slice:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice,unsafe.Sizeof(slice),cap(slice),len(slice),slice)

//!-----------------------------------------------------------------
	  //growth cap*2 : cap 10 => cap 20
	  slice=append(slice,11)
	  splitter("growth cap")
	  fmt.Printf("growth slice:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice,unsafe.Sizeof(slice),cap(slice),len(slice),slice)


//! Copy -----------------------------------------------------------------
	  /*
	      b = make([]T, len(a))
	      copy(b, a)
	  */
	  //1-best way to copy slice : Copy and Trim(cap 20) => cap 11
	  slice1 := make([]int,len(slice))
	  copy(slice1,slice)

	  slice1[3]=0
	  splitter("copy: b = make([]T, len(a)) , copy(b, a) ./. slice1[3]=0 : updatebyValue after copy.have been 2 seprate value")
	  fmt.Printf("slice:   Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice,unsafe.Sizeof(slice),cap(slice),len(slice),slice)
	  Pgreen("slice1:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice1,unsafe.Sizeof(slice1),cap(slice1),len(slice1),slice1)
//*-----------------------------------------------------------------
	  /*
		  b = append([]T(nil), a...)
	  */

	// These two are often a little slower than the above one,
	// but they would be more efficient if there are more
	// elements to be appended to b after copying.	  
	//2-best way to copy slice : Copy and Trim(cap 20) => cap 12
	slice2 := append([]int(nil), slice...)

	slice2[4]=0
	splitter("copy: b = append([]T(nil), a...) ./. slice2[4]=0 ")
	fmt.Printf("slice:   Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice,unsafe.Sizeof(slice),cap(slice),len(slice),slice)
	fmt.Printf("slice1:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice1,unsafe.Sizeof(slice1),cap(slice1),len(slice1),slice1)
	Pgreen("slice2:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice2,unsafe.Sizeof(slice2),cap(slice2),len(slice2),slice2)

//!*-----------------------------------------------------------------
	/*
		b = append(a[:0:0], a...)
	*/
	//3-best way to copy slice : Copy and Trim(cap 20) => cap 12
	slice3 := append(slice[:0:0], slice...)

	splitter("copy: b = append(a[:0:0], a...)  ./. slice3[:0:0] = []")
	fmt.Printf("slice:   Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice,unsafe.Sizeof(slice),cap(slice),len(slice),slice)
	fmt.Printf("slice1:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice1,unsafe.Sizeof(slice1),cap(slice1),len(slice1),slice1)
	fmt.Printf("slice2:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice2,unsafe.Sizeof(slice2),cap(slice2),len(slice2),slice2)
	Pgreen("slice3:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice3,unsafe.Sizeof(slice3),cap(slice3),len(slice3),slice3)
//!*-----------------------------------------------------------------
	/*
		b = append(make([]T, 0, len(a)), a...)
	*/
	// This one-line implementation is equivalent to the above
	// two-line make+copy implementation logically. But it is
	// actually a bit slower (as of Go toolchain v1.16).
	slice4 := append(make([]int, 0, len(slice)), slice...)

	splitter("copy: b = append(make([]T, 0, len(a)), a...)")
	fmt.Printf("slice:   Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice,unsafe.Sizeof(slice),cap(slice),len(slice),slice)
	fmt.Printf("slice1:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice1,unsafe.Sizeof(slice1),cap(slice1),len(slice1),slice1)
	fmt.Printf("slice2:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice2,unsafe.Sizeof(slice2),cap(slice2),len(slice2),slice2)
	fmt.Printf("slice3:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice3,unsafe.Sizeof(slice3),cap(slice3),len(slice3),slice3)
	Pgreen("slice4:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n",slice4,unsafe.Sizeof(slice4),cap(slice4),len(slice4),slice4)

}