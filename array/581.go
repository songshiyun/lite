package array

import "sort"

/**
Input: nums = [2,6,4,8,10,9,15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9]
in ascending order to make the whole array sorted in ascending order.
*/
// 2 6，4-> 2,4 8,10,9
// 单调栈
func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	left, right := len(nums)-1, 0
	var stack1 []int
	for i := 0; i < len(nums); i++ {
		if len(stack1) == 0 || nums[i] >= nums[stack1[len(stack1)-1]] {
			stack1 = append(stack1, i)
		} else {
			pop := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			left = min(left, pop)
			i--
		}
	}
	var stack2 []int
	for i := len(nums) - 1; i >= 0; i-- {
		if len(stack2) == 0 || nums[i] <= nums[stack2[len(stack2)-1]] {
			stack2 = append(stack2, i)
		} else {
			pop := stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
			right = max(right, pop)
			i++ // 加回去能避免while循环
		}
	}
	if right-left > 0 {
		return right - left + 1
	}
	return 0
}

// 将数组排序，找到正确的位置
func findUnsortedSubarray1(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	left, right := len(nums)-1, 0
	for i, v := range sorted {
		if nums[i] != v {
			left = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if sorted[i] != nums[i] {
			right = i
			break
		}
	}
	if right-left > 0 {
		return right - left + 1
	}
	return 0
}
