package recursion

import (
	"github.com/rogermori/algorithms-datastructures-go/recursion"
	"testing"
)

func Test_should_return_array_including_13_when_array_contains_123(t *testing.T) {
	intList := []int{1, 2, 3}
	oddList := recursion.CollectOddValuesWithHelperFunction(intList)
	t.Logf("%v\n", oddList)

	oddList1 := recursion.CollectOddValuesPure(intList)
	t.Logf("%v\n", oddList1)

	if oddList[0] != 1 || oddList[1] != 3 {
		t.Logf("[helper] expected [1,3] , resulted %v\n", oddList)
		t.Fail()
	}
	if oddList1[0] != 1 || oddList1[1] != 3 {
		t.Logf("[pure] expected [1,3] , resulted %v\n", oddList)
		t.Fail()
	}

}
