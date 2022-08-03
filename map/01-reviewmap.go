package main

import "fmt"

func main() {
	employe := map[string]int{
		"fateme": 15000,
		"zahra":  42000,
		"mehri":  31000,
		"gafor":  42000,
		"mahnaz": 31000,
		"mmd":    42000,
		"yosef":  31000,
	}
	employeneed := "zahra"
	salary := employe[employeneed]
	fmt.Println("this is ", employeneed, " salary:", salary)
	fmt.Println("==================================================")
	//Deleting items from a map
	delete(employe, "fateme")
	fmt.Println(employe) //so delete fateme in map
	fmt.Println("==================================================")
	fmt.Println("example number 87")
	modi := employe
	modi["mmd"] = 25000
	fmt.Println("change value for mmd in employe: ", employe)
}
