func generateParenthesis(n int) []string {
	var res []string
	var dfs func(int,  int , string)
	dfs = func(open int, close int , path string) {
		if open == n && close == n {
			res = append(res, path)
			return
		}
		if open < n {
			dfs(open + 1, close, path + "(")
		}
		if close < open {
			dfs(open, close + 1, path + ")")
		}
	}
	dfs(0,0,"")
	return res
}
