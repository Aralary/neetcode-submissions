func twoSum(nums []int, target int) []int {
	st := make(map[int][]int)
	for id, el := range nums {
		st[el] = append(st[el], id)
	}
	for k, val := range st {
		need := target - k 
		if ids, ok := st[need]; ok {
			if need == k {
				if len(val) > 1 {
					return val[:2]
				}
			} else {
				return []int{min(val[0], ids[0]),max(val[0], ids[0])}
			}
		}
	}
	return []int{}
}
