package tree

const (
	White = iota
	Gray
)

type colorNode struct {
	Node  *TreeNode
	Color int
}

//中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	stack := []*colorNode{{Node: root, Color: White}}

	for len(stack) != 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if item.Node == nil {
			continue
		}
		if item.Color == White {
			stack = append(stack, &colorNode{Node: item.Node.Right, Color: White})
			stack = append(stack, &colorNode{Node: item.Node, Color: Gray})
			stack = append(stack, &colorNode{Node: item.Node.Left, Color: White})
		} else {
			res = append(res, item.Node.Val)
		}
	}
	return
}
