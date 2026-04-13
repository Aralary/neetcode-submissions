func permute(nums []int) [][]int {
    res := [][]int{}
    used := make([]bool, len(nums))

    var dfs func(path []int)
    dfs = func(path []int) {
        if len(path) == len(nums) {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }

        for i := 0; i < len(nums); i++ {
            if used[i] {
                continue
            }

            used[i] = true 
            path = append(path, nums[i])

            dfs(path)

            path = path[:len(path)-1]
            used[i] = false
        }
    }

    dfs([]int{})
    return res
}