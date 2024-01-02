package main

import "fmt"

func main() {
	var ar []int

	ar = append(ar, 1, 3, 2, 6, 1, 2)
	res := 0
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if i == j {
				continue
			}
			if (ar[i]+ar[j])%3 == 0 {
				res++
			}
		}
	}

	fmt.Println(res)
}

/*
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    // Write your code here

    if n<2 || n>100{
        return -1
    }
    if k<1 || k>100{
        return -1
    }

    var res int32

    for i := 0; i < len(ar); i++ {
        if ar[i]<1 || ar[i]>100{
            return -1
        }
        for j := i; j < len(ar); j++ {
            if i == j {
                continue
            }
            if (ar[i]+ar[j])%k == 0 {
                res++
            }
        }
    }
    return res
}
*/
