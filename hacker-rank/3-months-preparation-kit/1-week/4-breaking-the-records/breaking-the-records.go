package main

import (
	"fmt"
	"math"
)

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	min := int32(math.MaxInt32)
	max := int32(math.MinInt32)
	var counters []int32
	minCounter := -1
	maxCounter := -1

	for i := 0; i < len(scores); i++ {
		if scores[i] < min {
			min = scores[i]
			minCounter++
		}
		if scores[i] > max {
			max = scores[i]
			maxCounter++
		}
	}

	counters = append(counters, int32(maxCounter), int32(minCounter))
	fmt.Println("Min: ", min, "Cant: ", minCounter)
	fmt.Println("Max: ", max, "Cant: ", maxCounter)

	return counters
}

func main() {

	var scores []int32

	scores = append(scores, 12, 24, 10, 24)
	result := breakingRecords(scores)

	fmt.Println(result)
}
