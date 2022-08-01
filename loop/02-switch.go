package main

import (
	"fmt"
	"math/rand"
)

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
	fmt.Println("part4:Enter your first name") //use switch with scanf
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
	fmt.Println("part5:breaking switch")
	switch manoal := -5; {
	case manoal < 50:
		if manoal < 0 {
			break
		}
		fmt.Printf("%d is lesser than 50\n", manoal)
	case manoal < 100:
		fmt.Printf("%d is lesser than 100\n", manoal)
		fallthrough
	case manoal < 200:
		fmt.Printf("%d is lesser than 200\n", manoal)
	}
	//in above program The break statement terminates the switch before it completes and the program doesn't print anything
	//======================================================
	fmt.Println("=>part6:Breaking the outer for loop")
randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Printf("Generated even number %d\n", i)
			break randloop
		}
	}

	//if the break statement is used without the label, the switch statement will only be broken and the loop will continue running.
	//So labeling the loop and using it in the break statement inside the switch is necessary to break the outer for loop.
	//this loop is infinite loop & rand.Intn generate random number

	//=======================================================

	fmt.Println("=>part 7:challenge in this tutorial")
reandloop:
	for {
		switch h := rand.Intn(100); {
		case h%3 == 0:
			fmt.Printf("Generated even number %d\n", h)
			break reandloop

			fallthrough

		case h%2 == 0:
			fmt.Printf("Generated even number %d\n", h)
			break reandloop

		}
	}
	//==========================================================
	fmt.Println("=>part8:print number in range 0 to 100 which dividable to 2 or 3 or 6")
	for i := 0; i <= 100; i++ {
		switch {
		case i%2 == 0:
			fmt.Printf("%d dividable to 2\n", i)
		case i%3 == 0:
			fmt.Printf("%d dividable to 3\n ", i)
		case i%6 == 0:
			fmt.Printf("%d dividable to 6\n ", i)

		}
	}
	//==========================================================
	fmt.Println("=>part9:print number in range 0 to 100 which dividable to 2 and 3 and 6")

	for i := 0; i <= 100; i++ {
		switch {
		case i%2 == 0 && i%3 == 0 && i%6 == 0:
			fmt.Printf("%d dividable to 2&3&6\n", i)

		}
	}

	//============================================================

	fmt.Println("========================================")
	fmt.Println("Enter your age:")
	var age int
	fmt.Scanf("%d", &age)
	switch {
	case age == 0:
		fmt.Println("newborn")
	case age >= 1 && age <= 3:
		fmt.Println("toddler")
	case age < 13:
		fmt.Println("child")
	case age < 18:
		fmt.Println("teenager")
	default:

		fmt.Println("adult")

	}
}
