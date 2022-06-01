package recursion

import (
	"github.com/rogermori/algorithms-datastructures-go/recursion"
	"testing"
)

func Test_should_return_array_including_1_when_array_only_contains_1(t *testing.T) {
	intList := []int{1}
	oddList := recursion.CollectOddValues(intList)
	if len(oddList) != 1 {
		t.Logf("expected [1] but resulted %v\n", oddList)
		t.Fail()
	}
	if oddList[0] != 1 {
		t.Logf("expected 1 but resulted %v\n", oddList[0])
		t.Fail()
	}

}
