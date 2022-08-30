package main

import "fmt"

type Describer interface {
	Describe()
}
type Peerson struct {
	name string
	age  int
}

func (p Peerson) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType("Naveen")
	p := Peerson{
		name: "Naveen R",
		age:  25,
	}
	findType(p)
}
