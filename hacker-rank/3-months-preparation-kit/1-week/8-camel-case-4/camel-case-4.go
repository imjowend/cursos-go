package main

import (
	"fmt"
	"strings"
)

const (
	semicolon  = ";"
	operation  = 0
	typeString = 1
	word       = 2
)

func main() {
	fmt.Println(camelCase4("S;M;plasticCup()"))
	fmt.Println(camelCase4("C;V;mobile phone"))
	fmt.Println(camelCase4("C;C;coffee machine"))
	fmt.Println(camelCase4("S;C;LargeSoftwareBook"))
	fmt.Println(camelCase4("C;M;white sheet of paper"))
	fmt.Println(camelCase4("S;V;pictureFrame"))
}

func camelCase4(s string) string {
	stringProps := strings.Split(s, semicolon)
	strOperation := stringProps[operation]
	strTypeString := stringProps[typeString]
	strWord := stringProps[word]
	var res string
	if strOperation == "S" {
		res = separateString(strWord)
		return res
	} else {
		res = combinateString(strWord)
	}

	if strTypeString == "M" {
		res = methodString(strWord)
	} else if strTypeString == "C" {
		res = classString(strWord)
	} else {
		res = variableString(strWord)
	}

	return res
}

func separateString(s string) string {
	return s
}

func combinateString(s string) string {
	return s
}

func methodString(s string) string {
	return s
}

func classString(s string) string {
	return s
}

func variableString(s string) string {
	return s
}
