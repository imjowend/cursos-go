package main

import "fmt"

func main() {
	var s []int32
	s = append(s, 2, 2, 1, 3, 2)
	d := 4
	m := 2
	fmt.Println(birthday(s, int32(d), int32(m)))
}

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here

	//Recorrer con un for, el Array "S" y empezar desde el valor "M"

	// Sumar los valores consiguos desde la posicion "M" hasta la posicion 0

	// Si la suma da igual a "d" sumar el contador

	var counter int32
	k := 0
	for i := int(m) - 1; i < len(s); i++ {
		result := 0
		for j := i; j >= k; j-- {
			result += int(s[j])
		}
		if result == int(d) {
			counter++
		}
		k++
	}

	return counter
}
