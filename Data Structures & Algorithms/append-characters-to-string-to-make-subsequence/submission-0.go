func appendCharacters(s string, t string) int {
	return len(t) - find(s, t, 0, 0)
}

func find(s string , t string , sStart int, tStart int) int {
	if tStart >= len(t) || sStart >= len(s) {
		return tStart
	}
	for i := sStart ; i < len(s) ; i++ {
		if s[i] == t[tStart] {
			tStart++
			if tStart >= len(t) {
				break
			}
		}
	}
	return tStart
}