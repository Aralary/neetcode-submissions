func subsets(nums []int) [][]int {
    res := [][]int{}
    subset := []int{}

    var dfs func(i int)
    dfs = func(i int) {
        if i == len(nums) {
            // делаем копию — иначе все ссылки будут на один слайс
            tmp := make([]int, len(subset))
            copy(tmp, subset)
            res = append(res, tmp)
            return
        }
        // Выбор 1: включаем nums[i]
        subset = append(subset, nums[i])
        dfs(i + 1)

        // Выбор 2: не включаем nums[i] (backtrack)
        subset = subset[:len(subset)-1]
        dfs(i + 1)
    }

    dfs(0)
    return res
}