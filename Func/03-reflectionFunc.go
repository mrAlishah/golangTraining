package main

import (
	"fmt"
	"reflect"
)

//Using Reflection
//If you really want to do this dynamically on a function of fixed number of arguments, you can use reflection:
func sum(a, b, c int) int {
  return a + b + c
}

func main() {
  a := []int{1, 2, 3}

  var args []reflect.Value
  for _, v := range a {
    args = append(args, reflect.ValueOf(v))
  }
  fun := reflect.ValueOf(sum)
  result := fun.Call(args)
  sum := result[0].Interface().(int)
  fmt.Println("sum=",sum)
}