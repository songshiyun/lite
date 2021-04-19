package design

import (
	"container/heap"
)

func Constructor1(cap int) *LFUCache1 {
	return &LFUCache1{
		Map:      make(map[int]*Item),
		PQ:       make([]*Item, 0, 0),
		Capacity: cap,
	}
}

func (lfu *LFUCache1) Get(key int) int {
	lfu.Counter++
	if item, ok := lfu.Map[key]; ok {
		lfu.PQ.update(item, item.V, lfu.Counter, item.Frequency+1)
		return item.V
	}
	return -1
}

func (lfu *LFUCache1) Put(key, val int) {
	lfu.Counter++
	if item, ok := lfu.Map[key]; ok {
		lfu.PQ.update(item, val, lfu.Counter, item.Frequency+1)
		return
	}
	//不存在，切已经大于cap
	if len(lfu.Map) >= lfu.Capacity {
		item := heap.Pop(&lfu.PQ).(*Item)
		delete(lfu.Map, item.K)
	}
	newItem := &Item{
		K:         key,
		V:         val,
		Frequency: 1,
		Count:     lfu.Counter,
	}
	heap.Push(&lfu.PQ, newItem)
	lfu.Map[key] = newItem
	return
}

type LFUCache1 struct {
	Map      map[int]*Item
	PQ       PriorityQueue
	Capacity int
	Counter  int //like logical clock,相同capacity,删除最久访问的
}

type Item struct {
	K         int
	V         int
	Frequency int
	Count     int
	Index     int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) update(item *Item, value int, count int, frequency int) {
	item.V = value
	item.Count = count
	heap.Fix(pq, item.Index)
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Frequency == pq[j].Frequency {
		return pq[i].Count < pq[j].Count
	}
	return pq[i].Frequency < pq[j].Frequency
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Pop() interface{} {
	oldpq := *pq
	n := len(oldpq)
	old := oldpq[n-1]
	oldpq[n-1] = nil
	old.Index = -1
	*pq = oldpq[0 : n-1]
	return old
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}
