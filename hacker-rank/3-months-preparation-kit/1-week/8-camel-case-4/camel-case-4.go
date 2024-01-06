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
	fmt.Println(camelCase4("S;V;iPad"))
	fmt.Println(camelCase4("C;M;mouse pad"))
	fmt.Println(camelCase4("C;C;code swarm"))
	fmt.Println(camelCase4("S;C;OrangeHighlighter"))

}

func camelCase4(s string) string {
	stringProps := strings.Split(s, semicolon)
	strOperation := stringProps[operation]
	strTypeString := stringProps[typeString]
	strWord := stringProps[word]
	var capsPosition []int
	var res string
	if strOperation == "S" {
		res = separateString(strWord)
		return res
	} else {
		res, capsPosition = combinateString(strWord)
	}
	var newCapsPosition []int

	for i := 0; i < len(capsPosition); i++ {
		newCapsPosition = append(newCapsPosition, (capsPosition[i] - i))
	}
	if strTypeString == "M" {
		res = methodString(res, newCapsPosition)
	} else if strTypeString == "C" {
		res = classString(res, newCapsPosition)
	} else {
		res = variableString(res, newCapsPosition)
	}

	return res
}

func separateString(s string) string {
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

func combinateString(s string) (string, []int) {

	newStrClean := ""
	var capsPosition []int

	for i, char := range s {
		//Primero los junto
		if unicode.IsLetter(char) {
			newStrClean += string(char)
		}
		if unicode.IsSpace(char) {
			capsPosition = append(capsPosition, i)
		}
	}
	//newStrClean := strings.ReplaceAll(s, " ", "")
	return newStrClean, capsPosition
}

func methodString(s string, pos []int) string {
	newStrClean := ""
	for i, char := range s {
		var capLetter string
		letter := string(char)
		for j := 0; j < len(pos); j++ {

			if i == pos[j] {
				capLetter = strings.ToUpper(string(char))
				letter = capLetter
			}
		}
		newStrClean += letter
	}

	return newStrClean + "()"
}

func classString(s string, pos []int) string {
	newStrClean := ""
	for i, char := range s {
		var capLetter string
		letter := string(char)
		if i == 0 {
			letter = strings.ToUpper(string(char))
		}
		for j := 0; j < len(pos); j++ {

			if i == pos[j] {
				capLetter = strings.ToUpper(string(char))
				letter = capLetter
			}
		}
		newStrClean += letter
	}

	return newStrClean
}

func variableString(s string, pos []int) string {
	newStrClean := ""
	for i, char := range s {
		var capLetter string
		letter := string(char)
		for j := 0; j < len(pos); j++ {

			if i == pos[j] {
				capLetter = strings.ToUpper(string(char))
				letter = capLetter
			}
		}
		newStrClean += letter
	}

	return newStrClean
}
