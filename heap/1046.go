package heap

import "container/heap"

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	h := &IntSlice{}
	heap.Init(h)
	for _, item := range stones {
		h.Push(item)
	}
	for h.Len() >= 2 {
		y := h.Pop()
		x := h.Pop()
		if y.(int) == x.(int) {
			continue
		}
		h.Push(y.(int) - x.(int))
	}
	if h.Len() == 0 {
		return 0
	}
	return h.Pop().(int)
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (i *IntSlice) Pop() interface{} {
	old := *i
	n := len(old)
	x := old[n-1]
	*i = old[0 : n-1]
	return x
}

func (i *IntSlice) Push(x interface{}) {
	*i = append(*i, x.(int))
}
