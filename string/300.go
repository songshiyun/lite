package string

/**
给定一个蒸熟数组nums,找到其中子序列严格递增的长度
子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/

/**
1.暴力，找到所有递增的子序列，然后找出其中最长的子序列
2.dp
*/

//没有邀请连续递增
//状态转移方程： dp[i] = max(dp[j]+1,dp[i])
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	//dp[i]表示以i结尾的最长子序列的长度
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	result := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(dp[i], result)
	}
	return result
}

//最长连续递增子序列
func lengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	//dp[i]表示以i结尾的最长子序列的长度
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	result := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			} else {
				break
			}
		}
		result = max(dp[i], result)
	}
	return result
}

func lengthOfLCISII(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	result, counter := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			counter++
		} else {
			result = max(result, counter)
			counter = 1
		}
	}
	return max(result, counter)
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
