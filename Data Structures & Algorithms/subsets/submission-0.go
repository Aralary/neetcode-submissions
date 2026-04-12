func subsets(nums []int) [][]int {
    res := [][]int{{}}

    for _, num := range nums {
        n := len(res)
        for i := 0; i < n; i++ {
            tmp := make([]int, len(res[i]), len(res[i])+1)
            copy(tmp, res[i])
            res = append(res, append(tmp, num))
        }
    }

    return res
}
