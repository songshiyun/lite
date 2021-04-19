package backtrace

func generateParenthesis(n int) (ans []string) {
	var dfs func(sub string, left, right int)
	dfs = func(sub string, left, right int) {
		if left == 0 && right == 0 {
			ans = append(ans, sub)
			return
		}
		//只要还有（剩下，就可以选，
		//只有当")"剩余比"("多的时候才能选择")"
		if left > 0 {
			dfs(sub+"(", left-1, right)
		}
		//括号的组合一定是（开始，则一定右边的剩余应该多余左边
		if right > left {
			dfs(sub+")", left, right-1)
		}
	}
	dfs("", n, n)
	return
}
