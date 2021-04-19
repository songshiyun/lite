package tree

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var dfs func(root *LeveWithNode) *LeveWithNode
	dfs = func(root *LeveWithNode) *LeveWithNode {
		current := root.Node
		if current == nil {
			return root
		}
		left := dfs(&LeveWithNode{Node: current.Left, Level: root.Level + 1})
		right := dfs(&LeveWithNode{Node: current.Right, Level: root.Level + 1})
		if left == right {
			return &LeveWithNode{Node: current, Level: left.Level}
		}
		if left.Level > right.Level {
			return &LeveWithNode{left.Node, left.Level}
		}
		return &LeveWithNode{right.Node, right.Level}
	}
	return dfs(&LeveWithNode{root, -1}).Node
}

type LeveWithNode struct {
	Node  *TreeNode
	Level int
}
