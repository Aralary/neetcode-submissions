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

func findKthLargest(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)
	for _, el := range nums {
		heap.Push(h, el)
	}
	res := heap.Pop(h).(int)
	for i := 1 ; i < k ; i++ {
		res = heap.Pop(h).(int)
	}
	return res
}
