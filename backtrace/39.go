package backtrace

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
	[2,2,2,2],
	[2,3,3],
	[3,5]
]
 */
func combinationSum(candidates []int, target int) (ans [][]int) {
	var dfs func(comb []int,start, sum int)
	dfs = func(comb []int,start,sum int) {
		if sum >= target { //剪枝条件
			if sum == target {
				ans = append(ans, append([]int(nil), comb...))
			}
			return
		}
		//不产生重复元素：[223]和[322]是同一个组合
		//只要限制下一次选择的起点，是基于本次的选择，
		//这样下一次就不会选到本次选择同层左边的数。
		//即通过控制 for 遍历的起点，去掉会产生重复组合的选项。
		for i := start; i < len(candidates); i++ {
			comb = append(comb, candidates[i])
			sum += candidates[i]
			dfs(comb,i, sum) //由于可以重复选择，这里是i,不是i+1
			comb = comb[:len(comb)-1]
			sum -= candidates[i]
		}
	}
	dfs([]int{}, 0,0)
	return
}
