package dp

func maximumScore(nums []int, multipliers []int) int {
	// 这是一个子问题求解问题，不能用贪心，局部最优不一定是全局最优
	minInt := -1 << 31
	n := len(nums)
	m := len(multipliers)
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, m)
		for j := range cache[i] {
			cache[i][j] = minInt
		}
	}
	var dpFunc func(i, j int) int
	dpFunc = func(i, j int) int {
		k := n - (j - i + 1)
		if k == m {
			return 0
		}
		ans := cache[i][k]
		if ans != minInt {
			return ans
		}
		i1 := dpFunc(i+1,j) + nums[i] * multipliers[k]
		j1 := dpFunc(i,j-1) + nums[j] * multipliers[k]
		ans = Max(i1,j1)
		cache[i][k] = ans
		return ans
	}
	return dpFunc(0, n-1)
}

/**

dp(i,j,m)-> i 开始，j结尾的乘以m的最大值，
dp(i,j,m) -> max(nums[i] * m + dp[i+1,j,m+1],nums[j] * m + dp[i,j-1,m+1])
i,j,m是啥有关联的，三维可以简化为二维
*/

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
