
type MinHeap []int
func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    a := *h; n := len(a)
    x := a[n-1]; *h = a[:n-1]; return x
}

type KthLargest struct {
    k int
    h *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
    h := &MinHeap{}
    heap.Init(h)
    obj := KthLargest{k: k, h: h}
    for _, v := range nums { obj.Add(v) }
    return obj
}

func (this *KthLargest) Add(val int) int {
    heap.Push(this.h, val)
    if this.h.Len() > this.k {
        heap.Pop(this.h)
    }
    return (*this.h)[0]
}