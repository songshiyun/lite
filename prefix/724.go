package prefix

//前缀和
//https://github.com/chefyuan/algorithm-base/blob/main/animation-simulation/%E5%89%8D%E7%BC%80%E5%92%8C/leetcode724%E5%AF%BB%E6%89%BE%E6%95%B0%E7%BB%84%E7%9A%84%E4%B8%AD%E5%BF%83%E7%B4%A2%E5%BC%95.md
//presum[i] = nums[i] + presum[i-1]
func pivotIndex(nums []int) int {
	var sum int
	for _, item := range nums {
		sum += item
	}
	var left int
	for i := 0; i < len(nums); i++ {
		if left == sum-nums[i]-left { //左边等于右边
			return i
		}
		left = left + nums[i]
	}
	return -1
}
