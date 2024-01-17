package main

import "fmt"

func main() {
	s1 := "10101"
	s2 := "00101"
	fmt.Println(strings_xor(s1, s2))
}

func strings_xor(s1 string, s2 string) string {
	var res string
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}
