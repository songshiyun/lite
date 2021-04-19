package backtrace

import "sort"

//含重复元素的全排列

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	var t []int
	var dfs func(list []int, used map[int]bool)
	dfs = func(list []int, used map[int]bool) {
		if len(list) == len(nums) {
			ans = append(ans, append([]int(nil), list...))
			return
		}
		for i := 0; i < len(nums); i++ {
			//剪枝条件
			if i-1 >= 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			list = append(list, nums[i])
			used[i] = true
			dfs(list, used)
			list = list[:len(list)-1]
			used[i] = false
		}
	}
	dfs(t, map[int]bool{})
	return
}
