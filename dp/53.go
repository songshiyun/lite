package dp

// dp[i] = num[i] + dp[i-1] (dp[i-1] > 0)
func maxSubArray(nums []int) int {
	var res int
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res = nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]

		}
		res = max(res, dp[i])
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, index := nums[0], 0, 0
	for index < len(nums) {
		res += nums[index]
		if res >= maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		index++
	}
	return maxSum
}
