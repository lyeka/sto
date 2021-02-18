package sto

// dfs
func movingCount_dfs(m int, n int, k int) int {
	d := [][]int{}
	for i:= 0; i<m; i++ {
		s := []int{}
		for j:= 0; j<n; j++ {
			s = append(s, 0)
		}
		d = append(d, s)
	}

	return dfs_2(m, n, k, 0, 0, d)
}

func dfs_2(m, n, k, x, y int, d [][]int) int {
	if x >= m || y >= n || d[x][y] == 1  || getSum(x) + getSum(y) > k {
		return 0
	}
	d[x][y] = 1
	return 1 + dfs_2(m, n, k, x+1, y, d) + dfs_2(m, n, k, x, y+1, d)
}



