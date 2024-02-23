package main

import "sort"

const (
	nonDegenerateTriangleExists = -1
)

func main() {
	sticks := make([]int32, 0)
	sticks = append(sticks, 1, 2, 3)
	maximumPerimeterTriangle(sticks)
}

func maximumPerimeterTriangle(sticks []int32) []int32 {
	// Write your code here
	sidesLength := make([]int32, 0)

	n := len(sticks)

	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] > sticks[j]
	})

	found := false

	for i := 0; i < n-2; i++ {
		if sticks[i] < sticks[i+1]+sticks[i+2] {
			sidesLength = append(sidesLength, sticks[i+2], sticks[i+1], sticks[i])
			found = true
			break
		}
	}

	if !found {
		sidesLength = append(sidesLength, nonDegenerateTriangleExists)
	}

	return sidesLength
}
