package backtrace

import "sort"

/**
给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的每个数字在每个组合中只能使用一次。
说明：
所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例1:
输入: candidates =[10,1,2,7,6,1,5], target =8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例2:
输入: candidates =[2,5,2,1,2], target =5,
所求解集为:
[
	[1,2,2],
	[5]
]
*/
//1.元素不能重复使用
//2.组合不能重复 [2,5]和[5,2]是同一个组合
//

//当给定的元素重复的时候，先将元素排序，方便去重
//for 枚举出选项时，加入下面判断，从而忽略掉同一层重复的选项，避免产生重复的组合。比如[1,2,2,2,5]，
//选了第一个 2，变成 [1,2]，它的下一选项也是 2，跳过它，因为如果选它，就还是 [1,2]。
//if (i - 1 >= start && candidates[i - 1] == candidates[i]) {
//    continue;
//}
//当前选择的数字不能和下一个选择的数字重复，给子递归传i+1，避免与当前选的i重复。
//dfs(i + 1, temp, sum + candidates[i])

func combinationSum2(candidates []int, target int) (ans [][]int) {
	var dfs func(start, sum int, comb []int)
	dfs = func(start, sum int, comb []int) {
		if sum >= target {
			if sum == target {
				ans = append(ans, append([]int(nil), comb...))
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			//排序之后，保证同一层不会选到重复元素
			if i-1 >= start && candidates[i] == candidates[i-1] {
				continue
			}
			comb = append(comb, candidates[i])
			dfs(i+1, sum+candidates[i], comb) //不能重复选
			comb = comb[:len(comb)-1]
		}
	}
	sort.Ints(candidates)
	dfs(0, 0, []int{})
	return
}
