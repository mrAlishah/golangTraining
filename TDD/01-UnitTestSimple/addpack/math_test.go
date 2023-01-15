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
}