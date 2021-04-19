package backtrace

func subsets(nums []int) (ans [][]int) {
	var t []int
	var dfs func(curr int)
	dfs = func(curr int) {
		if curr == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		t = append(t, nums[curr])
		dfs(curr + 1)
		t = t[:len(t)-1]
		dfs(curr + 1)
	}
	dfs(0)
	return
}

func subsets2(nums []int) (ans [][]int) {
	var t []int
	var dfs func(curr int, list []int)
	dfs = func(curr int, list []int) {
		dst := make([]int, len(list))
		copy(dst, list)
		ans = append(ans, dst)
		for i := curr; i < len(nums); i++ {
			list = append(list, nums[i])
			dfs(i+1, list)
			list = list[:len(list)-1]
		}
	}
	dfs(0, t)
	return
}
