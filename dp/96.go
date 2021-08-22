package dp

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	// dp[i] 表示有n个节点的二叉树
	// dp[0] = 1
	// dp[1]表示有一个节点的数
	// dp[2]表示以1为头节点的数+以2为头节点的数
	// 有三个节点表示 1为头节点+2为头节点+3为头节点
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
