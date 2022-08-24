package main

import "fmt"

type Coordinate struct {
	x, y int
}

func shiftby(x, y int, coord *Coordinate) { //inja 3 ta parameter darim x,y,coord ke coord pointer has
	coord.x += x
	coord.y += y

}
func (coord *Coordinate) shiftby(x, y int) { //here in first ()we have reciver and then we have func
	coord.x += x
	coord.y += y
}
func main() {
	coord := Coordinate{5, 5}
	shiftby(1, 1, &coord)
	fmt.Println(coord)
}
 