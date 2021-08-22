package dp

/**
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
 */

// 完全背包问题
// 物品是wordDict, 背包s,能否凑满整个字符串，背包中的能否组成整个字符串
// 能否组成不关心顺序，组合问题（先遍历背包和先遍历物品都可以）

//如果求组合数就是外层for循环遍历物品，内层for遍历背包。
//如果求排列数就是外层for遍历背包，内层for循环遍历物品。
func wordBreak(s string, wordDict []string) bool {
	wordset := make(map[string]struct{})
	for _,item := range wordDict {
		wordset[item] = struct{}{}
	}
	dp := make([]bool,len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j <i; j++ {
			temp := s[j:i]
			_,ok := wordset[temp]
			if ok && dp[j]{
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func wordBreak1(s string, wordDict []string) bool  {
	wordset := make(map[string]struct{})
	for _,item := range wordDict {
		wordset[item] = struct{}{}
	}
	memory := make([]int,len(s)+1)
	for i,_ := range memory {
		memory[i] = -1
	}
	var backTrace func(index int,memory []int) bool
	backTrace = func(index int, memory []int) bool {
		if index >= len(s) {
			return true
		}
		for i := index; i <len(s); i++ {
			word := s[index: i-index +1]
			_,ok := wordset[word]
			if ok && backTrace(i+1,memory){
				memory[index] = 1
				return true
			}
		}
		memory[index] = 0
		return false
	}
	return backTrace(0,memory)
}
