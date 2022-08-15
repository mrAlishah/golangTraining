package main

import "fmt"

// Variadic Functions
//Variadic functions can be called with any number of trailing arguments.
func sum(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func print(args ...int) {
    for i, n := range args {
        fmt.Printf("arg%d = %d \n",i,n)
    }
}

func main() {
  a := []int{1, 2, 3}

  print(a...)

  t := sum(a...)
  fmt.Println("sum=",t)

  t = sum(2, 3, 4)
  fmt.Println("sum=",t)
}