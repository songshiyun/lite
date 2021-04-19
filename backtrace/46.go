package backtrace

//没有重复数字的全排列

func permute(nums []int) (ans [][]int) {
	var list []int
	var dfs func(list []int, used map[int]bool)
	dfs = func(list []int, used map[int]bool) {
		if len(list) == len(nums) {
			ans = append(ans, append([]int(nil), list...))
			return
		}
		for i := 0; i < len(nums); i++ {
			//剪枝条件：当前这个元素已经被选择过
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
	dfs(list, map[int]bool{})
	return
}
