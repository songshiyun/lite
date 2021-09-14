package array

// prefix_sum

//nums = [1,1,1], k = 2
// prefixSum[n+1] -> prefixSum[0] = 0, prefixSum[i] = prefixSum[i-1] + nums[i-1]
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	var count int
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	for i := 0; i < len(nums); i++ {
		for j := i+1; j <= len(nums); j++ {
			if preSum[j]-preSum[i] == k {
				count++
			}
		}
	}
	return count
}

func subarraySum1(nums []int, k int) int  {
	if len(nums) == 0 {
		return 0
	}
	orderMap := make(map[int]int)
	orderMap[0] = 1
	var sum,ans int
	for _,item := range nums {
		sum += item
		ans += orderMap[sum-k]
		orderMap[sum]++
	}
	return ans
}
