package tree

/**
z字形遍历
 */
func zigzagLevelOrder(root *TreeNode) (res [][]int) {
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
		if len(res)%2 == 0 {
			reverse(temp)
		}
	}
	return
}

func reverse(data []int) {
	i, j := 0, len(data)-1
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}
