package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}
type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}
func (c Contract) CalculateSalary() int {
	return c.basicpay
}
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("total expense per month $%d", expense)
}
func main() {
	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
	}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

}
