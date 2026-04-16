func isSubsequence(s string, t string) bool {
	return find(s, t, 0, 0)
}

func find(s string, t string, sStart, tStart int) bool {
	if len(t) - tStart  < len(s) - sStart { return false }
	if sStart >= len(s) - 1{ return true }
	if tStart >= len(t) -1 { return false }

	for i := tStart ; i < len(t) ; i++ {
		tStart = i + 1
		if t[i] == s[sStart] {
			sStart++
			break
		}
	}
	return find(s,t, sStart, tStart)
}
