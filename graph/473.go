package graph

func networkDelayTime(times [][]int, n int, k int) int {
	inf := 0x3f3f3f3f
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
		for j, _ := range matrix[i] {
			matrix[i][j] = inf
		}
	}
	for _, item := range times {
		matrix[item[0]-1][item[1]-1] = item[2]
	}
	dist := make([]int, n)
	for index, _ := range dist {
		dist[index] = inf
	}
	visited := make([]bool, n)
	dist[k-1] = 0
	for i := 0; i < n-1; i++ {
		tmp := -1
		for j, u := range visited {
			if !u && (tmp == -1 || dist[j] < dist[tmp]) {
				tmp = j
			}
		}
		visited[tmp] = true
		for t, dis := range matrix[tmp] {
			dist[t] = min(dist[t], dis+dist[tmp])
		}
	}
	var ans int
	for _, d := range dist {
		if d == inf {
			return -1
		}
		if ans > d {
			ans = d
		}
	}
	return ans
}

func networkDelayTime1(times [][]int, n int, k int) int {
	maxPath := 100 * 101
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		for j, _ := range dp[i] {
			dp[i][j] = maxPath
		}
	}
	for _, item := range times {
		dp[item[0]-1][item[1]-1] = item[2]
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 0
	}
	// dp[i][j]表示从节点i->j节点的最小值
	for x := 0; x < n; x++ {
		// 对于每一个节点，
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dp[i][j] = min(dp[i][j], dp[i][x]+dp[x][j])
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[k-1][i])
	}
	if res == maxPath {
		return -1
	}
	return res

}

// bellman ford

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
