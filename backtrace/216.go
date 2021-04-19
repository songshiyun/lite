package backtrace

//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字
/**
说明：
所有数字都是正整数。
解集不能包含重复的组合。
示例 1:
输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:
输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
*/

func combinationSum3(k int, n int) (ans [][]int) {
	var dfs func(start, sum int, comb []int)
	dfs = func(start, sum int, comb []int) {
		if len(comb) == k {
			if sum == n {
				ans = append(ans, append([]int(nil), comb...))
			}
			return
		}
		for i := start; i <= 9; i++ {
			comb = append(comb, i)
			sum += i
			dfs(i+1, sum, comb)
			sum -= i
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0, 0, []int{})
	return
}
