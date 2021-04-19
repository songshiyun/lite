package tree

var res int

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode, k *int)
	dfs = func(root *TreeNode, k *int) {
		if root != nil {
			dfs(root.Left, k)
			*k--
			if *k == 0 {
				res = root.Val
				return
			}
			dfs(root.Right, k)
		}
	}
	dfs(root, &k)
	return res
}
