package tree

import "math"

func maxPathSum(root *TreeNode) int {
	res := -math.MaxInt32
	var subMax func(current *TreeNode) int
	subMax = func(current *TreeNode) int {
		if current == nil {
			return 0
		}
		left := max(subMax(current.Left), 0)   //选择左边节点或者不选
		right := max(subMax(current.Right), 0) //选择右边节点或者不选，能够有正向的值就选
		newValue := left + right + current.Val //选择(不)左右和当前的值与已经存在的值对比
		res = max(res, newValue)
		return max(left, right) + current.Val
	}
	subMax(root)
	return res
}
