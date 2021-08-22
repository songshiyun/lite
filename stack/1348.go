package stack

/**
给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，
该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。
如果不存在满足条件的子数组，则返回 0 。

*/
/**
维护两个单调栈：
1.单调递减栈 (第一个元素最大)
2.单调递增栈（第一个元素最小）
滑动窗口内的最大差为两个栈顶元素的差，如果两个栈顶元素的差大于limit, 如果num[left] == stack[0],则将当前栈顶元素移除，否则移动左边指针

maxVal = max(maxVal,right-left+1)
right++
*/
func longestSubarray(nums []int, limit int) int {
	if len(nums) <= 1 {
		return 0
	}
	var maxQueue, minQueue []int
	left, right, maxValue := 0, 0, 0
	for right < len(nums) {
		for len(maxQueue) != 0 && maxQueue[len(maxQueue)-1] < nums[right] {
			maxQueue = maxQueue[:len(maxQueue)-1]
		}
		for len(minQueue) != 0 && minQueue[len(minQueue)-1] > nums[right] {
			minQueue = minQueue[:len(minQueue)-1]
		}
		maxQueue = append(maxQueue, nums[right])
		minQueue = append(minQueue, nums[right])
		for len(maxQueue) != 0 && len(minQueue) != 0 && maxQueue[0]-minQueue[0] > limit {
			if maxQueue[0] == nums[left] {
				maxQueue = maxQueue[1:]
			}
			if minQueue[0] == nums[left] {
				minQueue = minQueue[1:]
			}
			left++
		}
		maxValue = max(maxValue, right-left+1)
		right++
	}
	return maxValue
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
