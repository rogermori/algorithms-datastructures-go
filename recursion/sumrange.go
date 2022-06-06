package recursion

func SumRange(n int) int {
	// base case n = 1
	if n == 1 {
		return 1
	}
	return n + SumRange(n-1)
}
