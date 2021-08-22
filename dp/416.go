package dp

/**
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
*/

func canPartition(nums []int) bool {
	// 0-1背包问题
	var sum int
	for _, item := range nums {
		sum += item
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	//dp[j]表示放入j个元素后的和
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[target] == target {
		return true
	}

	return false
}

// dfs 遍历所有元素，选择或者不选择
func canPartitionDfs(nums []int) bool {
	// 0-1背包问题
	var sum int
	for _, item := range nums {
		sum += item
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	var dfs func(currentSum, index int) bool

	dfs = func(currentSum, index int) bool {
		if currentSum > target || index > len(nums) {
			return false
		}
		if currentSum == target {
			return true
		}
		return dfs(currentSum+nums[index], index+1) || dfs(currentSum, index+1)
	}

	return dfs(0, 0)
}

func canPartitionWithDfsMemo(nums []int) bool {
	var sum int
	for _, item := range nums {
		sum += item
	}
	if sum%2 != 0 {
		return false
	}
	memos := make([][]int, len(nums))
	for i := 0; i < len(memos); i++ {
		memos[i] = make([]int, len(nums))
	}
	target := sum / 2
	var dfs func(currentSum, index int) bool

	dfs = func(currentSum, index int) bool {
		if currentSum > target || index > len(nums) {
			return false
		}
		if currentSum == target {
			return true
		}
		if memos[currentSum][index] != 0 {
			//当前已经计算过了
			return memos[currentSum][index] == 1
		}
		res := dfs(currentSum+nums[index], index+1) || dfs(currentSum, index+1)
		if res {
			memos[currentSum][index] = 1
		} else {
			memos[currentSum][index] = 2
		}
		return res
	}
	return dfs(0, 0)
}
