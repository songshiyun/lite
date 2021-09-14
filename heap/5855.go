package heap

import (
	"bytes"
	"sort"
)

func kthLargestNumber(nums []string, k int) string {
	//sort.Sort(stringHeap(nums))
	sort.Sort(sortString{nums})
	return nums[k-1]
}

type stringHeap []string

func (p stringHeap) Len() int { return len(p) }
func (p stringHeap) Less(i, j int) bool {
	if len(p[i]) < len(p[j]) {
		return false
	}
	if len(p[i]) > len(p[j]) {
		return true
	}
	result := bytes.Compare([]byte(p[i]), []byte(p[j]))
	if result > 0 {
		return true
	}
	return false
}
func (p stringHeap) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// heap.Pop(p)
// heap.Push(p,item)
/*func (s *stringHeap)Push(i interface{}) {
	*s = append(*s, i.(string))
}

func (s *stringHeap)Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0:n-1]
	return x
}*/

type sortString struct {
	sort.StringSlice
}

func (s sortString)Less(i,j int) bool {
	if len(s.StringSlice[i]) < len(s.StringSlice[j]) {
		return false
	}
	if len(s.StringSlice[i]) > len(s.StringSlice[j]) {
		return true
	}
	result := bytes.Compare([]byte(s.StringSlice[i]), []byte(s.StringSlice[j]))
	if result > 0 {
		return true
	}
	return false
}