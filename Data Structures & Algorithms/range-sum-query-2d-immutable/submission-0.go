type NumMatrix struct {
    prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    rows, cols := len(matrix), len(matrix[0])
    // размер (rows+1)×(cols+1), нулевые строка и столбец = 0
    prefix := make([][]int, rows+1)
    for i := range prefix {
        prefix[i] = make([]int, cols+1)
    }

    for i := 1; i <= rows; i++ {
        for j := 1; j <= cols; j++ {
            prefix[i][j] = matrix[i-1][j-1] +
                prefix[i-1][j] +
                prefix[i][j-1] -
                prefix[i-1][j-1]
        }
    }
    return NumMatrix{prefix: prefix}
}

func (this *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
    return this.prefix[row2+1][col2+1] -
        this.prefix[row1][col2+1] -
        this.prefix[row2+1][col1] +
        this.prefix[row1][col1]
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
