package heap

type MinHeap []int

func NewIntHeap() MinHeap {
	return NewIntHeapWithLength(0)
}

func NewIntHeapWithLength(length int) MinHeap {
	return make([]int, length)
}

func (heap *MinHeap) IsEmpty() bool {
	return heap.Len() == 0
}

func (heap *MinHeap) Len() int {
	return len(*heap)
}

func (heap *MinHeap) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *MinHeap) Poll() int {
	top := (*heap)[0]
	heap.Swap(0, heap.Len()-1)
	*heap = (*heap)[:heap.Len()-1]
	heap.shiftDown(0)
	return top
}

func (heap *MinHeap) Peek() int {
	n := len(*heap) - 1
	old := *heap
	return old[n]
}

func (heap *MinHeap) Push(x int) {
	*heap = append(*heap, x)
	heap.shiftUp(heap.Len() - 1)
}

func (heap *MinHeap) shiftUp(child int) {
	for {
		parent := (child - 1) / 2
		if parent == child || (*heap)[parent] < (*heap)[child] {
			break
		}
		heap.Swap(parent, child)
		child = parent
	}
}

func (heap *MinHeap) shiftDown(parent int) {
	for {
		leftChild := 2*parent + 1
		if leftChild >= heap.Len() || leftChild < 0 {
			break
		}
		j := leftChild
		if rightChild := j + 1; rightChild < heap.Len() && (*heap)[rightChild] < (*heap)[leftChild] { //找左右子节点较小的
			j = rightChild
		}
		if (*heap)[j] >= (*heap)[parent] {
			break
		}
		heap.Swap(parent, j)
		parent = j
	}
}
