package stack

func checkValidString(s string) bool {
	// ( , * 入栈，
	// ) -> 是否有 （， 是否有*
	// 对于最终还是剩下的 只有 （ 和 *
	var temp []byte
	var start int
	for _,item := range []byte(s) {
		if item == '(' || item == '*' {
			temp  = append(temp,item)
			if item == '*' {
				start++
			}
		}else {
			// )
			if len(temp) == 0 {
				return false
			}
			var hasLeft bool
			ln := len(temp)
			for i := ln-1; i >=0; i-- {
				if temp[i] == '('  {
					hasLeft = true
					temp = append(temp[0:i],temp[i+1:]...)
					break
				}
			}
			if !hasLeft {
				temp = temp[0:len(temp)-1]
				start--
			}
		}
	}
	if len(temp) == 0 {
		return true
	}
	var t []byte
	for _,item := range temp {
		if item == '(' {
			t = append(t,item)
		}else{
			if len(t) != 0 {
				t = t[0:len(t)-1]
			}
		}
	}
	if len(t) != 0 {
		return false
	}
	return true
}
