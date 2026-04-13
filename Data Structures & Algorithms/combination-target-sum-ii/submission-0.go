func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    res := [][]int{}

    var dfs func(start, sum int, path []int)
    dfs = func(start, sum int, path []int) {
        if sum == target {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }

        for i := start; i < len(candidates); i++ {
            if i > start && candidates[i] == candidates[i-1] {
                continue
            }
            if sum + candidates[i] > target {
                continue
            }
            path = append(path, candidates[i])
            dfs(i + 1, sum+candidates[i], path)
            path = path[:len(path)-1]
        }
    }

    dfs(0, 0, []int{})
    return res
}
