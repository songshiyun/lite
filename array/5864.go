package array


func numberOfWeakCharacters(properties [][]int) int {
	var count int
	mp := make(map[int][][]int)
	for _,item := range properties {
		sum := item[0]+item[1]
		mp[sum] = append(mp[sum], item)
	}
	for _,item := range properties {
		if muchBigger(item,mp) {
			count++
		}
	}
	return count
}

func muchBigger(item []int,mp map[int][][]int) bool {
	total := item[0] + item[1]
	for sum,s := range mp {
		if sum > total {
			if countainItem(item,s) {
				return true
			}
		}
	}
	return false
}

func countainItem(item []int,properties [][]int) bool {
	for _,v := range properties {
		if v[0] > item[0] && v[1] > item[1] {
			return true
		}

	}
	return false

}
