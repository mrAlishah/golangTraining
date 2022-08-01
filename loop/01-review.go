package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("%d", i)
	}
	for j := 0; j <= 10; j++ {
		fmt.Println("%d", j)
		if j >= 5 {
			break
		}
	}
	for k := 0; k <= 10; k++ {

		if k == 5 {
			continue
		}
		fmt.Println("%d", k)
	}
	//--------------------------------------
	fmt.Println("=>simple for")
	for n := 0; n < 100; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("%d", n)
	}
	//--------------------------------------
	fmt.Println("=>number of divided to 6")
	for ii := 0; ii <= 100; ii++ {
		if ii%6 == 0 {
			fmt.Printf("%d\t", ii)
		}
	}
	//--------------------------------------

	fmt.Println("\n=>normal nested")
	for b := 0; b < 3; b++ {
		for c := 1; c < 4; c++ {
			fmt.Printf("b = %d,  c = %d\n", b, c)

		}
	}
	//--------------------------------------
	fmt.Println("=>nested with break without condition")
	for d := 0; d <= 4; d++ {
		for e := 1; e < 5; e++ {
			fmt.Printf("d = %d, e = %d\n", d, e)
			break
		}
	}
	//----------------------------------------
	fmt.Println("=>nested with break")
	for s := 0; s <= 8; s++ {
		for m := 2; m < 30; m++ {
			fmt.Printf("s=%d , m=%d\n", s, m)
			if s == m {
				break
			}
		}
	}
	//------------------------------------
	fmt.Println("=>nested with break & label")
outer:
	for g := 0; g < 10; g++ {
		for h := 1; h < 12; h++ {
			fmt.Printf("g=%d , h=%d\n", g, h)

			if g == h {
				break outer
			}

		}
	}
	//-----------------------------------
	fmt.Println("=>print even number")

	for A := 0; A <= 10; A++ {

		if A%2 == 0 {
			fmt.Printf("A=%d\n", A)
		}

	}
	//	//------------------------------------
	fmt.Println("=>create a simple calculator")
	for D := 0; D <= 5; D++ {
		for V := 0; V <= 5; V++ {
			fmt.Printf("%d*%d=%d\n", D, V, D*V)
		}
	}
	//-------------------------------------
	fmt.Println("simple example when we have same variable in and out side of loop")
	for n := 1; n <= 5; n++ {
		for n := 1; n <= 3; n++ {
			fmt.Println("n1:", n)
		}
		fmt.Println("n2:", n)
		fmt.Println("====================")

	}
	//above example run correctly & print -n-
	//---------------------------------------
	/* fmt.Println("infinite loop ")
	for n := 1; n <= 5; n++ {
		for n = 1; n <= 3; n++ {
			fmt.Println("n1:", n)
		}
		fmt.Println("n2:", n)
		fmt.Println("====================")

	}*/ //this loop is infinite so i command that
	//Above example shows that when we omit : compiler thinks internal variable is a new variable
	//----------------------------------------

	//----------------------------------------

}
