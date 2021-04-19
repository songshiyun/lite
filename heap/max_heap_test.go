package heap

import (
	"fmt"
	"log"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	heap := NewMaxHeap()
	heap.Push(2)
	heap.Push(20)
	heap.Push(12)
	heap.Push(22)
	heap.Push(4)
	heap.Push(1)
	heap.Push(3)
	fmt.Println(fmt.Sprintf("%+v",heap))
	for !heap.IsEmpty() {
		log.Println()
		log.Println(heap.Pop())
		fmt.Println(fmt.Sprintf("%+v",heap))
	}
}
