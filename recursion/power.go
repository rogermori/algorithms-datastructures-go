package recursion

func Power(base, exponent int) int {
	if exponent == 1 {
		return base
	} else if exponent == 0 {
		return 1
	}
	return base * Power(base, exponent-1)
}
