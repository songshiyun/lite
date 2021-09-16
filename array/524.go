package array

import "strings"

/*
	s = "abpcplea",
	dictionary = ["ale","apple","monkey","plea"]
	  *
	abpcplea
	ale      1
	apple    2
	monkey   0
	plea     1
*/

// 1.可以每个排序后比较
// 2.可以单个依次比较
func findLongestWord(s string, dictionary []string) string {
	var result string
	index := make([]int, len(dictionary))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(dictionary); j++ {
			if index[j] < len(dictionary[j]) && s[i] == dictionary[j][index[j]] {
				index[j]++
				if index[j] == len(dictionary[j]) && (len(dictionary[j]) > len(result) ||
					len(result) == len(dictionary[j]) && strings.Compare(result, dictionary[j]) == +1) {
					result = dictionary[j]
				}
			}
		}
	}
	return result
}
