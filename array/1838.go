package array

import "sort"

// todo important
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	l, ans := 0, 0
	var sum int64
	for r := 0; r < len(nums); r++ {
		sum = sum + int64(nums[r])
		for l < r && (sum+int64(k)) < int64(nums[r]*(r-l+1)) {
			sum -= int64(nums[l])
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}
