package main

import (
	"fmt"
	"math"
)

type Float float64

func (round Float) Addedtwofunc() (li float64) {
	if round < 0 {
		return float64(-round)

	}
	return float64(round)
}
func main() {
	round := Float(-math.Sqrt2)
	fmt.Println(round.Addedtwofunc())
}
