package dfs

var (
	digitsMap = map[string]string{
		"0": "",
		"1": "",
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
)

func letterCombinations(digits string) (res []string) {
	if digits == "" {
		return
	}
	var dfs func(index int, dst string)
	dfs = func(index int, dst string) {
		if index == len(digits) {
			res = append(res, dst)
			return
		}
		//dfs,1. 先遍历找到外层字母
		//2.再遍历内层数组
		letter := digitsMap[string(digits[index])]
		for i := 0; i < len(letter); i++ {
			dfs(index+1, dst+string(letter[i]))
		}
		return
	}
	dfs(0, "")
	return
}

var letterMap = map[string][]string{
	"0": {},
	"1": {},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

//backtrace
func letterCombinations3(digits string) (res []string) {
	if digits == "" {
		return
	}
	var dfs func(src, tmp string)
	dfs = func(src, tmp string) {
		if src == "" {
			res = append(res, tmp)
			return
		}
		k := src[0:1]
		src = src[1:]
		for i := 0; i < len(letterMap[k]); i++ {
			tmp += letterMap[k][i] //选择当前这个
			dfs(src, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(digits, "")
	return
}

func letterCombinations2(digits string) (res []string) {
	if len(digits) == 0 {
		return
	}
	var dfs func(index int, dst string)
	dfs = func(index int, dst string) {
		if index == len(digits) {
			res = append(res, dst)
			return
		}
		current := letterMap[string(digits[index])]
		for _, item := range current {
			dfs(index+1, dst+string(item))
		}
		return
	}
	dfs(0, "")
	return
}
