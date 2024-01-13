package main

import "fmt"

const (
	SOS       = "SOS"
	SosLength = 3
)

func main() {
	fmt.Println(marsExploration("SOSSPSSQSSOR"))
}

func marsExploration(s string) int32 {
	// Write your code here
	sosSize := len(s) / SosLength
	var sosWord string
	var changedLetter int32
	for i := 0; i < sosSize; i++ {
		sosWord += SOS
	}

	for i := 0; i < len(sosWord); i++ {
		if !(sosWord[i] == s[i]) {
			changedLetter++
		}
	}
	return changedLetter
}
