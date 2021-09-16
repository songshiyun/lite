package array

import "testing"

func Test524(t *testing.T) {
	t.Log(findLongestWord("abpcplea",[]string{"ale","apple","monkey","plea"}))
	t.Log(findLongestWord("abpcplea",[]string{"a","b","c","d"}))
}