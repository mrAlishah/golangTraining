package main

import "fmt"

func Comparestring(str1, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s , %s is equal", str1, str2)
		return
	}
	fmt.Printf("%s , %s is not equal", str1, str2)

}
func main() {
	string1 := "fateme"
	string2 := "jadidi"
	Comparestring(string1, string2)
	fmt.Printf("\n")
	d := "math"
	j := "computer"
	Comparestring(d, j)
	fmt.Printf("\n")

}
