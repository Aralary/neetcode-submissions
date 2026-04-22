func orangesRotting(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    queue := [][2]int{}
    for i := 0 ; i < rows ; i++{
        for j := 0 ; j < cols; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, [2]int{i,j})
            }
        }
    }
    dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    for len(queue) > 0 {
        tmpPos := queue[0]
        queue = queue[1:]
        for _, dir := range dirs {
            ni, nj := tmpPos[0] + dir[0], tmpPos[1] + dir[1]
            if ni >=0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == 1 {
                grid[ni][nj] = grid[tmpPos[0]][tmpPos[1]] + 1
                queue = append(queue, [2]int{ni,nj})
            }
        }
    }
    maxVal := 2
    for i := 0 ; i < rows ; i++{
        for j := 0 ; j < cols; j++ {
            if grid[i][j] == 1 {
                return -1
            }
            maxVal = max(maxVal, grid[i][j])
        }
    }
    return maxVal - 2
}
