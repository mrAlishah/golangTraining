package main

import (
	"fmt"
)

func main() {
	a := [...]int16{45, 3, 13, 14, 8, 21, 31, 78, 10, 36} //made array 'a' type int
	var b []int16 = a[1:4]                                //made slice 'b'
	fmt.Println("made slice from a", b)
	c := [3]string{"Multidimensional", "Readability", "reuse"} //made array c=>type string
	var d []string = c[1:3]                                    //made slice 'd'
	fmt.Println("made slice from c", d)
	disis := a[2:6]                  //made slice 'disis'
	fmt.Println("array a before", a) //print disis before change
	for i := range disis {
		disis[i]++
	}
	fmt.Println("array a after", disis) //print disis after change
	dil := a
	dil[1] = 150
	fmt.Println("dil is", dil) //dil is a copy of a but index 1 is 150

	fmt.Println("=====================================================")
	//capacity
	e := [...]string{"parham", "mahdi", "mostafa", "shiva", "shima", "saeid", "fateme"}
	f := e[2:6]
	fmt.Printf("length e: %d,capacity e: %d ", len(e), cap(e))
	fmt.Printf("length f: %d,capacity f:%d \n", len(f), cap(f))
	fmt.Println("=====================================================")
	g := [...]string{"a1", "a2", "a3"}
	var h []string = g[1:2]
	fmt.Println("h is a slice", h)
	fmt.Println("=====================================================")
	FJ := make([]int, 5, 6)
	FJ[0] = 20
	FJ[1] = 455

	fmt.Println("FJ slice:", FJ) //so use make for making new slice
	fmt.Println("=====================================================")
	fmt.Println("example number 45")
	DS := []string{"41", "12", "86", "73"} //if put ... in [] return error because we can appened to slice not arrays
	fmt.Printf("type:%T,len is:%d,cap is:%d\n", DS, len(DS), cap(DS))
	DS = append(DS, "20")

	fmt.Printf("len is:%d,cap is:%d\n", len(DS), cap(DS))

	fmt.Println("=================================================================================")
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	fmt.Println("=====================================================")

	//The zero value of a slice type is nil.A nil slice has length and capacity 0. It is possible to append values to a nil slice using the append function.
	var sens []string
	if sens == nil {
		fmt.Println("slice is nil going to append")
		sens = append(sens, "this", "1", "is", "big", "one")
		fmt.Println("sens contents", sens)
		Hs := []string{"do", "it", "by", "yourself"}
		Js := []string{"one", "two"}
		VS := append(Hs, Js...)
		fmt.Println("this is vs:", VS)
	}

}
