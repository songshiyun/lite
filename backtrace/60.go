package backtrace

import (
	"strconv"
	"strings"
)

/**
给出集合[1,2,3,...,n]，其所有元素共有 n! 种排列。
按大小顺序列出所有排列情况，并一一标记，当n = 3 时, 所有排列如下：
"123"
"132"
"213"
"231"
"312"
"321"
给定n 和k，返回第k个排列。
示例 1：
输入：n = 3, k = 3
输出："213"
示例 2：
输入：n = 4, k = 9
输出："2314"
示例 3：
输入：n = 3, k = 1
输出："123"
*/

//todo 超过时间限制
func getPermutation(n int, k int) (res string) {
	var count int
	var dfs func(sub []string, used map[int]bool) string
	dfs = func(sub []string, used map[int]bool) string {
		if len(sub) == n {
			count++
			if count == k {
				return strings.Join(sub, "")
			}
			return ""
		}
		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}
			sub = append(sub, strconv.Itoa(i))
			used[i] = true
			rs := dfs(sub, used)
			sub = sub[:len(sub)-1]
			used[i] = false
			if rs != "" {
				return rs
			}
		}
		return ""
	}
	res = dfs([]string{}, map[int]bool{})
	return
}
