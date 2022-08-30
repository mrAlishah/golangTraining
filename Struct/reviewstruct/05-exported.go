package main

import (
	"fmt"
	"golangTraining/Struct/exported"
)

func main() {
	nuce := exported.Name{
		Maker: "apple",
		Price: 45,
		Model: "nike",
	}
	fmt.Println("maker:", nuce.Maker)
	fmt.Println("maker:", nuce.Price)
	fmt.Println("maker:", nuce.Model)
}
