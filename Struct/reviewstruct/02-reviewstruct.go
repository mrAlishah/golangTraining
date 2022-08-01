package main

import "fmt"

type coordinate struct {
	x, y int
}
type rectangle struct {
	a coordinate
	b coordinate
}

func width(rect rectangle) int {
	return (rect.b.x - rect.a.x)
}
func length(rect rectangle) int {
	return (rect.a.y - rect.b.y)
}
func area(rect rectangle) int {
	return length(rect) * width(rect)

}
func perimeter(rect rectangle) int {
	return (length(rect) * 2) + (width(rect) * 2)
}

func printinfo(rect rectangle) {
	fmt.Println("Area is ", area(rect))
	fmt.Println("Perimeter is ", perimeter(rect))
}

func main() {
	rect := rectangle{a: coordinate{0, 7}, b: coordinate{10, 0}}
	printinfo(rect)
	rect.a.y *= 2
	rect.b.x *= 2
}
