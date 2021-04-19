package heap

func kthLargestValue(matrix [][]int, k int) int {
	n := len(matrix)
	m := len(matrix[0])
	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	minHeap := NewIntHeap()
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ { //有减1这样的操作，可以多初始化一个，把[0,0]空出来避免边界处理
			if i-1 >= 0 && j-1 >= 0 {
				arr[i][j] = arr[i][j-1] ^ matrix[i-1][j-1]
				arr[i][j] = arr[i][j] ^ arr[i-1][j-1]
				arr[i][j] = arr[i][j] ^ arr[i-1][j]
				minHeap.Push(arr[i][j])
				if minHeap.Len() > k {
					minHeap.Poll()
				}
			}
		}
	}
	if minHeap.Len() == k {
		return minHeap.Poll()
	}
	return 0
}
