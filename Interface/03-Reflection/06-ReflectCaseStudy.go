package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

/*
Making Functions
Reflection doesn’t just let you make new places to store data. You can use reflection to make new functions using t
he reflect.MakeFunc function. This function expects the reflect.Type for the function that we want to make and
a closure whose input parameters are of type []reflect.Value and whose output parameters are also of type []reflect.Value.
Here’s a quick example, which creates a timing wrapper for any function that’s passed into it:
*/

func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now() //for end.Sub(start)
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func timeMe() {
	fmt.Println("timeMe()------------------------------------------------------")
	fmt.Println("starting")
	time.Sleep(1 * time.Second)
	fmt.Println("ending")
}

func timeMeToo(a int) int {
	fmt.Println("timeMeToo(a int)------------------------------------------------------")
	fmt.Println("starting")
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	fmt.Println("ending")
	return result
}

func main() {
	timed := MakeTimedFunction(timeMe).(func())
	timed()
	timedToo := MakeTimedFunction(timeMeToo).(func(int) int)
	fmt.Println(timedToo(2))
}
