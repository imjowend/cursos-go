package main

import (
	"fmt"
)

func main() {
	var inputString = []string{"ab", "ab", "abc"}
	var queryString = []string{"ab", "abc", "bc"}

	fmt.Println(matchingStrings(inputString, queryString))
}

func matchingStrings(inputString []string, queryString []string) []int32 {
	// Write your code here
	var res []int32

	for i := 0; i < len(queryString); i++ {
		var count int32
		for j := 0; j < len(inputString); j++ {

			fmt.Println("inputString: ", inputString[j])
			fmt.Println("queryString", queryString[i])
			if inputString[j] == queryString[i] {
				fmt.Println("+1")
				count++
			}
		}
		res = append(res, count)
	}

	return res
}
