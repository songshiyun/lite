package dp

// 给定一个整数，拆解成n个数的和，使得n个数的乘积最大

func integerBreaker(n int) int {

	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
		}
	}
	return dp[n]
}
