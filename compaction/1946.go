package compaction

import (
	"bytes"
	"strconv"
)

// 突变越靠前越好
func maximumNumber(num string, change []int) string {
	l := len(num)
	i := 0
	var result bytes.Buffer
	var last int = -1
	var notfirst bool = true
	for i < l {
		// 没有突变并且不满足突变条件，
		// 前一个已经突变但是不满足突变条件
		k := i
		i++
		index := int(num[k] - '0')
		c := change[index]
		if index >= c || index < c && (last+1 != k && !notfirst) {
			if index == c {
				last = k
			}
			result.WriteByte(num[k])
			//changed = false
			continue
		}
		// 需要突变
		notfirst = false
		last = k
		result.WriteString(strconv.Itoa(c))
	}
	return result.String()
}

func maximumNumber1(num string, change []int) string {
	l := len(num)
	i := 0
	notchanged := true
	var result bytes.Buffer
	for ; i < l; i++ {
		index := int(num[i] - '0')
		c := change[index]
		if notchanged && index >= c {
			result.WriteByte(num[i])
			continue
		} else if c >= index {
			notchanged = false
			result.WriteString(strconv.Itoa(c))
			continue
		}
		break
	}
	result.WriteString(num[i:l])
	return result.String()
}
