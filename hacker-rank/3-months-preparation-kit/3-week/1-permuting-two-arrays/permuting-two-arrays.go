package main

import (
	"fmt"
	"math"
	"sort"
)

const (
	NO_PERMUTING    = "NO"
	YES_PERMUTING   = "YES"
	MIN             = math.MinInt64
	ACTUAL_POSITION = 1
	MIN_VALUE_FOUND = 1
)

func main() {
	var a []int32
	var b []int32
	a = append(a, 0, 1)
	b = append(b, 0, 2)
	fmt.Println(twoArrays(1, a, b))
	var c []int32
	var d []int32
	c = append(c, 0)
	d = append(d, 0)
	fmt.Println(twoArrays(0, c, d))
	var e []int32
	var f []int32
	e = append(e, 2, 1, 3)
	f = append(f, 7, 8, 9)
	fmt.Println(twoArrays(10, e, f))
}

func twoArrays(k int32, A []int32, B []int32) string {
	// Write your code here

	// Ordena el slice en orden descendente
	sort.SliceStable(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	sort.SliceStable(B, func(i, j int) bool {
		return B[i] > B[j]
	})
	// aNum := searchUpperNumber(A)
	// bNum := searchUpperNumber(B)
	// aAsc := countingSortAsc(A, aNum)
	// bDesc := countingSortDesc(B, bNum)
	// for i := range aAsc {
	// 	if aAsc[i]+bDesc[i] < k {
	// 		return NO_PERMUTING
	// 	}
	// }
	// return YES_PERMUTING
	for i := range A {
		if A[i]+B[i] < k {
			return NO_PERMUTING
		}
	}
	return YES_PERMUTING
}

func searchUpperNumber(arr []int32) int32 {
	number := MIN
	for _, v := range arr {
		if v > int32(number) {
			number = int(v)
		}
	}

	return int32(number)
}

func countingSortDesc(arr []int32, length int32) []int32 {
	countingArr := make([]int32, length+ACTUAL_POSITION)
	for i := 0; i < len(arr); i++ {
		countingArr[arr[i]]++
	}
	var sortedArr2 []int32
	for i := len(countingArr) - 1; i >= 0; i-- {
		if countingArr[i] >= MIN_VALUE_FOUND {
			for j := 0; j < int(countingArr[i]); j++ {
				sortedArr2 = append(sortedArr2, int32(i))
			}
		}
	}
	return sortedArr2
}

func countingSortAsc(arr []int32, length int32) []int32 {
	countingArr := make([]int32, length+ACTUAL_POSITION)
	for i := 0; i < len(arr); i++ {
		countingArr[arr[i]]++
	}
	var sortedArr []int32
	for j, v := range countingArr {
		if v >= MIN_VALUE_FOUND {
			for i := 0; i < int(v); i++ {
				sortedArr = append(sortedArr, int32(j))
			}
		}
	}
	return sortedArr
}
