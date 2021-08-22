package dp

import "testing"

func TestWordBreak(t *testing.T) {
	src := []string{"leet","code"}
	s := "leetcode"
	t.Log(wordBreak(s,src))
	t.Log(wordBreak1(s,src))
}