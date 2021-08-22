package tree

//https://www.cnblogs.com/anniekim/archive/2013/06/15/morristraversal.html
//mirros遍历： O(n)时间复杂的，O(1)空间复杂度
func inorderMorrisTraversal(root *TreeNode) []int {
	var res []int
	curr := root
	var pre *TreeNode
	for curr != nil {
		if curr.Left == nil {
			res = append(res, curr.Val)
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Left != nil && pre.Right != curr {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = curr
				curr = curr.Left
			} else {
				pre.Right = nil
				res = append(res, curr.Val)
				curr = curr.Right
			}
		}
	}
	return res
}
