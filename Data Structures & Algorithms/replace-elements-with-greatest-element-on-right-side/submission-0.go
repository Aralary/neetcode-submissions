func replaceElements(arr []int) []int {
	n := len(arr)
	var prefix []int = make([]int, n)
	prefix[n-1] = arr[n-1]
	tmpMax := arr[n-1]
	for i := n-1 ; i > -1 ; i-- {
		prefix[i] = tmpMax
		if arr[i] > tmpMax {
			tmpMax = arr[i]
		}
	}
	prefix[n-1] = -1
	return prefix
}
