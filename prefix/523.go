package prefix

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	for i := 0; i < len(nums)-1; i++ {
		var prefix int
		for j := i; j < len(nums); j++ {
			prefix += nums[j]
			if prefix/k == 0 && (j-i+1) >= 2 {
				return true
			}
		}
	}
	return false
}
