package main

import "fmt"

type addres struct {
	city   string
	street string
}
type person struct {
	name   string
	age    int
	Addres addres
}

func main() {
	p := person{
		name: "fateme",
		age:  21,
		Addres: addres{
			city:   "Tehran",
			street: "kamlia",
		},
	}
	fmt.Println(p)
	fmt.Println(p.name)
}
