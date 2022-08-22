package main

import "fmt"

type Emp struct {
	name string
	age  int
}

func (e Emp) changeName(newName string) {
	e.name = newName

}
func (e *Emp) changeAge(newAge int) {
	e.age = newAge
}
func main() {
	e := Emp{
		name: "fatemejadidi",
		age:  50,
	}
	fmt.Printf("Emp name before change :%s", e.name)
	e.changeName("Fateme_jadidi")
	fmt.Printf("\nEmp name after change:%s", e.name)
	fmt.Printf("\n\nEmp age before :%d", e.age)
	(&e).changeAge(59) //also we can use e.changeAge() =>without &
	fmt.Printf("\nEmp age after change:%d\n", e.age)
}
