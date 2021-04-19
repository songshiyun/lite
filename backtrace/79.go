package backtrace

func exist(board [][]byte, word string) bool {
	var dfs func([][]bool, int, int, int) bool
	var inArea func(i int, j int) bool
	inArea = func(i int, j int) bool {
		return i >= 0 && i < len(board[0]) && j >= 0 && j < len(board[0])
	}
	dfs = func(marked [][]bool, i int, j int, index int) bool {
		if index == len(word)-1 {
			return board[i][j] == word[index]
		}
		if board[i][j] == word[index] {
			//anFindRest := canFind(r+1, c, i+1) || canFind(r-1, c, i+1) || canFind(r, c+1, i+1) || canFind(r, c-1, i+1)
			marked[i][j] = true
			for k := 0; k < 4; k++ {
				newX := i + direction[k][0]
				newY := j + direction[k][1]
				if inArea(newX, newY) && !marked[newX][newY] {
					if dfs(marked, newX, newY, index+1) {
						return true
					}
				}
			}
			marked[i][j] = false
		}
		return false
	}
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	marked := make([][]bool, m)
	for i := 0; i < m; i++ {
		marked[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(marked, i, j, 0) {
				return true
			}
		}
	}
	return false
}

var direction = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}
