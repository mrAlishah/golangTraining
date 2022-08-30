package main

import "fmt"

type Eduction struct {
	major string
	city  string
	score float64
}

func (e Eduction) Edprint() {
	fmt.Printf(" major: %s,city: %s,score: %.2f", e.major, e.city, e.score)
}

type Person struct {
	fname string
	lname string
	Eduction
}

func main() {
	p := Person{
		fname: "fateme",
		lname: "jadidi",
		Eduction: Eduction{
			major: "computer_sience",
			city:  "Semnan",
			score: 17.02,
		},
	}
	p.Edprint() //calling value receiver with a pointer
	fmt.Println("\n")
}
