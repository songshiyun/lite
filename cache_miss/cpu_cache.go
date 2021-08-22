package cache_miss

var matrix [][]int

func init() {
	matrix = make([][]int, 128)
	for index, _ := range matrix {
		matrix[index] = make([]int, 128)
	}
}

func PrintRow() {
	row := len(matrix)
	coloumn := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < coloumn; j++ {
			_ = matrix[i][j]
		}
	}
}

func PrintColumn() {
	row := len(matrix)
	coloum := len(matrix[0])
	for i := 0; i < coloum; i++ {
		for j := 0; j < row; j++ {
			_ = matrix[i][j]
		}
	}
}
