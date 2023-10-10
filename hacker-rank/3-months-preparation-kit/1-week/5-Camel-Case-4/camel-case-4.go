package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var tests []string

	tests = append(tests, "S;M;plasticCup()", "C;V;mobile phone", "C;C;coffee machine", "S;C;LargeSoftwareBook", "C;M;white sheet of paper", "S;V;pictureFrame")
	fmt.Println(tests)

	for i, test := range tests {

		fmt.Println("Test N°"+strconv.Itoa(i)+" \n word: ", test)
		partsOfTest := strings.Split(test, ";")
		for i, part := range partsOfTest {
			//The first part must contains (S or C)
			if i == 1 {
				if part == "S" {
					//Split
				} else {
					//Combinatex
				}
			}
			if i == 2 {
				if part == "M" {
					//Method
				}
				if part == "V" {
					//Variable
				}
				if part == "C" {
					//Class
				}
			}
			// At the step 3 we need to use split or concatenate
			fmt.Println("Part N°"+strconv.Itoa(i)+"\n part: ", part)
		}

	}
	/*fmt.Print("Type a word: ")
	fmt.Scan(&word)
	fmt.Println("Your word is:", word)*/
}
