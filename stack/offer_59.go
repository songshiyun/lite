package stack

func Constructor() MaxQueue {
	return MaxQueue{
		que: make([]int, 0),
		deq: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.deq) == 0 {
		return -1
	}
	return this.deq[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.que = append(this.que, value)
	for len(this.deq) != 0 && value > this.deq[len(this.deq)-1] {
		this.deq = this.deq[:len(this.deq)-1]
	}
	this.deq = append(this.deq, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.que) == 0 {
		return -1
	}
	if this.que[0] == this.deq[0] {
		this.deq = this.deq[1:]
	}
	x := this.que[0]
	this.que = this.que[1:]
	return x
}

//单调栈
type MaxQueue struct {
	que []int
	deq []int //单调递减的队列，队头存储最大的元素
}
