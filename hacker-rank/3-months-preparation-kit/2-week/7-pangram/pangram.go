package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize"))
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the prize"))
}

var alphabet = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

const minLetterAmount = 1

func pangrams(s string) string {
	// Write your code here

	lowS := strings.ToUpper(s)
	splitS := strings.ReplaceAll(lowS, " ", "")
	for _, letter := range alphabet {
		letterCounter := 0
		for i := 0; i < len(splitS); i++ {
			if letter == splitS[i] {
				letterCounter++
			}
		}
		if letterCounter < minLetterAmount {
			return "not pangram"
		}
	}
	return "pangram"
}
