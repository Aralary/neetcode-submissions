func customSortString(order string, s string) string {
	stat := make(map[rune]int)
	for _, el := range s {
		stat[el]++
	}
	var res string
	for _, el := range order {
		if n, ok := stat[el]; ok {
			res = res + strings.Repeat(string(el), n)
			delete(stat, el)
		}
	} 
	for k, v := range stat {
		res = res + strings.Repeat(string(k), v)
	}
	return res
}
