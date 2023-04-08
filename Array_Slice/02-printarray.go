package main

import "fmt"

func Printarray(a [3][2]float64) {
	for _, b := range a {
		for _, c := range b {
			fmt.Printf("%.3f ", c)
		}
		fmt.Printf("\n")

	}

}
func main() {
	a := [3][2]float64{
		{24, 14.2},
		{41, 15.23},
		{11.56, 56.666},
	}
	Printarray(a)
}
