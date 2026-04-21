func numIslands(grid [][]byte) int {
    rows, cols := len(grid), len(grid[0])
    count := 0

    var bfs func(r, c int)
    bfs = func(r, c int) {
        queue := [][2]int{{r, c}}
        grid[r][c] = '0'

        dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
        for len(queue) > 0 {
            cur := queue[0]
            queue = queue[1:]

            for _, d := range dirs {
                nr, nc := cur[0]+d[0], cur[1]+d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '1' {
                    grid[nr][nc] = '0'
                    queue = append(queue, [2]int{nr, nc})
                }
            }
        }
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '1' {
                bfs(r, c)
                count++
            }
        }
    }
    return count
}