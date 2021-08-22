package tree

func pathSum(root *TreeNode, target int) (res [][]int) {
	if root == nil {
		return
	}
	var dfs func(cur *TreeNode, temp []int, remain int)
	dfs = func(cur *TreeNode, temp []int, remain int) {
		if cur == nil {
			return
		}
		//左右节点同时为空
		temp = append(temp, cur.Val)
		if cur.Left == nil && cur.Right == nil {
			if cur.Val == remain {
				t := make([]int, len(temp)) //tmp指针会被外层修改，这里需要拷贝
				copy(t, temp)
				res = append(res, t)
			}
			return
		}
		dfs(cur.Left, temp, remain-cur.Val)
		dfs(cur.Right, temp, remain-cur.Val)
		temp = temp[:len(temp)-1]
	}
	dfs(root, []int{}, target)
	return
}
