package string

/*func main() {
	fmt.Println(checkOneSegment("10001"))
	fmt.Println(checkOneSegment("110"))
	fmt.Println(checkOneSegment("1111"))
	fmt.Println(checkOneSegment("0001"))
	fmt.Println(checkOneSegment("0001001"))
	fmt.Println(checkOneSegment("000000"))
	fmt.Println(checkOneSegment("0"))
	fmt.Println(checkOneSegment("1"))
}*/

func checkOneSegment(s string) bool {
	if len(s) == 0 {
		return true
	}
	slow, fast, count := 0, 0, 0
	bytes := []byte(s)
	for fast <= len(s)-1 {
		if bytes[fast] != '1' {
			if fast != slow {
				count++
			}
			fast++
			slow = fast
		} else {
			fast++
		}
	}
	// 1111
	if fast != slow {
		count++
	}
	if count <= 1 {
		return true
	}
	return false
}
