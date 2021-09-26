package array

func gridGame(grid [][]int) int64 {
	prefix := make([][]int, 2)
	for i := range prefix {
		prefix[i] = make([]int, len(grid[i]))
	}
	for i := range grid {
		for j := range grid[i] {
			if j == 0 {
				prefix[i][j] = grid[i][j]
			} else {
				prefix[i][j] = prefix[i][j-1] + grid[i][j]
			}
		}
	}
	var res int
	l := len(prefix[0])
	// j是转弯的位置
	for j := 0; j < len(prefix[0]); j++ {
		if j == 0 {
			res = max(res,prefix[0][l-1] - prefix[0][0])
		}else {
			right := prefix[0][l-1] - prefix[0][j]
			left := prefix[1][j-1]
			res = min2(res,max2(right,left))
		}

	}
	return int64(res)
}

func max2(i,j int) int {
	if i > j {
		return i
	}
	return j
}
func min2(i,j int) int {
	if i > j {
		return j
	}
	return i
}

func gridGame1(grid [][]int) int64 {
	var prefix1 int
	for _,item := range grid[0] {
		prefix1 += item
	}
	var res = 1 << 32 -1
	var prefix2 int
	for i,v := range grid[0] {
		prefix1 -= v
		res = min2(res,max2(prefix1,prefix2))
		prefix2 += grid[1][i]

	}
	return int64(res)
}