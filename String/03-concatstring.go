package main

import "fmt"

func main() {
	//first way to concat two string
	var (
		f       = "Fateme"
		j       = "Jadidi"
		result1 = f + "" + j
	)
	fmt.Printf(result1)
	fmt.Println("\n")
	//=====================================
	//second way to concat two string
	var (
		m       = "Mostafa"
		a       = "Alishah"
		result2 = fmt.Sprintf("%s %s\n", m, a)
	)
	fmt.Printf(result2)
	//=====================================

}
