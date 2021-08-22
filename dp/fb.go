package dp

//fn(n) = fn(n-1)+fn(n-2)
//fn(0) = 0, fn(1) = 1

//dp[i] = dp[i-1] + dp[i-2]
//dp[0] = 0, dp[1] = 1
func fb1(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fb2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	// 求当前值只需要前两个值，并不需要记录整个路径上的值
	for i := 2; i <= n; i++ {
		temp := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = temp
	}
	return dp[1]
}

//dp数组含义
//状态转移方程
//dp数组初始化
//遍历顺序
//举例推倒dp数组
