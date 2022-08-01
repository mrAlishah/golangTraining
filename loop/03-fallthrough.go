//package main
//
//import (
//	"fmt"
//)
//fmt.Println("=>Part2")
//func number() int {
//	num := 15 * 5
//	return num
//}
//
//func main() {
//
//	switch num := number(); { //num is not a constant
//	case num < 30:
//		fmt.Printf("%d is lesser than 30\n", num)
//		fallthrough
//	case num < 50:
//		fmt.Printf("%d is lesser than 50\n", num)
//		fallthrough
//	case num < 60:
//		fmt.Printf("%d is lesser than 60\n", num)
//		fallthrough
//	case num < 100:
//		fmt.Printf("%d is lesser than 100\n", num)
//	}
//
//}

// use a fallthrough for executed other lines
//Fallthrough happens even when the case evaluates to false

//===================================================================

package main

import "fmt"

func main() {
	fmt.Println("=>Part2")

	var numbi int
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &numbi)

	switch {
	case numbi > 100:
		fmt.Printf("%d bigger than 100\n", numbi)
		fallthrough
	case numbi < 100:
		fmt.Printf("%d is lesser than 100\n", numbi)

	}

}
