package stack

// 单调栈，记录的是数组元素的index
func findPeakElement(nums []int) int {
	var minStack []int
	for i := range nums {
		if len(minStack) == 0 {
			minStack = append(minStack, i)
			continue
		}
		if nums[minStack[len(minStack)-1]] < nums[i] {
			if i == len(nums)-1 {
				return len(nums) - 1
			}
			minStack = append(minStack, i)
			continue
		}
		return i - 1
	}
	return 0
}
