package array

//前缀和
func maxSideLength(mat [][]int, threshold int) int {
	row, column := len(mat), len(mat[0])
	var dp [][]int
	dp = make([][]int, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]int, column+1)
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			dp[i][j] = mat[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	var result int
	for k := 0; k <= min(row, column); k++ {
		for i := 1; i <= row; i++ {
			for j := 1; j <= column; j++ {
				if i-k < 0 || j-k < 0 {
					continue
				}
				tmp := dp[i][j] - dp[i-k][j] - dp[i][j-k] + dp[i-k][j-k]
				if tmp <= threshold {
					result = max(result, k)
				}
			}
		}
	}
	return result
}

//前缀+二分
func maxSideLength2(mat [][]int, threshold int) int {
	row, column := len(mat), len(mat[0])
	dp := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]int, column+1)
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			// dp[i][j] -> mat[i-1][j-1]
			dp[i][j] = mat[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	var helper func(int, int) bool
	helper = func(mid int, t int) bool {
		for i := 1; i <= row; i++ {
			for j := 1; j <= column; j++ {
				if i < mid || j < mid {
					continue
				}
				if dp[i][j]-dp[i-mid][j]-dp[i][j-mid]+dp[i-mid][j-mid] <= t {
					return true
				}
			}
		}
		return false
	}
	l, h := 0, min(row, column)
	for l <= h {
		mid := l + (h-l)/2
		if l == h || l+1 == h {
			break
		}
		if helper(mid, threshold) {
			l = mid
		} else {
			h = mid - 1
		}
	}
	if helper(h, threshold) {
		return h
	}
	return l
}
