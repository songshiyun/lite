package tree

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) != 0 {
		temp := queue
		queue = nil
		for i, current := range temp {
			if i+1 < len(temp) {
				current.Next = temp[i+1]
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}
	return root
}

func connect1(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) != 0 {
		size := len(queue)
		var pre *Node
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			if pre != nil {
				pre.Next = current
			}
			pre = current
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
