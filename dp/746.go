package dp

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// todo ???
func minCostClimbingStairs2(cost []int) int {
	dp := make([]int, 2)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dpi := min(dp[1], dp[0]+cost[i])
		dp[0] = dp[1]
		dp[1] = dpi
	}
	return min(dp[0], dp[1])
}
