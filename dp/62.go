package dp

func uniquePaths(m int, n int) int {
	// 从[0,0]到[i,j]的路径数
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}

//可以用一维dp, 滚动数组
// dp[i] = dp[i] + dp[i-1]
func uniquePath2(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			dp[i] += dp[i-1]
		}
	}
	return dp[n-1]
}
