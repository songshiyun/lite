package backtrace

import "sort"

func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	var t []int
	var dfs func(bool2 bool, int2 int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur] == nums[cur-1] {
			return
		}
		t = append(t, nums[cur])
		dfs(true, cur+1)
		t = t[:len(t)-1] //移除最后选择的一个
	}
	dfs(false, 0)
	return
}

func subsetsWithDup2(nums []int) (ans [][]int) {
	sort.Ints(nums)
	var t []int
	var dfs func([]int, int)
	dfs = func(sl []int, curr int) {
		dst := make([]int, len(sl))
		copy(dst, sl)
		ans = append(ans, dst)
		for i := curr; i < len(nums); i++ {
			if i != curr && nums[i] == nums[i-1] {
				continue
			}
			sl = append(sl, nums[i])
			dfs(sl, i+1)
			sl = sl[:len(sl)-1]
		}
	}
	dfs(t, 0)
	return
}
