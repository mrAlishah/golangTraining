package main

import "fmt"

func main() {
	var firststruct struct {
		field string
		a, b  int
	}
	firststruct.field = "fateme"
	firststruct.a = 15
	firststruct.b = 51
	fmt.Println(firststruct)
	//===============================================
	secondstruct := struct {
		field string
		c, d  int
		h, j  rune
	}{
		"fateme",
		45, 84,
		'f', 'j',
	}
	fmt.Println(secondstruct)
	//===============================================

}
