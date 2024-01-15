package main

import "fmt"

func main() {
	var a []int32
	var b []int32
	a = append(a, 0, 1)
	b = append(b, 0, 2)
	fmt.Println(twoArrays(1, a, b))
}

func twoArrays(k int32, A []int32, B []int32) string {
	// Write your code here
	return ""
}

func countingSortDesc(arr []int32) []int32 {
	// Write your code here
	countingArr := make([]int32, 10)

	for i := 0; i < len(arr); i++ {
		countingArr[arr[i]]++
	}

	//Exercise ends here, return countingArr and run

	var sortedArr2 []int32

	for i := len(countingArr) - 1; i >= 0; i-- {

		if countingArr[i] >= 1 {

			for j := 0; j < int(countingArr[i]); j++ {

				sortedArr2 = append(sortedArr2, int32(i))
			}
		}
	}

	return sortedArr2
}

func countingSortAsc(arr []int32) []int32 {
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
