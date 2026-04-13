func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}

    var dfs func(start, sum int, path []int)
    dfs = func(start, sum int, path []int) {
        if sum == target {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }

        for i := start; i < len(nums); i++ {
            if sum + nums[i] > target {
                continue // pruning: этот элемент слишком большой
            }
            path = append(path, nums[i])
            dfs(i, sum+nums[i], path)
            path = path[:len(path)-1]
        }
    }

    dfs(0, 0, []int{})
    return res
}