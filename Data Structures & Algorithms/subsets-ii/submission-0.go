func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}

    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        tmp := make([]int, len(path))
        copy(tmp, path)
        res = append(res, tmp)

        for i := start; i < len(nums); i++ {
            if i > start && nums[i] == nums[i-1] {
                continue
            }
            path = append(path, nums[i])
            dfs(i+1, path)
            path = path[:len(path)-1]
        }
    }

    dfs(0, []int{})
    return res
}