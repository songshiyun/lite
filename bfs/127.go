package bfs

//BFS的普适题型一共是三种：树的层次遍历，图的最短路径，字符串匹配最少次数。找准关键字“最短路径”，“层次遍历”。
func ladderLength(beginWord string, endWord string, wordList []string) int {
	visited := make(map[string]bool)
	for _, item := range wordList {
		visited[item] = true
	}
	queue := []string{beginWord}
	step := 1
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			if current == endWord {
				return step
			}
			for k := 0; k < len(current); k++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := current[:k] + string(j) + current[k+1:]
					if visited[newWord] {
						queue = append(queue, newWord)
						visited[newWord] = false
					}
				}
			}
		}
		step++
	}
	return 0
}
