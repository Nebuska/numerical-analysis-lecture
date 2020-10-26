func FindInPascal(line int) []int {
	pascal := make([]int,2)
	pascal[0] = 1
	pascal[1] = 1

	for i := 1; i < line; i++ {
		lenTemp := len(pascal) + 1
		temp := make([]int,lenTemp)
		temp[0] = 1
		temp[lenTemp - 1] = 1
		for j := 1; j < lenTemp - 1;j++ {
			temp[j] = pascal[j - 1] + pascal[j]
		}
		pascal = temp
	}
	return pascal
}
