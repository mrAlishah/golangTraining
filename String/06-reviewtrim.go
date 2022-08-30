package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "##this is so good##"
	b := "**not bad**"

	fmt.Println(a, b)
	res1 := strings.Trim(a, "#")
	res2 := strings.Trim(b, "*")
	fmt.Println(res1, res2)
	c := "++today all thing is good**"
	d := "--but i need money^^"
	fmt.Println(c, d)
	res3 := strings.TrimRight(c, "*")
	res4 := strings.TrimRight(d, "^")
	fmt.Println(res3, res4)
	e := "@@money for what?$$"
	f := "((book&&"
	fmt.Println(e, f)
	res5 := strings.TrimLeft(e, "@")
	res6 := strings.TrimLeft(f, "(")
	fmt.Println(res5, res6)
	g := "		found and		"
	h := "		lost that		"
	fmt.Println(g, h)
	res7 := strings.TrimSpace(g)
	res8 := strings.TrimSpace(h)
	fmt.Println(res7, res8)
	i := "welcome,to Microsoft"
	j := "you should learn go or python"
	fmt.Println(i, j)
	res9 := strings.TrimSuffix(i, "in Microsoft")
	res10 := strings.TrimSuffix(j, "or python")
	fmt.Println(res9, res10)
	k := "bad day,be happy"
	l := "earn money, get happiness"
	fmt.Println(k, l)
	res11 := strings.TrimPrefix(k, "bad day,")
	res12 := strings.TrimPrefix(l, "earn money,")
	fmt.Println(res11, res12)

}
