package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//char in golang is byte(ASCII) & rune(UTF8)
func main() {
	var char1 byte = 'a'
	var char2 rune = 'a'

	//var char3 byte = '♥' //ERROR: it'snot able to assign
	var char4 rune = '♥'

	fmt.Printf("byte as ASCII Char: | %T | %d | %v | %c | %d \n",char1,unsafe.Sizeof(char1),char1,char1,char1)
	fmt.Printf("rune as ASCII Char: | %T | %d | %v | %c | %U \n",char2,unsafe.Sizeof(char2),char2,char2,char2)
	//fmt.Printf("byte as UTF Char: | %T | %d | %v | %c | %U \n",char3,unsafe.Sizeof(char3),char3,char3,char3)
	fmt.Printf("rune as UTF Char:   | %T | %d | %v | %c | %U \n",char4,unsafe.Sizeof(char4),char4,char4,char4)
	
	fmt.Printf("Size: %d | Type: %s | Character: %c\n", unsafe.Sizeof(char1), reflect.TypeOf(char1), char1)
	fmt.Printf("Size: %d | Type: %s | Character: %c\n", unsafe.Sizeof(char2), reflect.TypeOf(char2), char2)
 
}
