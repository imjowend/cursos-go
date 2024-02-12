package main

import "fmt"

func main() {
	// var socks []int32
	// length := 9
	// socks = append(socks, 10, 20, 20, 10, 10, 30, 50, 10, 20)

	// fmt.Println(sockMerchant(int32(length), socks))

	var socks2 []int32
	length2 := 10
	socks2 = append(socks2, 1, 1, 3, 1, 2, 1, 3, 3, 3, 3)

	fmt.Println(sockMerchant(int32(length2), socks2))

}

// func sockMerchant(n int32, ar []int32) int32 {
// 	// Write your code here
// 	var repeats []int32

// 	numOfPairs := 0
// 	for i := 0; i < int(n)-1; i++ {
// 		isRepeat := false
// 		if i >= 1 && len(repeats) >= 1 {
// 			for j := 0; j < len(repeats); j++ {
// 				if ar[i] == repeats[j] {
// 					isRepeat = true
// 					break
// 				}
// 			}
// 		}
// 		if isRepeat {
// 			continue
// 		}
// 		counter := 1
// 		for k := i + 1; k < int(n)-1; k++ {
// 			if ar[i] == ar[k] {
// 				counter++
// 			}
// 		}
// 		pairs := counter / 2
// 		numOfPairs += pairs
// 		if numOfPairs == int(n)/2 {
// 			break
// 		} else {
// 			repeats = append(repeats, ar[i])
// 		}
// 	}

// 	return int32(numOfPairs)
// }

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	fmt.Println(ar)
	pairs := make(map[int32]int32)
	var numPairs int32

	for _, i := range ar {
		fmt.Println(i)
		pairs[i]++
	}
	fmt.Println(pairs)
	for _, v := range pairs {
		numPairs += v / 2
	}

	return numPairs
}
