package main

import "fmt"

type Findvowel interface {
	Findvowels() []rune
}
type myString string

func (bs myString) Findvowels() []rune {
	var vowels []rune
	for _, rune := range bs {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels

}
func main() {
	var name myString
	fmt.Scanln(&name)
	var v Findvowel
	v = name
	fmt.Printf("vowels are %c\n", v.Findvowels())
}
