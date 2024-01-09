package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	num       = 35601423
	maxLength = 32
	one       = "1"
	zero      = "0"
)

func main() {
	fmt.Println(flippingBits(int64(num)))
}

func flippingBits(n int64) int64 {
	var invertedBinNum string
	newInt := strconv.FormatInt(n, 2)
	newBinNum := completeZeros(newInt)
	for i := 0; i < len(newBinNum); i++ {
		if newBinNum[i] == '0' {
			invertedBinNum += one
		}
		if newBinNum[i] == '1' {
			invertedBinNum += zero
		}
	}
	var entero int64

	for i, j := 0, len(invertedBinNum)-1; i <= len(invertedBinNum)-1; i, j = i+1, j-1 {
		if invertedBinNum[i] == '1' {
			entero += int64(math.Pow(2, float64(j)))
		}
	}
	return entero
}

func completeZeros(str string) string {
	lengthOfZeros := maxLength - len(str)
	var newstr string
	for i := 0; i < lengthOfZeros; i++ {
		newstr += zero
	}

	return newstr + str
}
