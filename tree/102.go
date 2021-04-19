package tree

/**
示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		var t []int
		for i := 0; i < length; i++ {
			temp := queue[0]
			queue = queue[1:]
			t = append(t, temp.Val)
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		res = append(res, t)
	}
	return
}
