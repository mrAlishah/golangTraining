package main

import "fmt"

func main() {
	/*type Member struct {
	fname  string //fields
	lname  string //fields
	salary int32  //fields*/

	//}
	member := struct {
		fname  string
		lname  string
		salary int
	}{
		fname:  "fateme",
		lname:  "jadidi",
		salary: 0,
	}

	fmt.Println(member)
	fmt.Println("fname :", member.fname) //Accessing individual fields of a struct

}
