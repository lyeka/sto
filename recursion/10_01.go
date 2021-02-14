package recursion

func fib(n int) int {
	x, y := 0, 1
	sum := x
	for i := 0; i < n; i++ {
		sum = (x + y) % 1000000007
		x = y
		y = sum
	}
	return x
}
