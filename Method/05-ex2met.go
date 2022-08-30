package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
	pii    float64
}

const pii = math.Pi

func circlearea(c circle) {
	fmt.Printf("area func %.2f\n", (c.radius * c.radius * pii))
}
func (c circle) circlearea() {
	fmt.Printf("method is %.2f\n", (c.radius * c.radius * pii))
}
func main() {
	c := circle{
		radius: 45,
		pii:    pii,
	}
	circlearea(c)
	c.circlearea()
	p := &c
	fmt.Printf("%.2f\n", p)
}
