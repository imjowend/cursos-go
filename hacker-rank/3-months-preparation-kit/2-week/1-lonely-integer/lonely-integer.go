package main

import "fmt"

func main() {
	a := []int32{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(lonelyInteger(a))
}

func lonelyInteger(a []int32) int32 {
	var possibleUnique int32
	if len(a) == 1 {
		possibleUnique = a[0]
		return possibleUnique
	}
	for i := 0; i < len(a); i++ {
		var point int32
		for j := 0; j < len(a); j++ {
			if a[i] != a[j] {
				possibleUnique = a[i]
				point++
			}
		}
		if int(point) == len(a)-1 {
			break
		}
	}
	return possibleUnique
}
