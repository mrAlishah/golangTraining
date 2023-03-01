package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main() {
	fmt.Println("\nsl := []int{1, 2, 3} ------------------------------------------------------")
	sl := []int{1, 2, 3}
	slType := reflect.TypeOf(sl)
	//fmt.Println("type.String():", slType.String(), "type.Name():", slType.Name())
	examiner(slType, 0)

	fmt.Println("\ngreeting := \"hello\" ------------------------------------------------------")
	greeting := "hello"
	gType := reflect.TypeOf(greeting)
	examiner(gType, 0)

	fmt.Println("\ngreetingPtr := &greeting ------------------------------------------------------")
	greetingPtr := &greeting
	grpType := reflect.TypeOf(greetingPtr)
	examiner(grpType, 0)

	fmt.Println("\nf := Foo{A: 10, B: \"Salutations\"} ------------------------------------------------------")
	f := Foo{A: 10, B: "Salutations"}
	fType := reflect.TypeOf(f) //reflect.StructField{}
	examiner(fType, 0)

	fmt.Println("\nfp := &f ------------------------------------------------------")
	fp := &f
	fpType := reflect.TypeOf(fp)
	examiner(fpType, 0)
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "TypeString: ", t.String(), "TypeName: ", t.Name(), ", kind: ", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "=> name: ", f.Name, ", type: ", f.Type.Name(), ", kind: ", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag: ", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1: ", f.Tag.Get("tag1"), ", tag2: ", f.Tag.Get("tag2"))
			}
		}
	}
}
