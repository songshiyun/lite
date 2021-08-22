package tree

//bfs, 找到每一层最左边的
/**
给定一个二叉树，在树的最后一行找到最左边的值。
示例 1:
输入:
    2
   / \
  1   3
输出:
1
示例 2:
输入:
        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7
输出:
7

*/
func findBottomLeftValue(root *TreeNode) (res int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			if i == 0 { //每一层的第一个
				res = current.Val
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}
	return
}
