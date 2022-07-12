package main

import (
	"fmt"
)
type Point struct  {
	x int32
	y int32
}

type Circle struct {
	redius float32
	center *Point
}

func main() {
	p1 := Point{1,2}
	p2 := Point{y: -4,x: 3}
	p3 := &Point{11,0}
	fmt.Println("===> Part 1")
	fmt.Println(p1.x, p1.y)
	fmt.Println(p2)
	fmt.Println(p3)
//----------------------------------------------------------------

	fmt.Println("===> Part 2")
	ChangeX(&p1)
	ChangeX(p3)
	fmt.Println(p1)
	fmt.Println(p3)	
//----------------------------------------------------------------

	fmt.Println("===> Part 3")
	//*p3.x = 444 //incorrect
	(*p3).x = 444 //correct
	fmt.Println(p3)		
	p3.x = 222 //correct
	fmt.Println(p3)		
//----------------------------------------------------------------

	fmt.Println("===> Part 4")
	p4 := &Point{2,3}
	c1 := Circle{1.2, p4}
	c2 := Circle{2.3, &Point{3,4}}
	c3 := Circle{3,&Point{}}
	c3.center.x = 3
	c3.center.y = 44
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}

func ChangeX(p *Point){
	p.x = 100
}