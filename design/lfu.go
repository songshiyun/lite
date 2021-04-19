package design

import "container/list"

type LFUCache struct {
	Nodes map[int]*list.Element
	List  map[int]*list.List
	Cap   int
	Min   int
}

type node struct {
	k         int
	v         int
	frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Nodes: make(map[int]*list.Element),
		List:  make(map[int]*list.List),
		Cap:   capacity,
		Min:   0,
	}
}

//最近最少访问： 只需要记录min次数的节点，不需要维护其他的次数
//通过两个map分别记录 key -> (key,value,frequency)和 frequency ->(key,value,frequency)
//判断元素是否存在
//存在:
//从原来的list中删除对应的val
//更新frequency的值
//在新的frequency的list中插入currentNode
//将插入后的心的element的节点赋值给node[key]
func (this *LFUCache) Get(key int) int {
	if ele, ok := this.Nodes[key]; ok {
		currentNode := ele.Value.(*node)
		this.List[currentNode.frequency].Remove(ele)
		currentNode.frequency++
		if _, ok := this.List[currentNode.frequency]; !ok {
			this.List[currentNode.frequency] = list.New()
		}
		newNode := this.List[currentNode.frequency].PushFront(currentNode)
		this.Nodes[key] = newNode
		//当前操作的是min
		if this.Min == currentNode.frequency-1 && this.List[currentNode.frequency-1].Len() == 0 {
			this.Min++
		}
		return currentNode.v
	}
	return -1
}

/*func findMin(data map[int]*list.List) int {
	if len(data) == 0 {
		return 0
	}
	var min int
	for k, _ := range data {
		if k < min {
			min = k
		}
	}
	return min
}*/

//case 1:原来存在
// 找到原来的节点，更新capacity, ->get 操作
// 原来不存在： 删除访问次数最最小的，然后插入新的节点
//case 2:原来不存在，

func (this *LFUCache) Put(key int, value int) {
	if this.Cap == 0 {
		return
	}
	if ele, ok := this.Nodes[key]; ok {
		currentNode := ele.Value.(*node)
		currentNode.v = value
		this.Get(key)
		return
	}
	//原来不存在，且节点已经满了
	if len(this.Nodes) >= this.Cap {
		oldList := this.List[this.Min]
		oldNode := oldList.Back()
		delete(this.Nodes, oldNode.Value.(*node).k)
		oldList.Remove(oldNode)
	}
	this.Min = 1
	newNode := &node{
		k:         key,
		v:         value,
		frequency: 1,
	}
	if _, ok := this.List[newNode.frequency]; !ok {
		this.List[newNode.frequency] = list.New()
	}
	newEle := this.List[newNode.frequency].PushFront(newNode)
	this.Nodes[key] = newEle
	return
}
