package dp

/**
给你一个整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

示例 1：

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
*/
/**

能够凑成target, 则满足:
x-y = target
x+y = sum
则： x = (sum + target) /2 ,则是否能够分成两部分满足重量为x
dp[j] 表示凑成质量为j的方法
dp[0] = 1,表示凑成质量为0的有几种方法，1种
dp[j] 表示:填满j(包括j)这么大容积的包，有dp[i]种方法
*/
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, item := range nums {
		sum += item
	}
	if target > sum || (sum+target)%2 == 1 {
		return 0
	}
	t := (sum + target) / 2
	dp := make([]int, t+1)
	dp[0] = 1 //
	// dp[i][j]
	for i := 0; i < len(nums); i++ {
		for j := t; j >= nums[i]; j-- {
			dp[j] = dp[j] + dp[j-nums[i]]
		}
	}
	return dp[t]
}
