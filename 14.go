package sto

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	if n == 2 {
		return 1
	}
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = maxInt(dp[i], maxInt(j*dp[i-j], j*(i-j)))
		}
	}
	return dp[n]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
