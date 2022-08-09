package main

import "fmt"

func Printbytes(s string) {
	fmt.Printf("bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
}
func Printchar(s string) {
	fmt.Printf("characters:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
}

func main() {
	name := "Hello world"
	fmt.Printf("string %s\n", name)
	Printbytes(name)
	fmt.Printf("\n")

	Printchar(name)

}
