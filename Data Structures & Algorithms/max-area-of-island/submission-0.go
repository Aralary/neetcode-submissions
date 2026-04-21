func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    maxSize := 0

    var bfs func(r, c int) int
    bfs = func(r, c int) int {
        size := 1
        queue := [][2]int{{r, c}}
        grid[r][c] = 0 // int 0, не '0'

        dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
        for len(queue) > 0 {
            cur := queue[0]
            queue = queue[1:]

            for _, d := range dirs {
                nr, nc := cur[0]+d[0], cur[1]+d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
                    grid[nr][nc] = 0
                    queue = append(queue, [2]int{nr, nc})
                    size++
                }
            }
        }
        return size
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 1 {
                maxSize = max(maxSize, bfs(r, c))
            }
        }
    }
    return maxSize
}