package main

import "fmt"

type family struct {
	fname string
	lname string
	age   int
	idfn  bool
}
type firend struct {
	likefamily family
}

func main() {
	mehri := family{"mehri", "jadidi", 29, true}
	zahra := family{"zahra", "jadidi", 26, true}
	fateme := family{"fateme", "jadidi", 21, true}
	fmt.Println(mehri, zahra, fateme)
	var (
		ghafor = family{fname: "ghafor", age: 49}
		mahnaz = family{fname: "mahnaz", age: 44}
	)
	fmt.Println(ghafor, mahnaz)
	var sara family
	sara.fname = "Sara"
	sara.age = 21
	fmt.Println(sara)
	sara.idfn = false
	Firend := firend{sara}
	fmt.Println(Firend)
	fmt.Println(Firend.likefamily.lname, "this is so good")
}
