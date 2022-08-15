package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	//when we use this format go fill array with 0
	var b [5]float64
	fmt.Println(b)
	fmt.Println("==================================================")

	//there you are, array with assign number
	var c [3]int
	c[0] = 25
	c[1] = 45
	c[2] = 22

	//in other way to assign number to array
	fmt.Println(c)
	fmt.Println("==================================================")

	d := [3]float64{4.23, 7, 8.9}
	fmt.Println(d)
	//short hand declaration
	e := [3]float64{4.2595635}
	//when we have size 3 and just assign 1 number so other number assign 0 by default
	fmt.Println(e)
	fmt.Println("==================================================")

	f := [...]int{45, 13}

	// ... makes the compiler determine the length
	fmt.Println(f)
	g := [...]string{"Iran", "Argentina", "Brazil", "Iraq", "china"}
	h := g //g copy of h
	h[0] = "USA"
	//h[6]="Germany"=>we cant do that because this is index 6 but we have 5 index

	fmt.Println(g)
	fmt.Println(h)

	//arrays are value type
	fmt.Println("==================================================")

	i := [...]rune{'r', 'f', 'j'}
	fmt.Println("length of i is", len(i))
	//Length of an array =>^^^^^^^^^^^^^use len
	arr := [...]int{123, 33, 93}
	for j := 0; j < len(arr); j++ {
		fmt.Printf("%d*%d=%d\n", arr[j], j, arr[j]*j)

	}
	fmt.Println("==================================================")

	//also we can use range
	l := [...]float64{4.15, 45, 33.1654165165}
	sum := float64(0)     //we use float64 because we need a float var
	for m, o := range l { //range return both the index and the value

		fmt.Printf("%d is element,%.3f is value\n", m, o)
		sum += o //summation af all value

	}
	fmt.Println("this is sum\n", sum)
	fmt.Println("==================================================")
	p := [4]float64{49.5, 34.2205859999999994566, 87, 12.165446}
	for q, r := range p {
		fmt.Printf("\nelement:%d,valu:%.2f", q, r)
	}
	fmt.Println("==================================================")

	arra := [5]float64{32, 0, 45465.564}
	for _, tsi := range arra {
		fmt.Println("\nvalue", tsi) //ignore the index=>use _

	}
	fmt.Println("==================================================")

	// It is possible to create multidimensional arrays.
	s := [2][2]string{
		{"mandar", "masiha"},
		{"idk", "ydk"}, // add comma in the end of sentence for simple rules

	}
	t := [2][1]string{
		{"sisi"},
		{"rabna"},
	}
	fmt.Println(s, t)
	fmt.Println("==================================================")

}
