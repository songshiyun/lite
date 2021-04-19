package tree

func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	stack := []*colorNode{{Node: root, Color: White}}
	for len(stack) != 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last.Node == nil {
			continue
		}
		if last.Color == White {
			stack = append(stack, &colorNode{Node: last.Node.Right, Color: White})
			stack = append(stack, &colorNode{Node: last.Node.Left, Color: White})
			stack = append(stack, &colorNode{Node: last.Node, Color: Gray}) //前序遍历，根结点先出栈
		} else {
			res = append(res, last.Node.Val)
		}
	}
	return
}
