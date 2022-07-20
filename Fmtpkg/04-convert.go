package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("=>convert string to integer ")
	str := "22"
	inte, err := strconv.Atoi(str)
	fmt.Println(str, inte, err, reflect.TypeOf(inte), reflect.TypeOf(str))

	//------------------------------------------------------------------------
	fmt.Println("=>convert integer to string")
	integ := 14
	stri := strconv.Itoa(integ)
	fmt.Println(stri, reflect.TypeOf(integ), reflect.TypeOf(stri))
	//------------------------------------------------------------------------
	fmt.Println("=>convert float ot integer to string")
	flo := 22.55
	a := fmt.Sprintf("%f", flo)
	fmt.Printf("%s\t", a)
	fmt.Println(reflect.TypeOf(a))
	//-------------------------------------------------------------------------

}
