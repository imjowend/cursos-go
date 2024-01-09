package main

import (
	"fmt"
	"strconv"
)

const (
	num = 2147483647
)

func main() {
	number := 10
	fmt.Println(flippingBits(int64(number)))
}

func flippingBits(n int64) int64 {
	// Write your code here
	//cacmbiar a decimal

	newInt := strconv.FormatInt(n, 2)
	b, _ := strconv.Atoi(newInt)
	newBinNum := fmt.Sprintf("%032d", int64(b))
	fmt.Println(newBinNum)
	a, _ := strconv.Atoi(newBinNum)

	return int64(a)
}
