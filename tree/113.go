package tree

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点
 */
func pathSum1(root *TreeNode, targetSum int) (res [][]int) {
	var dfs func(current *TreeNode, path []int, remain int)
	dfs = func(current *TreeNode, path []int, remain int) {
		if current == nil {
			return
		}
		path = append(path, current.Val)
		if current.Left == nil && current.Right == nil { //叶子节点
			if remain == current.Val {
				temp := make([]int,len(path))
				copy(path,temp)
				res = append(res, temp)
				return
			}
		}
		dfs(current.Left, path, remain-current.Val)
		dfs(current.Right, path, remain-current.Val)
		path = path[:len(path)-1] //找不到需要出栈
	}
	dfs(root, []int{}, targetSum)
	return
}
