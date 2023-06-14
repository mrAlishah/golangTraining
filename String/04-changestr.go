package main

import (
	"fmt"
	"unicode/utf8"
)

func converti(s []rune) string {
	s[6] = 'd'
	s[8] = 'd'
	return string(s)

}
func main() {
	h := "Hello baby"
	fmt.Println(converti([]rune(h)))
	fmt.Println(len(h))
	fmt.Println(utf8.RuneCountInString(h))
}
