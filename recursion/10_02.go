package recursion

func numWays(n int) int {
	x, y := 1, 1
	var sum int
	for i := 0; i < n; i++ {
		sum = (x + y) % 1000000007
		x = y
		y = sum
	}
	return x
}
