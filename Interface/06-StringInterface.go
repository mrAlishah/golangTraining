package main

import (
	"fmt"
)

type richErrorByValue struct {
	op      string
	message string
}

func RichErrorNewByValue() richErrorByValue {
	return richErrorByValue{}
}

func (r richErrorByValue) WithOpByValue(op string) richErrorByValue {
	r.op = op
	return r
}

func (r richErrorByValue) WithMessageByValue(msg string) richErrorByValue {
	r.message = msg
	return r
}

func (r richErrorByValue) String() string {
	return fmt.Sprintf("op: %s , message: %s ", r.op, r.message)
}

type richErrorByRef struct {
	op      string
	message string
}

func RichErrorNewByRef() *richErrorByRef {
	return &richErrorByRef{}
}

func (r *richErrorByRef) WithOpByRef(op string) *richErrorByRef {
	r.op = op
	return r
}

func (r *richErrorByRef) WithMessageByRef(msg string) *richErrorByRef {
	r.message = msg
	return r
}

func (r *richErrorByRef) String() string {
	return fmt.Sprintf("op: %s , message: %s ", r.op, r.message)
}

func (r *richErrorByRef) Error() string {
	return r.message
}

func main() {
	fmt.Println("\nString() struct ------------------------------")
	r1 := richErrorByValue{
		message: "test",
	}
	fmt.Println(r1)

	fmt.Println("\nWithByRef vs WithByValue struct RichErrorNewByValue------------------------------")
	r2 := RichErrorNewByValue()
	r2.WithMessageByValue("test2")
	fmt.Println(r2, " > r2.WithMessageByValue(\"test2\")")
	r2.WithMessageByValue("test2").WithOpByValue("srv2")
	fmt.Println(r2, " > r2.WithMessageByValue(\"test2\").WithOpByValue(\"srv2\")")

	fmt.Println("\nWithByRef vs WithByValue struct RichErrorNewByRef ------------------------------")
	r3 := RichErrorNewByRef()
	r3.WithMessageByRef("test3").WithOpByRef("srv3")
	fmt.Println(r3, " > r2.WithMessageByRef(\"test2\").WithOpByRef(\"srv3\")")

}
