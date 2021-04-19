package backtrace

/**
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入:n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
func combine(n int, k int) (ans [][]int) {
	var dfs func(comb []int, start int)
	dfs = func(comb []int, start int) {
		if len(comb) == k {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		//不能重复选择
		for i := start; i <= n; i++ {
			comb = append(comb, i)
			dfs(comb, i+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs([]int{}, 1)
	return
}
