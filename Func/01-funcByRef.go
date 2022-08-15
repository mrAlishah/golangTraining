package main

import "fmt"

//! int Basic Data Type -----------------------------------------------------------------
func updateIntByValue(param int) {
	param = 10 
}

func updateIntByRef(param *int) {
	*param = 10 
}

//! string Basic Type -----------------------------------------------------------------
func updateStringByValue(str string) {
	str = "changed"
}

func updateStringByRef(str *string) {
	*str = "changed"
	//*str[2]='1' //ERROR : It's not possible like another languages
}

//! Array Aggregate Type -----------------------------------------------------------------
//declare explicitly array length as param and return
func updateArrayByValue(array [3]int) {
	array[0]=10
}

func updateArrayByRef(array *[3]int) {
	array[0]=10
}

func updateArrayByReturn(array [3]int) [3]int{
	array[1]=20
	return array
}

//! Struct Aggregate Type -----------------------------------------------------------------
type Employee struct {
	Name string
	Age int
	Designation string
  }

func updateStructByValue(emp Employee) {
	emp.Name = "Robe";
}

func updateStructByRef(emp *Employee) {
	emp.Name = "Robe";
}

func main(){

//! simple send Basic Data Type argument -----------------------------------------------------------------
	fmt.Println("==========================================================================")
	arg1 := 2
	fmt.Println("arg1 = ", arg1)
	updateIntByValue(arg1)
	fmt.Println("updateIntByValue = ", arg1)
	updateIntByRef(&arg1)
	fmt.Println("updateIntByRef = ", arg1)

//! simple send Basic Type string -----------------------------------------------------------------
	fmt.Println("==========================================================================")
	arg2 := "string"
	//arg2[2] = "1" //ERROR : It's not possible like another languages
	fmt.Println("arg2 = ", arg2)
	updateStringByValue(arg2)
	fmt.Println("updateStringByValue = ", arg2)
	updateStringByRef(&arg2)
	fmt.Println("updateStringByRef = ", arg2)

//! simple send Aggregate Type Array -----------------------------------------------------------------
	//Arrays in Golang are value types unlike other languages like C, C++, and Java where arrays are reference types.
	fmt.Println("==========================================================================")
	arg3 := [3] int{1, 2, 3}
	fmt.Println("arg3 = ", arg3)
	updateArrayByValue(arg3)
	fmt.Println("updateArrayByValue = ", arg3)
	updateArrayByRef(&arg3)
	fmt.Println("updateArrayByRef = ", arg3)
	arg3=updateArrayByReturn(arg3)
	fmt.Println("updateArrayByReturn = ", arg3)

//! simple send Aggregate Type Struct -----------------------------------------------------------------
	//Struct in Golang are value types not reference types.
	fmt.Println("==========================================================================")
	arg4 := Employee{"john", 20, "Developer"}
	fmt.Println("arg4 = ", arg4)
	updateStructByValue(arg4)
	fmt.Println("updateStructByValue = ", arg4)
	updateStructByRef(&arg4)
	fmt.Println("updateStructByRef = ", arg4)


}