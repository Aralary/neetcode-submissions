func hasDuplicate(nums []int) bool {
    stat := make(map[int]int)
	for _, el := range nums {
		stat[el]++
		if stat[el] > 1 { return true }
	}
	return false
}
