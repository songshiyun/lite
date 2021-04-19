package heap

type MaxHeap []int

func NewMaxHeap() MaxHeap {
	return make([]int, 0)
}

func (heap *MaxHeap) IsEmpty() bool {
	return heap.Len() == 0
}

func (heap *MaxHeap) Len() int {
	return len(*heap)
}

func (heap *MaxHeap) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *MaxHeap) Push(x int) {
	*heap = append(*heap, x)
	heap.shiftUp(heap.Len() - 1)
}

func (heap *MaxHeap) Pop() int {
	top := (*heap)[0]
	heap.Swap(0, heap.Len()-1)
	*heap = (*heap)[:heap.Len()-1]
	heap.shiftDown(0)
	return top
}

func (heap *MaxHeap) Peek() int {
	return (*heap)[0]
}

func (heap *MaxHeap) shiftUp(child int) {
	for {
		parent := (child - 1) / 2
		if child == parent || (*heap)[child] < (*heap)[parent] {
			break
		}
		heap.Swap(parent, child)
		child = parent
	}
}

func (heap *MaxHeap) shiftDown(parent int) {
	for {
		leftChild := 2*parent + 1
		if leftChild >= heap.Len() || leftChild < 0 {
			break
		}
		j := leftChild
		if rightChild := j + 1; rightChild < heap.Len() && (*heap)[rightChild] > (*heap)[leftChild] {
			j = rightChild
		}
		if (*heap)[j] <= (*heap)[parent] {
			break
		}
		heap.Swap(parent, j)
		parent = j
	}
}
