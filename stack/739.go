package stack

/**
请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，
至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用0 来代替。
例如，给定一个列表temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是[1, 1, 4, 2, 1, 1, 0, 0]。
提示：气温 列表长度的范围是[1, 30000]。每个气温的值的均为华氏度，都是在[30, 100]范围内的整数。
*/
//单调栈
//栈顶存储的是最大的元素的下标，
//遍历数组： 如果当前元素大于栈顶元素对应的值，则栈顶元素出栈，
// x = stack.pop()
// res[x] = i - x
//否则将当前i压入栈顶
func dailyTemperatures(T []int) []int {
	var temp, res []int
	res = make([]int, len(T))
	if len(T) == 0 {
		return res
	}
	for i := 0; i < len(T); i++ {
		for len(temp) != 0 && T[i] > T[temp[len(temp)-1]] {
			x := temp[len(temp)-1]
			res[x] = i - x
			temp = temp[:len(temp)-1]
		}
		temp = append(temp, i)
	}
	return res
}
