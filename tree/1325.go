package tree

//给你一棵以root为根的二叉树和一个整数 target ，请你删除所有值为 target 的 叶子节点 。
//
//注意，一旦删除值为target的叶子节点，它的父节点就可能变成叶子节点；如果新叶子节点的值恰好也是 target ，那么这个节点也应该被删除。
//
//也就是说，你需要重复此过程直到不能继续删除。
//叶子节点移除以后，新的父节点也可能变成叶子节点被移除，递归重复调用
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var dfs func(current, parent *TreeNode, isLeft bool)
	dfs = func(current, parent *TreeNode, isLeft bool) {
		if current == nil {
			return
		}
		//后续遍历
		dfs(current.Left, current, true)
		dfs(current.Right, current, false)
		if current.Val == target && parent != nil && current.Left == nil && current.Right == nil {
			if isLeft {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
	}
	dummy := new(TreeNode) //由于根节点可能改变，增加一个虚拟节点
	dummy.Left = root
	dfs(root, dummy, true)
	return dummy.Left
}
