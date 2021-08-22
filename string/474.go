package string

/**
二进制字符串str和两个证书m和n
找出并且返回strs的最大子元素的大小，该子集中最多有m个0和n个1
*/

/**
s1: 所有的子集，找出其中满足条件的集合
s2:
*/
func findMaxForm(strs []string, m, n int) int {
	mp := make([][]int, 0)
	for i := 0; i <= m; i++ {
		l := make([]int, n+1) // n + 1
		mp = append(mp, l)
	}
	one := make([]int, len(strs))
	zero := make([]int, len(strs))
	for i, _ := range strs {
		for _, item := range strs[i] {
			if item == '0' {
				zero[i]++
			} else {
				one[i]++
			}
		}
	}
	for i, _ := range strs {
		for j := m; j >= zero[i]; j-- {
			for z := n; z >= one[i]; z-- {
				mp[j][z] = max(mp[j-zero[i]][z-one[i]]+1, mp[j][z])
			}
		}
	}
	return mp[m][n]
}
