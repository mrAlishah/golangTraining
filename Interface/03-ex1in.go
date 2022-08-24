package main

import (
	"fmt"
)

type Onlight interface {
	lighton() bool
}
type room struct {
	pricelight int
	pricecable int
	color      string
	onOroff    bool
}
type library struct {
	pricelight int
	onOroff    bool
}

func (o room) Onlight() int {
	return o.pricelight + o.pricecable
}
func (o library) Onlight() int {
	return o.pricelight
}

func main() {
	light1 := library{
		pricelight: 45000,
		onOroff:    false,
	}
	light2 := room{
		pricelight: 7500,
		pricecable: 4500,
		color:      "blue",
		onOroff:    false,
	}
	light3 := room{
		pricelight: 750,
		pricecable: 2000,
		color:      "red",
		onOroff:    true,
	}
	fmt.Println(light1)
	fmt.Println(light2)
	fmt.Println(light3)
}
