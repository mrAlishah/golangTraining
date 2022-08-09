package main

import "fmt"

func converti(s []rune) string {
	s[6] = 'd'
	s[8] = 'd'
	return string(s)

}
func main() {
	h := "Hello baby"
	fmt.Println(converti([]rune(h)))
}
