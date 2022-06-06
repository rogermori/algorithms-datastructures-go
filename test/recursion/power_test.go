package recursion

import (
	"github.com/rogermori/algorithms-datastructures-go/recursion"
	"testing"
)

func TestPower_to_0_should_return_1(t *testing.T) {
	power := recursion.Power(10, 0)
	if power != 1 {
		t.Logf("Expect %d but resulted %d", 1, power)
		t.Fail()
	}
}

func TestPower_to_1_should_return_base(t *testing.T) {
	power := recursion.Power(10, 1)
	if power != 10 {
		t.Logf("Expect %d but resulted %d", 1, power)
		t.Fail()
	}
}

func TestPower_2_power_5_should_return_32(t *testing.T) {
	power := recursion.Power(2, 5)
	if power != 32 {
		t.Logf("Expect %d but resulted %d", 1, power)
		t.Fail()
	}
}
