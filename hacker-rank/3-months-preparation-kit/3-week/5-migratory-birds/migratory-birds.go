package main

import "fmt"

func main() {
	arr := make([]int32, 0)
	arr = append(arr, 1, 1, 2, 2, 3)
	migratoryBirds(arr)
}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	counterIds := make(map[int32]int32)

	for _, v := range arr {
		counterIds[v]++
		fmt.Println(counterIds[v])
	}
	fmt.Println(counterIds)
	return 0
}
