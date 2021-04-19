package design

import "container/list"


//lru
//双向链表+map
//map中保存的是双向链表中的element
type LRUCache struct {
	Cap  int
	Map  map[int]*list.Element
	List *list.List
}

func ConstructorLru(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Map:  make(map[int]*list.Element, capacity),
		List: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	ele, ok := this.Map[key]
	if ok {
		this.List.MoveToFront(ele)
		return ele.Value.(Pair).V
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	ele, ok := this.Map[key]
	if ok {
		ele.Value = Pair{K: key, V: value}
		this.List.MoveToFront(ele)
	} else {
		el := this.List.PushFront(Pair{K: key, V: value})
		this.Map[key] = el
	}
	if this.List.Len() > this.Cap {
		el := this.List.Back()
		this.List.Remove(el)
		delete(this.Map, el.Value.(Pair).K)
	}
}

type Pair struct {
	K, V int
}
