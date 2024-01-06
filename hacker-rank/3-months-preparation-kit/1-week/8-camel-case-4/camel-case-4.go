package main

import (
	"fmt"
	"strings"
	"unicode"
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
	fmt.Println(s)
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
	fmt.Println(s)

	var words []string
	newStrClean := ""
	delimiter := 0
	for i, char := range s {
		//Primero los junto
		if unicode.IsLetter(char) {
			newStrClean += string(char)
		}
		if i > 0 && unicode.IsUpper(char) {
			word := newStrClean[delimiter:i]
			words = append(words, word)
			delimiter = i
		}
	}

	word := newStrClean[delimiter:]
	words = append(words, word)

	capsWords := strings.Join(words, " ")
	strClean := strings.ToLower(capsWords)
	return strClean
}

func combinateString(s string) string {

	newStrClean := strings.ReplaceAll(s, " ", "")
	return newStrClean
}

func methodString(s string) string {
	var res string
	for i, v := range s {
		if unicode.IsUpper(v) && i == 0 {
			res += strings.ToLower(string(v))
		}
		res += string(v)
	}
	return res
}

func classString(s string) string {
	return s
}

func variableString(s string) string {
	return s
}
