package tree

import "math"

//二叉搜索树，终须终须中序遍历有序
func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	var pre *TreeNode
	stack := []*colorNode{{root, White}}
	for len(stack) != 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last.Node == nil {
			continue
		}
		if last.Color == White {
			stack = append(stack, &colorNode{last.Node.Right, White})
			stack = append(stack, &colorNode{last.Node, Gray})
			stack = append(stack, &colorNode{last.Node.Left, White})
		} else {
			if pre != nil {
				if pre.Val >= last.Node.Val {
					return false
				}
			}
			pre = last.Node
		}
	}
	return true
}

func isValidBST1(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}
