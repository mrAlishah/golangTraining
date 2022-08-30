package main

import "fmt"

type Worker interface {
	Work() //worker interface has one method
}
type Person struct {
	name string
}

func (p Person) Work() {

	fmt.Println(p.name, "is working")

}
func describe(w Worker) {
	fmt.Printf("interface type %T value %v\n", w, w)
}
func main() {
	p := Person{
		name: "fateme",
	}
	var w Worker = p
	describe(w)
	w.Work()
}
