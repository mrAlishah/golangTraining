package number

import (
	"testing"
)

var table = []struct {
	in string
	out *NumberData
	err string
} {
	{
		in: "not a number",
		out: &NumberData{
			value: 0,
			isNumeric: false,
			isInteger: false,
			isNegative: false,
		},
		err: "Did not evaluate non-number correctly",
	},
	{ // Heap Map to coverage result.isNumeric = false
		in: "not a number",
		out: &NumberData{
			value: 0,
			isNumeric: false,
			isInteger: false,
			isNegative: false,
		},
		err: "Did not evaluate non-number correctly",
	},
	{
		in: "42.2",
		out: &NumberData{
			value: 42.2,
			isNumeric: true,
			isInteger: false,
			isNegative: false,
		},
		err: "Did not evaluate positive, non-integer number correctly",
	},	
}

//go test ./number -coverprofile=number.out -covermode=count
//go tool cover -html=number.out
func TestEvaluatesNotNumberCorrectly(t *testing.T){
	for _,entry := range table {
		//arrange
		condidate := entry.in 

		//act
		result := NumberEvaluator(condidate)

		//assert
		if  result.isNumeric != entry.out.isNumeric ||
			result.value != entry.out.value ||
			result.isNegative != entry.out.isNegative ||
			result.isInteger != entry.out.isInteger {
				t.Log(result)
				t.Log(entry.out)
				t.Error(entry.err)
			}

	}
}