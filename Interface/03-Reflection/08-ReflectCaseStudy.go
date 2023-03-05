package main

import "fmt"

// -------------------------------------
type Fooo struct {
	A int
}

func (f Fooo) Double() int {
	return f.A * 2
}

// -------------------------------------
type Bar struct {
	Fooo
	B int
}

// -------------------------------------
type Doubler interface {
	Double() int
}

func DoDouble(d Doubler) {
	//fmt.Println(d.Double())
	fmt.Printf("%+v\n", d.Double())
}

// -------------------------------------
func main() {
	f := Fooo{10}
	b := Bar{Fooo: f, B: 20}
	fmt.Printf("f: %+v\n", f)
	fmt.Printf("b: %+v\n", b)
	DoDouble(f) // passed in an instance of Foo; it meets the interface, so no surprise here
	DoDouble(b) // passed in an instance of Bar; it works!
}
