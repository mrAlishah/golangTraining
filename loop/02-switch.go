package main

import "fmt"

func main() {
	fmt.Println("part1:simple example")
	human_body := 2
	fmt.Printf("this part of body help to us for thinking:")

	switch human_body {
	case 1:
		fmt.Println("heart")
	case 2:
		fmt.Println("brain")
	case 3:
		fmt.Println("hand")
	case 4:
		fmt.Println("leg")

	}
	//-----------------------------------------------------
	fmt.Println("part 2:default case")
	letter := "a"
	fmt.Printf("%s is ", letter)
	switch letter {
	case "a", "u", "o", "i", "e":
		fmt.Println("vowel sound")
	default:
		fmt.Println("not a vowel")
	}
	//-----------------------------------------------------
	fmt.Println("=part3: Expressionless switch\n")
	num := 120
	switch {
	case num < 100:
		fmt.Printf("%d is smaller than 100\n", num)
	case num > 100:
		fmt.Printf("%d is bigger than 100\n", num)

	}
	//----------------------------------------------------
	fmt.Println("part4:use switch with scanf")
	var i string
	fmt.Scanf("%s", &i)
	switch i {
	case "Mahdi":
		fmt.Println(" Mehrabi", i)
	case "Mostafa":
		fmt.Println("Alishah", i)
	case "Fateme":
		fmt.Println("Jadidi", i)
	default:
		fmt.Println("we cant guess your family", i)

	}
	//-----------------------------------------------------

}
