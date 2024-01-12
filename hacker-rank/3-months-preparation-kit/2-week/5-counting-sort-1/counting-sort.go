package main

import "fmt"

func main() {

	var numbers []int32
	numbers = append(numbers, 6, 2, 7, 1, 9, 7, 5, 8, 8, 9)
	fmt.Println(countingSort(numbers))

}

func countingSort(arr []int32) []int32 {
	// Write your code here
	countingArr := make([]int32, 10)

	for i := 0; i < len(arr); i++ {
		countingArr[arr[i]]++
	}
	//Exercise ends here, return countingArr and run
	var sortedArr []int32

	for j, v := range countingArr {
		if v >= 1 {
			for i := 0; i < int(v); i++ {
				sortedArr = append(sortedArr, int32(j))
			}
		}
	}
	return sortedArr
}
