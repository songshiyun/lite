package dfs

func findFarmland(land [][]int) [][]int {
	n, m := len(land), len(land[0])
	var res [][]int
	// 元素为，左边为边界或者0，右边为边界或者0
	for i, row := range land {
		for j, item := range row {
			if item == 0 || j > 0 && row[j-1] == 1 || i > 0 && land[i-1][j] == 1 {
				continue
			}
			ix := i
			for ; ix+1 < n && land[ix+1][j] == 1; ix++ {
			}
			jx := j
			for ; jx+1 < m && land[i][jx+1] == 1; jx++ {
			}

			res = append(res, []int{i, j, ix, jx})
		}
	}
	return res
}

var dir = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func findFarmlandWithDfs(land [][]int) [][]int {
	n, m := len(land), len(land[0])
	var dfs func(i, j, u, v int)
	var maxu, maxv int
	dfs = func(i, j, u, v int) {
		if u < 0 || v < 0 {
			return
		}
		if u >= n {
			maxu = n - 1
			return
		}
		if v >= m {
			maxv = m - 1
			return
		}
		if land[u][v] == 0 { // 当前是0
			if u > maxu {
				maxu = u - 1
			}
			if v > maxv {
				maxv = v - 1
			}
			return
		}
		land[u][v] = 0
		for _, d := range dir {
			dfs(i, j, u+d[0], v+d[1])
		}
	}
	var res [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if land[i][j] == 1 {
				maxu, maxv = i, j
				dfs(i, j, i, j)
				res = append(res, []int{i, j, maxu, maxv})
			}
		}
	}
	return res
}
