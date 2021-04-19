package skiplist

import (
	"math"
	"math/rand"
	"time"
)

const (
	DefaultDeep             = 16
	DefaultPriority float64 = 1 / math.E
)

type elementNode struct {
	Next []*Element
}

type Element struct {
	elementNode
	Key   int64
	Value interface{}
}

type Skiplist struct {
	elementNode
	maxLevel       int
	Length         int
	randSource     rand.Source
	probability    float64
	probTable      []float64
	prevNodesCache []*elementNode
	existMap       map[int64]int
}

func Constructor() Skiplist {
	return NewWithMexLevel(DefaultDeep)
}

func NewWithMexLevel(level int) Skiplist {
	if level < 1 || level > DefaultDeep {
		level = DefaultDeep
	}
	return Skiplist{
		elementNode: elementNode{make([]*Element, level)},
		maxLevel:    level,
		randSource:  rand.NewSource(time.Now().UnixNano()),
		probability: DefaultPriority,
		probTable:   probabilityTable(DefaultPriority, level),
		existMap:    make(map[int64]int),
	}
}

func probabilityTable(probability float64, MaxLevel int) (table []float64) {
	for i := 1; i <= MaxLevel; i++ {
		prob := math.Pow(probability, float64(i-1))
		table = append(table, prob)
	}
	return table
}

func (skl *Skiplist) Add(num int) {
	var element *Element
	prev := skl.getPrevElementNodes(int64(num))
	if element = prev[0].Next[0]; element != nil && element.Key <= int64(num) {
		//element.Value =
		//存在
	}
	element = &Element{
		elementNode: elementNode{
			Next: make([]*Element, skl.randLevel()),
		},
	}
	for i := range element.Next {
		element.Next[i] = prev[i].Next[i]
		prev[i].Next[i] = element
	}
	skl.Length++
	skl.fastSet(int64(num))
	return
}

func (skl *Skiplist) randLevel() (level int) {
	r := float64(skl.randSource.Int63()) / (1 << 63)
	level = 1
	for level < skl.maxLevel && r < skl.probTable[level] {
		level++
	}
	return
}

func (skl *Skiplist) Search(target int) bool {
	var prev *elementNode = &skl.elementNode
	var next *Element
	for i := skl.maxLevel - 1; i >= 0; i-- {
		next = prev.Next[i]
		for next != nil && int64(target) > next.Key {
			prev = &next.elementNode
			next = next.Next[i]
		}
	}
	if next != nil && next.Key <= int64(target) {
		return true
	}
	return false
}

func (skl *Skiplist) fastSet(key int64) {
	count, ok := skl.existMap[key]
	if ok {
		skl.existMap[key] = count + 1
	} else {
		skl.existMap[key] = 1
	}
}

func (skl *Skiplist) fastExist(key int64) bool {
	_, ok := skl.existMap[key]
	return ok
}
func (skl *Skiplist) slowRemove(key int64) {
	count, _ := skl.existMap[key]
	if count-1 == 0 {
		delete(skl.existMap, key)
	} else {
		skl.existMap[key] = count - 1
	}
}

func (skl *Skiplist) Erase(num int) bool {
	if !skl.fastExist(int64(num)) {
		return false
	}
	_ = skl.Remove(int64(num))
	skl.slowRemove(int64(num))
	return true
}

func (skl *Skiplist) Remove(key int64) *Element {
	prevs := skl.getPrevElementNodes(key)
	if element := prevs[0].Next[0]; element != nil && element.Key <= key {
		for k, v := range element.Next {
			prevs[k].Next[k] = v
		}
		skl.Length--
		return element
	}
	return nil
}

func (skl *Skiplist) getPrevElementNodes(key int64) []*elementNode {
	var prev *elementNode = &skl.elementNode
	var next *Element
	prevs := skl.prevNodesCache
	for i := skl.maxLevel - 1; i >= 0; i-- {
		next = prev.Next[i]
		for next != nil && key > next.Key {
			prev = &next.elementNode
			next = next.Next[i]
		}
		prevs[i] = prev
	}
	return prevs
}
