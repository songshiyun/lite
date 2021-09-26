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
	var index int
	l := len(prefix[0])
	// j是转弯的位置
	for j := 0; j < len(prefix[0]); j++ {
		var tmp int
		if j ==0 {
			tmp = prefix[0][j] + prefix[1][l-1]
		}else {
			tmp = prefix[0][j] + prefix[1][l-1] - prefix[1][j-1]
		}
		if tmp >= res {
			res = tmp
			index = j
		}
	}
	if index !=  0 {
		right := prefix[0][l-1] - prefix[0][index]
		left := prefix[1][index-1]
		return int64(max2(right,left))
	}
	return int64(prefix[0][l-2] - prefix[0][0])
}

func max2(i,j int) int {
	if i > j {
		return i
	}
	return j
}