package tree

/**
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]

*/
func levelOrder1(root *Node1) (res [][]int) {

	return
}

type Node1 struct {
	Val      int
	Children []*Node1
}
