package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B := []string{"lo", "eo"}
	fmt.Println(wordSubsets(A, B))
}

func wordSubsets(A []string, B []string) []string {
	charSet := make([]string, len(A))
	for index, item := range A {
		charSet[index] = sortSting(item)
	}
	var result []string
	for i, item := range charSet {
		var isSub bool
		for index, value := range B { //B中所有的都是这个元素的子集
			if index != 0 && !isSub {
				break
			}
			if strings.Index(item, sortSting(value)) != -1 {
				isSub = true
			} else {
				isSub = false
			}
		}
		if isSub {
			result = append(result, A[i])
		}
	}
	return result
}

func wordSubsets1(A []string, B []string) []string {
	return nil
}

func mapSet(A []string) map[string]map[byte]int {
	charset := make(map[string]map[byte]int, len(A))
	for _, item := range A {
		mp, ok := charset[item]
		if !ok {
			mp = make(map[byte]int)
		}
		for _, b := range []byte(item) {
			mp[b]++
		}
	}
	return charset
}

func sortSting(string2 string) string {
	buf := []byte(string2)
	sort.Slice(buf, func(i, j int) bool {
		return buf[i] < buf[j]
	})
	return string(buf)
}

func commonData(a, b []string) []string {
	var result []string
	for _, item := range a {
		for _, data := range b {
			if item == data {
				result = append(result, item)
			}
		}
	}
	return result
}
