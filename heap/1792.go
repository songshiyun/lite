package heap

import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	var hpx ItemHeap
	heap.Init(&hpx)
	for i := range classes {
		rat := ratio(classes, i, 1) - ratio(classes, i, 0)
		heap.Push(&hpx, Item{
			Ratio: rat,
			Index: i,
		})
	}
	for ; extraStudents > 0; extraStudents-- {
		it := heap.Pop(&hpx).(Item)
		classes[it.Index][0]++
		classes[it.Index][1]++
		rat1 := ratio(classes, it.Index, 1) - ratio(classes, it.Index, 0)
		heap.Push(&hpx, Item{
			Ratio: rat1,
			Index: it.Index,
		})
	}
	var total float64
	for i := range classes {
		total += ratio(classes, i, 0)
	}
	return total / float64(len(classes))
}
func ratio(class [][]int, index, delta int) float64 {
	return float64(class[index][0]+delta) / float64(class[index][1]+delta)
}

type Item struct {
	Ratio float64
	Index int
}
type ItemHeap []Item

func (ih ItemHeap) Len() int {
	return len(ih)
}
func (ih ItemHeap) Less(k, j int) bool {
	return ih[k].Ratio > ih[j].Ratio
}
func (ih ItemHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}
func (ih *ItemHeap) Push(data interface{}) {
	*ih = append(*ih, data.(Item))
}
func (ih *ItemHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	data := old[n-1]
	*ih = old[0 : n-1]
	return data
}
