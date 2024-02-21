package main

import "fmt"

func main() {
	arr := make([]int32, 0)
	arr = append(arr, 1, 1, 2, 2, 3)
	fmt.Println(migratoryBirds(arr))

}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	birdSeen := map[int32]int32{}
	max := int32(-1)
	maxBird := int32(-1)

	for _, b := range arr {
		birdSeen[b]++
		if birdSeen[b] > max || (birdSeen[b] == max && b < maxBird) {
			max = birdSeen[b]
			maxBird = b
		}
	}

	return maxBird
}
