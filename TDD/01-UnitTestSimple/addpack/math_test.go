package addpack

import (
	"testing"
)

func TestCanAddNumbers(t *testing.T) {
	result := Add(1,2)

	if (result != 3){
		t.Log("Err --> Failed to add 1+2")
		t.Fail()
	}

    //----------------------------------------------------------------
	
    result = Add(1,2,3,4)

	if result != 10 {
		t.Error("Err --> Failed to add more than two numbers")
	}
}

var addTable = []struct {
	in []int
	out int
} {
	{[]int{1,2}, 3},
	{[]int{3,2}, 5},
	{[]int{1,2,3,4}, 10},
}
func TestCanAddNumByTables(t *testing.T) {
	t.Parallel()

	for _, entry := range addTable {
		result := Add(entry.in...)
		if result != entry.out {
			t.Error("Err --> Failed to add numbers as expected")
		}
	}
}