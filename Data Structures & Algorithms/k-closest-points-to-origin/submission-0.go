type MinHeap [][]int

func (h *MinHeap) Push(a interface{}) {
	*h = append(*h, a.([]int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return x
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	l1 := h[i][0] * h[i][0] + h[i][1]*h[i][1]
	l2 := h[j][0] * h[j][0] + h[j][1]*h[j][1]
	return l1 < l2
}

func (h MinHeap) Swap(i,j int) {
	h[i], h[j] = h[j], h[i]
}

func kClosest(points [][]int, k int) [][]int {
	h := &MinHeap{}
	heap.Init(h)
	for _, el := range points {
		heap.Push(h, el)
	}
	var res [][]int
	for i := 0 ; i < k ; i++ {
		tmp := heap.Pop(h).([]int)
		res = append(res, tmp)
	}
	return res
}
