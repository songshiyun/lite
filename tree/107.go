package tree

/**
例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层序遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/
func levelOrderBottom(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	var reverseSlice [][]*TreeNode
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		var temp []*TreeNode
		for i := 0; i < length; i++ {
			first := queue[0]
			queue = queue[1:]
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			temp = append(temp, first)
		}
		reverseSlice = append(reverseSlice, temp)
	}

	for i := len(reverseSlice) - 1; i >= 0; i-- {
		data := reverseSlice[i]
		var temp []int
		for j := 0; j < len(data); j++ {
			temp = append(temp, data[j].Val)
		}
		res = append(res, temp)
	}
	return
}

func levelOrderBottom2(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			first := queue[0]
			queue = queue[1:]
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			temp = append(temp, first.Val)
		}
		res = append(res, temp)
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return
}
