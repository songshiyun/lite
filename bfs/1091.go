package bfs

var direction = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func shortestPathBinaryMatrix(grid [][]int) (res int) {
	weight := len(grid)
	if grid[0][0] == 1 || grid[weight-1][weight-1] == 1 {
		return -1
	}
	if weight == 1 {
		return weight
	}
	var inArea func(x, y int) bool
	inArea = func(x, y int) bool {
		return x >= 0 && x < weight && y >= 0 && y < weight
	}
	visited := make([][]bool, weight)
	for i := 0; i < weight; i++ {
		visited[i] = make([]bool, weight)
	}
	var queue [][2]int
	queue = append(queue, [2]int{0, 0})
	visited[0][0] = true
	steps := 1
	for len(queue) != 0 {
		size := len(queue)
		for j := 0; j < size; j++ {
			current := queue[0]
			queue = queue[1:]
			for i := 0; i < 8; i++ { //八个方向
				newX := current[0] + direction[i][0]
				newY := current[1] + direction[i][1]
				if !inArea(newX, newY) {
					continue
				}
				if newX == weight-1 && newY == weight-1 {
					return steps + 1
				}
				if !visited[newX][newY] && grid[newX][newY] != 1 {
					visited[newX][newY] = true
					queue = append(queue, [2]int{newX, newY})
				}
			}
		}
		steps++
	}
	return -1
}
