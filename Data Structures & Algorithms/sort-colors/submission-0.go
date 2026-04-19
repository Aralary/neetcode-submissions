func sortColors(nums []int) {
    var stat []int = make([]int, 3, 3)
	for _, el := range nums {
		stat[el]++
	}
	tmp := 0
	for id,_ := range nums {
		if stat[tmp]== 0 {
			tmp++
		}
		if stat[tmp] > 0 {
			nums[id] = tmp
			stat[tmp]--
		}
	}
}
