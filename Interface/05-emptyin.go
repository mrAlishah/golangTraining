package main

import "fmt"

type mn struct {
	i int
}

func Describe(i interface{}) {
	fmt.Printf("Type:%T,value=%v\n", i, i)

}
func main() {
	s := "always bad"
	Describe(s)
	i := 55
	Describe(i)
	d := 'a'
	Describe(d)
	k := true
	Describe(k)
	Describe(mn{i: 45})
	strt := struct {
		name string
	}{
		name: "fateme jadidi",
	}
	Describe(strt)

}
