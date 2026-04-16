func getConcatenation(nums []int) []int {
	n := len(nums)
    var res []int = make([]int, 2*n)
	for i := 0 ; i < n; i++ {
		res[i], res[n+i] = nums[i], nums[i]
	}
	return res
}
