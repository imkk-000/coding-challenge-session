package main

import "math"

func FindDimension(matrix [][]int) (dimension int) {
	lenght := len(matrix)
	for i := 0; i < lenght; i++ {
		dimension += (matrix[i][i] - matrix[i][lenght-i-1])
	}
	dimension = int(math.Abs(float64(dimension)))
	return
}
