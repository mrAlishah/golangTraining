package main

import "fmt"

type Employee struct {
	Name string
	Age int
	Designation string
	Salary int
  }

func UpdateEmployeeByValue(emp Employee) {
	emp.Name = "Parham";
}

func UpdateEmployeeByRef(emp *Employee) {
	emp.Name = "Parham";
}

func (emp Employee) PrintName() {
	fmt.Println("User Name: ", emp.Name)
}

func (emp Employee) UpdateMethodByValue(newName string) {
	emp.Name = newName
}

func (emp *Employee) UpdateMethodByRef(newName string) {
	emp.Name = newName
}

func main(){

//** Creating Objects From Struct
	var newEmployee = Employee{"Saeed", 20, "Developer", 10}
	/*
	Problems with the above approach:
	1-We need to remember all the properties and their order while creating the object.
	2-All of the values need to be passed to the struct.
	*/
	
	var otherEmployee = Employee{Designation: "Developer", Name: "Mehdi", Age: 30}
	/*
	Advantage:
	1-No need to add values in the same order.
	2-No need to specify all the key-value pairs.

	we’re omitting one of the key-values (Salary). By default, Golang adds a default value to the property Salary.
	The int value is assigned with the value 0.
	The string values are assigned with “” (empty string).
	The bool values are assigned with the default value, false.
	*/

//================================================================================================
//** Object reference or object value?
	UpdateEmployeeByValue(newEmployee)
	fmt.Println("Object by Value")
	fmt.Println(newEmployee.Name)
	//when we’re updating the values in the UpdateEmployee function, 
	//the original object newEmployee isn’t impacted since we’ve passed the object as a value and not a reference.

	UpdateEmployeeByRef(&otherEmployee)
	// The original Object sent to the function Updated...
	fmt.Println("Object by Reference")
	fmt.Println(otherEmployee.Name)	
	/*
	the function (as well as the function call) has been updated to send and receive the memory location instead of the value 
	for the ibject. Now if the data is updated in the called function, then that’ll be replicated in the original data object.
	*/

	//The problem in this approach is that we have to explicitly use the address-of operator to extract 
	//the address of the object and send it to the function.
	var moreEmployee = new(Employee)
	moreEmployee.Name = "Reza"
	moreEmployee.Age = 30
	UpdateEmployeeByRef(moreEmployee)
	fmt.Println("Object by Ref and created by new()")
	fmt.Println(moreEmployee.Name)	
	/*
    In the above case, the object reference is returned from the new keyword. Therefore, while invoking the function, 
	we don’t need to use & to send the reference of the object.	
	*/
//================================================================================================
//** Adding a Method to the Struct

	/*
	we’re adding a new function that’s bound to the Employee struct. We need to explicitly bind the function to 
	the struct. This function defined can then take the reference of the object created using emp.
	*/
	fmt.Println("========================")
	fmt.Println("Add method to struct")
	newEmployee.PrintName()

	fmt.Println("reciver by value")
	newEmployee.UpdateMethodByValue("Sarah")
	newEmployee.PrintName()

	fmt.Println("reciver by Ref")
	newEmployee.UpdateMethodByRef("Sarah")
	newEmployee.PrintName()	
}