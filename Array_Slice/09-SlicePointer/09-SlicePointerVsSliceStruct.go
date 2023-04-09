package main

import (
	"fmt"
)

// go build -gcflags="-m -l" ./09-SlicePointerVsSliceStruct.go
func main() {
	SurveyOneStruct()
	//SurveySliceStruct()
}

func SurveyOneStruct() {
	fmt.Println("----------- ReturnOneStructWithStructs")
	pS := ReturnOneStructWithStructs(1)
	fmt.Printf("pS = %+v \n", pS)

	fmt.Println("----------- ReturnOneStructWithPointers")
	pP := ReturnOneStructWithPointers(1)
	fmt.Printf("pP = %+v \n", pP)

	fmt.Println("----------- ReturnOneStructAsParamPointers")
	var p Person
	ReturnOneStructAsParamPointers(&p)
	fmt.Printf("p = %+v \n", p)
}

func SurveySliceStruct() {
	n := 3

	fmt.Println("----------- ReturnSliceWithPointers")
	perS := ReturnSliceWithStructs(n)
	for i := 0; i < n; i++ {
		fmt.Printf("perS%d = %+v \n", i, perS[i])
	}

	fmt.Println("----------- ReturnSliceWithPointers")
	perP := ReturnSliceWithPointers(n)
	for i := 0; i < n; i++ {
		fmt.Printf("perP%d = %+v \n", i, *perP[i])
	}

	fmt.Println("----------- ReturnSliceWithParamPointerSliceStruct")
	var perPPS []Person
	ReturnSliceWithParamPointerSliceStruct(&perPPS, n)
	for i := 0; i < n; i++ {
		fmt.Printf("perPPS%d = %+v \n", i, perPPS[i])
	}

	fmt.Println("----------- ReturnSliceWithParamPointersSlicePointer")
	var perPPP []*Person
	ReturnSliceWithParamPointersSlicePointer(&perPPP, n)
	for i := 0; i < n; i++ {
		fmt.Printf("perPPP%d = %+v \n", i, *perPPP[i])
	}
}

type Person struct {
	Age int
}

/*************************************************************************************************/
/* Return Slice of Struct
/*************************************************************************************************/

func ReturnSliceWithStructs(size int) []Person {
	res := make([]Person, size)

	for i := 0; i < size; i++ {
		res[i] = Person{Age: i}
	}

	return res
}

func ReturnSliceWithPointers(size int) []*Person {
	res := make([]*Person, size)

	for i := 0; i < size; i++ {
		res[i] = &Person{Age: i}
	}

	return res
}

func ReturnSliceWithParamPointerSliceStruct(res *[]Person, size int) {
	*res = make([]Person, size)

	for i := 0; i < size; i++ {
		(*res)[i] = Person{Age: i}
	}

}

func ReturnSliceWithParamPointersSlicePointer(res *[]*Person, size int) {
	*res = make([]*Person, size)

	for i := 0; i < size; i++ {
		(*res)[i] = &Person{Age: i}
	}

}

// this code it's not important ,It's just for testing sth, we just set Make outside the func

func ReturnSliceWithParamPointerSliceStructMaked(res *[]Person, size int) {

	for i := 0; i < size; i++ {
		(*res)[i] = Person{Age: i}
	}

}

// this code it's not important ,It's just for testing sth, we just set Make outside the func

func ReturnSliceWithParamPointersSlicePointerMaked(res *[]*Person, size int) {

	for i := 0; i < size; i++ {
		(*res)[i] = &Person{Age: i}
	}

}

/*************************************************************************************************/
/* Return just one Struct
/*************************************************************************************************/

// to return a copy of the struct
// RESULT:: It could be ok for small structs (because the overhead is minimal)

func ReturnOneStructWithStructs(age int) Person {
	res := Person{Age: age}

	return res
}

// a pointer to the struct value created within the function
// RESULT:: It could be ok for bigger struct used

func ReturnOneStructWithPointers(age int) *Person {
	res := &Person{Age: age}

	return res
}

// it expects an existing struct to be passed in and overrides the value.
// RESULT:: if you want to be extremely memory efficient because you can easily reuse a single struct instance between calls.

func ReturnOneStructAsParamPointers(res *Person) {
	res.Age = 1
}
