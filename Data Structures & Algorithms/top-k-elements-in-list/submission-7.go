type MaxHeap struct {
    nums []int
    freq map[int]int
}

func (h *MaxHeap) Push(a interface{}) { h.nums = append(h.nums, a.(int)) }
func (h *MaxHeap) Pop() interface{} {
    x := h.nums[len(h.nums)-1]
    h.nums = h.nums[:len(h.nums)-1]
    return x
}
func (h MaxHeap) Len() int           { return len(h.nums) }
func (h MaxHeap) Less(i, j int) bool { return h.freq[h.nums[i]] > h.freq[h.nums[j]] }
func (h MaxHeap) Swap(i, j int)      { h.nums[i], h.nums[j] = h.nums[j], h.nums[i] }

func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, el := range nums {
        freq[el]++
    }

    h := &MaxHeap{freq: freq}
    heap.Init(h)
    for num := range freq {
        heap.Push(h, num)
    }

    res := make([]int, 0, k)
    for i := 0; i < k; i++ {
        res = append(res, heap.Pop(h).(int))
    }
    return res
}