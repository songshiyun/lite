package tree

/**
给定二叉树根结点root，此外树的每个结点的值要么是 0，要么是 1。

返回移除了所有不包含 1 的子树的原二叉树(不包含1，则子树的和为0)。

( 节点 X 的子树为 X 本身，以及所有 X 的后代。)

*/
//skill: 虚拟节点
func pruneTree(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) int
	//求二叉树节点和
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if left == 0 { //左子树和为0，移除
			root.Left = nil
		}
		if right == 0 { //右子树和为0，移除
			root.Right = nil
		}
		return root.Val + left + right //返回当前子树的和
	}
	dummy := new(TreeNode) //由于根节点可能会变化，所以用一个虚拟节点代替
	dummy.Left = root
	dfs(dummy)
	return dummy.Left
}
