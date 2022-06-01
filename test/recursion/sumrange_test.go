package recursion

import (
	"github.com/rogermori/algorithms-datastructures-go/recursion"
	"math/rand"
	"testing"
)

func Test_sumrange_should_return_1_when_n_equals_1(t *testing.T) {
	sum := recursion.SumRange(1)
	if sum != 1 {
		t.Logf("Expected 1, got %d", sum)
		t.Fail()
	}
}

func Test_sumrange_should_return_3_when_n_equals_2(t *testing.T) {
	sum := recursion.SumRange(2)
	if sum != 3 {
		t.Logf("Expected 1, got %d", sum)
		t.Fail()
	}
}

func Test_sumrange_should_return_ntimesnplus1divided2_when_n(t *testing.T) {
	n := rand.Intn(10)
	sum := recursion.SumRange(n)
	if ((n * (n + 1)) / 2) != sum {
		t.Logf("Expected 1, got %d", sum)
		t.Fail()
	}
}
