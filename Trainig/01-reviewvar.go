package multiply

import (
	"fmt"
	"math"
	"unsafe"
)

var i int64 = 25
var j int = 45
var k = "number"

const s = 45

var sd = 79
var fj int = 82
var b string
var (
	FirstName = "fateme"
	LastName  = "Jadidi"
	Age       = 21
)
var q = float64(45)

func main() {
	var age float64
	fmt.Scanln(&age)
	fmt.Println("i:", i, "k:", k)
	fmt.Println("j:", j, "s:", s)
	fmt.Println("sd:", sd, "fj:", fj)
	b = "name be farzand"
	fmt.Println(b)
	b = "strong like go"
	fmt.Println(b)
	c := "this is declaration"
	fmt.Println(c)
	fmt.Println(FirstName, LastName, Age)
	fmt.Println("this is size of age:", unsafe.Sizeof(age))
	var N = complex(7, 5)
	fmt.Println(N)
	fmt.Println(q)
	sum := age + q
	fmt.Printf("%.2f+%.2f=%.2f\n", age, q, sum)
	H := 45
	J := 85
	sizo := math.Max(float64(H), float64(J))
	fmt.Println(sizo)
}
