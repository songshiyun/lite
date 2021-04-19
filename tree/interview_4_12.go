package tree

/**
给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。
注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。
示例:
给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
3
解释：和为 22 的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]
*/
//todo
func pathSum2(root *TreeNode, sum int) int {
	var dfs func(root *TreeNode, path []int) (result int)
	dfs = func(root *TreeNode, path []int) (result int) {
		if root != nil {
			s := make([]int, len(path)+1)
			copy(s, path)
			for i := range s {
				s[i] += root.Val
				if s[i] == sum {
					result++
				}
			}
			result += dfs(root.Left, s) + dfs(root.Right, s)
		}
		return
	}
	return dfs(root, []int{})
}
//双递归
func pathSum3(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode, remain int) (sum int)
	dfs = func(root *TreeNode, remain int) (sum int) {
		if root == nil {
			return 0
		}
		if remain == root.Val {
			sum++
		}
		return sum + dfs(root.Left, remain-root.Val) + dfs(root.Right, remain-root.Val)
	}
	return dfs(root, sum) + pathSum3(root.Left, sum) + pathSum3(root.Right, sum)
}
