func Chio(matrix [][]float64) float64{
	result := 1.0
	for len(matrix) > 2 {
		result *= 1 / math.Pow(matrix[0][0],float64(len(matrix) - 2))
		newSize := len(matrix)-1
		temp := make([][]float64,newSize)
		for i:=0;i < newSize;i++ {
			temp[i] = make([]float64,newSize)
			for j:= 0; j < newSize;j++ {
				temp[i][j] = matrix[0][0] * matrix[i+1][j+1] - (matrix[i+1][0] * matrix[0][j+1])
			}
		}
		matrix = temp
	}
	result *= matrix[0][0] * matrix[1][1] - (matrix[1][0] * matrix[0][1])
	return result
}
