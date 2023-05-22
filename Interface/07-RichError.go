package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type RichError interface {
	WithMessage(msg string) RichError
	WithOp(op string) RichError
	Error() string
	String() string
	ToJson() string
}

type richError struct {
	op      string `json:"op"`
	message string `json:"message"`
}

func RichErrorNew() RichError {
	return &richError{}
}

func (r *richError) WithOp(op string) RichError {
	r.op = op
	return r
}

func (r *richError) WithMessage(msg string) RichError {
	r.message = msg
	return r
}

func (r *richError) ToJson() string {
	fmt.Printf("richError: %+v | size: %d | Type: %T \n", r, unsafe.Sizeof(r), r)
	// convert the interface value to a map
	m := map[string]interface{}{
		"op":      r.op,
		"message": r.message,
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		return "struct is not json serializable"
	}
	return string(bytes)
}

func (r *richError) String() string {
	return fmt.Sprintf("op: %s , message: %s ", r.op, r.message)
}

func (r *richError) Error() string {
	return r.message
}

func main() {
	fmt.Println("\nWithByRef vs WithByValue Interface RichError------------------------------")
	r5 := RichErrorNew()
	var ir5 RichError
	ir5 = r5
	ir5.WithOp("srv5").WithMessage("test5")
	fmt.Println(ir5)
	fmt.Println(ir5.ToJson())

	var err error
	err = ir5
	fmt.Println(err.Error())

	var ir6 RichError
	ir6 = err.(RichError)
	fmt.Println(ir6.String())

}
