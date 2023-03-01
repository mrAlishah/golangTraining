package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func main() {
	i := 10
	fmt.Printf("1- Type: %T , Value: %d \n", i, i)
	fmt.Println("------------------------------------------------------")
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Printf("2- Type: %T , Value: %d \n", o, o)

	//The concrete type of interface{} is represented by reflect.TypeOf()
	// The underlying value is represented by reflect.ValueOf().
	fmt.Println("------------------------------------------------------")
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	fmt.Println("3- Type: ", t)
	fmt.Println("3- Value: ", v)

	//Type represents the actual type of the interface{}, in this case main.Order
	//Kind represents the specific kind of the type. In this case, it's a struct.
	fmt.Println("------------------------------------------------------")
	t = reflect.TypeOf(o)
	k := t.Kind()
	fmt.Println("4- Type: ", t)
	fmt.Println("4- Kind: ", k)

	//The NumField() method returns the number of fields in a struct. reflect.ValueOf(o).NumField()
	//the Field(i int) method returns the reflect. Value of the ith field. reflect.ValueOf(o).Field(i)
	fmt.Println("------------------------------------------------------")
	if reflect.ValueOf(o).Kind() == reflect.Struct {
		v := reflect.ValueOf(o)
		fmt.Printf("5- Number of ValueOf fields:%d \n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("5- Field:%d , type:%T , value:%v\n", i, v.Field(i), v.Field(i))
		}

		t = reflect.TypeOf(o)
		fmt.Printf("5- Number of TypeOf fields: %d , Name: %v\n", t.NumField(), t.Name())
		for i := 0; i < t.NumField(); i++ {
			fmt.Printf("5- Field:%d , type:%T , value:%v , FieldName: %v\n", i, t.Field(i), t.Field(i), t.Field(i).Name)
		}
	}

	//The methods Int and String help extract the reflect.Value as an int64 and string respectively.
	fmt.Println("------------------------------------------------------")
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("6- type: %T value: %v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("6- type: %T value: %v\n", y, y)

	fmt.Println("------------------------------------------------------")
	if reflect.ValueOf(o).Kind() == reflect.Struct {
		t := reflect.TypeOf(o).Name()
		fmt.Printf("7- reflect.TypeOf(o).Name(): %v\n", t)

		v := reflect.ValueOf(o)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				fmt.Printf("7- v.Field(i): %d\n", v.Field(i).Int())
			case reflect.String:
				fmt.Printf("7- v.Field(i): %s\n", v.Field(i).String())
			default:
				fmt.Println("7- Unsupported type")
			}
		}
	}
}
