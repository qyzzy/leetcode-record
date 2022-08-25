package rangeSumQuery2dImmutable

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n, m := len(matrix), len(matrix[0])
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, m+1)
		for j := 0; j < m; j++ {
			sum[i][j+1] = sum[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{
		sum: sum,
	}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) (sum int) {
	for i := row1; i <= row2; i++ {
		sum += nm.sum[i][col2+1] - nm.sum[i][col1]
	}
	return
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
