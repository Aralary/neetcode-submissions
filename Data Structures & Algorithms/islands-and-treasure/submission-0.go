func islandsAndTreasure(grid [][]int) {
    rows, cols := len(grid), len(grid[0])
    inf := 2147483647
    queue := [][2]int{}

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 0 {
                queue = append(queue, [2]int{i, j})
            }
        }
    }

    dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        for _, d := range dirs {
            ni, nj := cur[0]+d[0], cur[1]+d[1]
            if ni >= 0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == inf {
                grid[ni][nj] = grid[cur[0]][cur[1]] + 1
                queue = append(queue, [2]int{ni, nj})
            }
        }
    }
}