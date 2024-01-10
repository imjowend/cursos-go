package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}

	fmt.Println(diagonalDifference(matrix))

}

func diagonalDifference(arr [][]int32) int32 {
	var result, firstDiag, secondDiag int32

	for i, j := 0, (len(arr[0]) - 1); i < len(arr[0]); i, j = i+1, j-1 {
		firstDiag += arr[i][i]
		secondDiag += arr[i][j]
	}

	result = int32(math.Abs(float64(firstDiag) - float64(secondDiag)))

	return result
}
