package main

import "fmt"

type Passenger struct {
	name         string
	ticketnumber int
	boarded      bool
}

type Bus struct {
	frontseat Passenger
}

func main() {

	casey := Passenger{"casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{name: "Bill", ticketnumber: 2}
		ella = Passenger{name: "Ella", ticketnumber: 3}
	)
	fmt.Println(bill, ella)

	var hidi Passenger
	hidi.name = "Hidi"
	hidi.ticketnumber = 4
	fmt.Println(hidi)

	casey.boarded = true
	bill.boarded = true
	if bill.boarded {
		fmt.Println("Bill has boarded to bus")

	}
	if casey.boarded {
		fmt.Println(casey.name, " has boarded to bus")

	}
	hidi.boarded = true
	bus := Bus{hidi}
	fmt.Println(bus)
	fmt.Println(bus.frontseat.name, "frontseat")
}
