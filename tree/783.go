package tree

import "math"

func minDiffInBST(root *TreeNode) int {
	var dfs func(root *TreeNode,lower,upper int) int
	dfs = func(root *TreeNode, lower, upper int) int {
		if root == nil {
			return upper -lower
		}
		left := dfs(root.Left,lower,root.Val)
		right := dfs(root.Right,root.Val,upper)
		return min(left,right)
	}

	return dfs(root,-math.MaxInt32,-math.MaxInt32)
}

func min(i,j int) int {
	if i < j {
		return i
	}
	return j
}