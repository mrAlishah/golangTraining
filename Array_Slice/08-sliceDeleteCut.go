package main

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/fatih/color"
)

var (
	Pred     = color.New(color.FgRed).PrintfFunc()
	Pgreen   = color.New(color.FgGreen).PrintfFunc()
	Pblue    = color.New(color.FgBlue).PrintfFunc()
	Pcyan    = color.New(color.FgCyan).PrintfFunc()
	Pmagneta = color.New(color.FgMagenta).PrintfFunc()
	Pyellow  = color.New(color.FgYellow).PrintfFunc()
)

type ST struct {
	num int
}

func splitter(title string) {
	// Pgreen := color.New(color.FgGreen).PrintfFunc()
	//color.Green(fmt.Sprintf("\n## %s %s \n",title,str))
	l := len(title)
	str := strings.Repeat("-", 130-l)
	Pyellow("\n## %s %s \n", title, str)
}

// reset or Copy slice for other test action
func resetSlice(srcSlice []int, desSlice *[]int) {
	*desSlice = append([]int(nil), srcSlice...)
}

func resetStructSlice(srcSlice []*ST, desSlice *[]*ST) {
	*desSlice = append([]*ST(nil), srcSlice...)
}

func main() {
	//!-----------------------------------------------------------------
	// an integer slice
	slice := make([]int, 0, 0)
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("origin slice:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice, unsafe.Sizeof(slice), cap(slice), len(slice), slice)

	slice3 := make([]int, 0, 0)

	//! Cut -----------------------------------------------------------------
	/*
		a = append(a[:i], a[j:]...)
	*/
	i := 2 //start
	j := 5 //end

	resetSlice(slice, &slice3)
	splitter("cut: a = append(a[:i], a[j:]...)")
	fmt.Printf("slice3:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3, unsafe.Sizeof(slice3), cap(slice3), len(slice3), slice3)
	fmt.Printf("slice3[:2]:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3[:i], unsafe.Sizeof(slice3[:i]), cap(slice3[:i]), len(slice3[:i]), slice3[:i])
	fmt.Printf("slice3[5:]:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3[j:], unsafe.Sizeof(slice3[j:]), cap(slice3[j:]), len(slice3[j:]), slice3[j:])

	slice3 = append(slice3[:i], slice3[j:]...)
	Pgreen("slice3 cutted:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3, unsafe.Sizeof(slice3), cap(slice3), len(slice3), slice3)

	//! Delete -----------------------------------------------------------------
	/*
		a = append(a[:i], a[i+1:]...)
	*/
	resetSlice(slice, &slice3)
	splitter("Delete: a = append(a[:i], a[i+1:]...)")
	fmt.Printf("slice3:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3, unsafe.Sizeof(slice3), cap(slice3), len(slice3), slice3)

	i = 2
	slice3 = append(slice3[:i], slice3[i+1:]...)
	Pgreen("slice3 deleted 2:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3, unsafe.Sizeof(slice3), cap(slice3), len(slice3), slice3)

	//* -----------------------------------------------------------------
	/*
		a = a[:i+copy(a[i:], a[i+1:])]
	*/
	resetSlice(slice, &slice3)
	splitter("Delete: a = a[:i+copy(a[i:], a[i+1:])]")
	fmt.Printf("slice3:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3, unsafe.Sizeof(slice3), cap(slice3), len(slice3), slice3)

	i = 4
	slice3 = slice3[:i+copy(slice3[i:], slice3[i+1:])]
	Pgreen("slice3 deleted 4:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3, unsafe.Sizeof(slice3), cap(slice3), len(slice3), slice3)

	//* -----------------------------------------------------------------
	/*
		a[i] = a[len(a)-1]
		a = a[:len(a)-1]
	*/
	resetSlice(slice, &slice3)
	splitter("Delete without preserving order: a[i] = a[len(a)-1] , a = a[:len(a)-1]")
	fmt.Printf("slice3:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3, unsafe.Sizeof(slice3), cap(slice3), len(slice3), slice3)

	i = 4
	slice3[i] = slice3[len(slice3)-1]
	slice3 = slice3[:len(slice3)-1]
	Pgreen("slice3 deleted 4:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", slice3, unsafe.Sizeof(slice3), cap(slice3), len(slice3), slice3)

	//! -----------------------------------------------------------------

	/*
	   NOTE If the type of the element is a pointer or a struct with pointer fields, which need to be garbage collected,
	   the above implementations of Cut and Delete have a potential memory leak problem:
	   some elements with values are still referenced by slice a and thus can not be collected. The following code can fix this problem
	*/

	/*
		copy(a[i:], a[j:])
		for k, n := len(a)-j+i, len(a); k < n; k++ {
		a[k] = nil // or the zero value of T
		}
		a = a[:len(a)-j+i]
	*/
	structSlice := make([]*ST, 0, 0)
	structSlice = append(structSlice, &ST{1}, &ST{2}, &ST{3}, &ST{4}, &ST{5})
	structSlice3 := make([]*ST, 0, 0)

	resetStructSlice(structSlice, &structSlice3)
	splitter("Cut GC: set nil element , see code")
	fmt.Printf("slice3:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", structSlice3, unsafe.Sizeof(structSlice3), cap(structSlice3), len(structSlice3), structSlice3)

	i = 2 //start
	j = 4 //end
	copy(structSlice3[i:], structSlice3[j:])
	for k, n := len(structSlice3)-j+i, len(structSlice3); k < n; k++ {
		structSlice3[k] = nil // or the zero value of T
	}
	structSlice3 = structSlice3[:len(structSlice3)-j+i]
	Pgreen("slice3 cute 2,4:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", structSlice3, unsafe.Sizeof(structSlice3), cap(structSlice3), len(structSlice3), structSlice3)

	//! -----------------------------------------------------------------
	/*
		if i < len(a)-1 {
			copy(a[i:], a[i+1:])
		  }
		  a[len(a)-1] = nil // or the zero value of T
		  a = a[:len(a)-1]
	*/

	resetStructSlice(structSlice, &structSlice3)
	splitter("Delete GC: set nil element , see code")
	fmt.Printf("slice3:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", structSlice3, unsafe.Sizeof(structSlice3), cap(structSlice3), len(structSlice3), structSlice3)

	i = 2
	if i < len(structSlice3)-1 {
		copy(structSlice3[i:], structSlice3[i+1:])
	}
	structSlice3[len(structSlice3)-1] = nil // or the zero value of T
	structSlice3 = structSlice3[:len(structSlice3)-1]
	Pgreen("slice3 delete 2:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", structSlice3, unsafe.Sizeof(structSlice3), cap(structSlice3), len(structSlice3), structSlice3)

	//! -----------------------------------------------------------------
	/*
		a[i] = a[len(a)-1]
		a[len(a)-1] = nil
		a = a[:len(a)-1]
	*/

	resetStructSlice(structSlice, &structSlice3)
	splitter("Delete without preserving order (GC): set nil element , see code")
	fmt.Printf("slice3:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", structSlice3, unsafe.Sizeof(structSlice3), cap(structSlice3), len(structSlice3), structSlice3)

	i = 2
	structSlice3[i] = structSlice3[len(structSlice3)-1]
	structSlice3[len(structSlice3)-1] = nil
	structSlice3 = structSlice3[:len(structSlice3)-1]
	Pgreen("slice3 delete 2:  Type:%T | size:%d | Cap:%d | Len:%d | %v \n", structSlice3, unsafe.Sizeof(structSlice3), cap(structSlice3), len(structSlice3), structSlice3)

}
