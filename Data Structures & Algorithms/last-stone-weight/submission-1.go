type MaxHeap []int

func (h *MaxHeap) Push(a interface{}) {
	*h = append(*h, a.(int))
}

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return x
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i,j int) {
	h[i], h[j] = h[j], h[i]
}

func AbsInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	heap.Init(h)
	for _, el := range stones{
		heap.Push(h, el)
	}
	for h.Len() != 1 {
		x, y := heap.Pop(h).(int), heap.Pop(h).(int)
		if x != y {
			heap.Push(h, AbsInt(x - y))
		}
		if h.Len() == 0 {
			heap.Push(h, 0)
		}
	}
	return (*h)[0]
}
