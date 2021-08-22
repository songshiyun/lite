package tree

/**
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。
*/
func sumNumbers(root *TreeNode) (res int) {
	nodeQueue := []*TreeNode{root}
	tempQueue := []int{0}
	for len(nodeQueue) != 0 {
		lastNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		lastValue := tempQueue[0]
		tempQueue = tempQueue[1:]
		currentVal := lastValue*10 + lastNode.Val
		if lastNode.Left == nil && lastNode.Right == nil {
			res += currentVal
			continue
		}
		if lastNode.Left != nil {
			nodeQueue = append(nodeQueue, lastNode.Left)
			tempQueue = append(tempQueue, currentVal)
		}
		if lastNode.Right != nil {
			nodeQueue = append(nodeQueue, lastNode.Right)
			tempQueue = append(tempQueue, currentVal)
		}
	}
	return
}

func sumNumbers1(root *TreeNode) (res int) {
	if root == nil {
		return
	}
	var dfs func(root *TreeNode, temp int)
	dfs = func(root *TreeNode, temp int) {
		if root == nil {
			return
		}
		if root.Right == nil && root.Left == nil {
			res += temp
			return
		}
		next := temp*10 + root.Val
		dfs(root.Left, next)
		dfs(root.Right, next)
	}
	dfs(root, 0)
	return
}
