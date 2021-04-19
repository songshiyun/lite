package dfs
//dp[k]是经过k调边能够到达的最小的值
//dp[i][k] 从原节点经过k调边到达的值，i节点的最小值
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//                 src,dst,weight
//src = 0, dst = 2, k = 1
/**
n 范围是 [1, 100]，城市标签从 0 到 n - 1
航班数量范围是 [0, n * (n - 1) / 2]
每个航班的格式 (src, dst, price)
每个航班的价格范围是 [1, 10000]
k 范围是 [0, n - 1]
航班没有重复，且不存在自环
*/
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	edges := make(map[int]map[int]int)
	for _, flight := range flights {
		if _, ok := edges[flight[0]]; !ok {
			edges[flight[0]] = make(map[int]int)
		}

		edges[flight[0]][flight[1]] = flight[2]
	}

	result := -1
	var dfs func(current, depth, cost int, used map[int]bool)
	dfs = func(current, depth, cost int, used map[int]bool) {
		if result != -1 && cost > result {
			return
		}

		if depth <= K {
			for next, weight := range edges[current] {
				if next != dst {
					used[next] = true
					dfs(next, depth+1, cost+weight, used)
					used[next] = false
				} else {
					if result == -1 {
						result = cost + weight
					} else {
						result = min(result, cost+weight)
					}
				}
			}
		}
	}
	dfs(src, 0, 0, map[int]bool{src: true})
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
