package main

import "fmt"

func main() {
	fmt.Println("=>part1")

	var dayy int
	var monthh string
	var yearr int
	fmt.Scanf("%d", &dayy)
	fmt.Scanf("%s", &monthh)
	fmt.Scanf("%d", &yearr)
	fmt.Printf("day :%d,month: %s,year: %d\n", dayy, monthh, yearr)
	//-------------------------------------------------------------------
	fmt.Println("=>part2")
	var day int
	var month string
	var year int
	fmt.Println("enter your birthdate:")
	fmt.Scanf("%d\n%s\n%d", &day, &month, &year)

	fmt.Printf("%d.%s.%d\n", day, month, year)

}
