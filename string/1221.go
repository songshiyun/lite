package string

func balancedStringSplit(s string) int {
	if len(s) <= 1 {
		return 0
	}
	var count int
	var temp []byte
	var last byte
	var presize int
	var nextSize int
	for _, b := range []byte(s) {
		if len(temp) == 0 {
			temp = append(temp, b)
			presize++
			last = b
		} else {
			if last == b {
				temp = append(temp, b)
				presize++
				last = b
			} else {
				nextSize++
				if nextSize == presize {
					nextSize = 0
					presize = 0
					temp = temp[:0]
					count++
				}else if temp[len(temp)-1] == b {
					last = b
				}
			}
		}
	}
	return count
}
