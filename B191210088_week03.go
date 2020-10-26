func Sarrus(matrix [3][3]float64) float64 {
	sum := 0.0
	for a := 0;a < 3;a++ {
		T1 := 1.0
		T2 := 1.0
		for i := 0;i < 3;i++ {
			j := a+i
			if j >= 3 {
				j = j-3
			}
			T1 *= matrix[i][j]
			j = 2 - j
			T2 *= matrix[i][j]
		}
		sum += T1
		sum -= T2
	}
	return sum
}
