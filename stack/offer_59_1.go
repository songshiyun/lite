package stack

//滑动窗口内的最大值
//滑动窗口的第一个元素为当前滑动窗口内的最大值
func maxSlidingWindow(nums []int, k int) []int {
	var res, temp []int
	if len(nums) == 0 || k == 0 {
		return res
	}
	for i := 0; i < min(k, len(nums)); i++ {
		for len(temp) != 0 && temp[len(temp)-1] < nums[i] {
			temp = temp[:len(temp)-1]
		}
		temp = append(temp, nums[i])
	}
	res = append(res, temp[0])
	for j := k; j < len(nums); j++ {
		if nums[j-k] == temp[0] { //todo for what ?
			temp = temp[1:]
		}
		for len(temp) != 0 && temp[len(temp)-1] < nums[j] {
			temp = temp[:len(temp)-1]
		}
		temp = append(temp, nums[j])
		res = append(res, temp[0])
	}

	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
