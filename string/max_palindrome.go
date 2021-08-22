package string

// 最长回文字符串
func longestPalindrome(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		res = maxPalindromeString(str, i, i, res)
		res = maxPalindromeString(str, i, i+1, res)
	}

	return res
}

func maxPalindromeString(str string, i, j int, res string) string {
	sub := ""
	for i < j && j < len(str) && str[i] == str[j] {
		sub = str[i : j+1]
		i++
		j--
	}
	if len(sub) > len(res) {
		return sub
	}
	return res
}
