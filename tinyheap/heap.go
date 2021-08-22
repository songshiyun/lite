package tinyheap

type Item interface {
	Less(item Item) bool
}

type Heap interface {
	Push(item Item)
	Pop() Item
	Peek() Item
	Len() int
}

type heap struct {
	length int
	data   []Item
}

func (h *heap) Push(item Item) {
	if h.data == nil {
		h.data = make([]Item, 0)
	}
	h.length++
	h.data = append(h.data, item)
	h.up(h.length - 1)
}

func (h *heap) Pop() Item {
	if h.length == 0 {
		return nil
	}
	top := h.data[0]
	h.length--
	if h.length > 0 {
		h.data[0] = h.data[h.length]
		h.down(0)
	}
	h.data = h.data[:h.length-1]
	return top
}

func (h *heap) Peek() Item {
	if h.length == 0 {
		return nil
	}
	return h.data[0]
}
func (h *heap) Len() int {
	return h.length
}

func (h *heap) down(pos int) {
	data := h.data
	halfLen := h.length >> 1
	item := data[pos]
	for pos < halfLen {
		left := (pos << 1) + 1
		right := left + 1
		best := data[left]
		if right < h.length && data[right].Less(best) {
			left = right
			best = data[right]
		}
		if !best.Less(item) {
			break
		}
		data[pos] = best
		pos = left
	}
	data[pos] = item
}

func (h *heap) up(pos int) {
	data := h.data
	item := data[pos]
	for pos > 0 {
		parent := (pos - 1) >> 1
		current := data[parent]
		if current.Less(item) {
			break
		}
		data[pos] = current
		pos = parent
	}
	data[pos] = item
}
