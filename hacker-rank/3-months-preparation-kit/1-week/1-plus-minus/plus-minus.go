package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func plusMinus(arr []int32) {
	positiveRatio := 0
	negativeRatio := 0
	zeroRatio := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positiveRatio += 1
		} else if arr[i] < 0 {
			negativeRatio += 1
		} else {
			zeroRatio += 1
		}

	}

	positive := float32(positiveRatio) / float32(len(arr))
	negative := float32(negativeRatio) / float32(len(arr))
	zero := float32(zeroRatio) / float32(len(arr))

	fmt.Printf("%.6f\n", positive)
	fmt.Printf("%.6f\n", negative)
	fmt.Printf("%.6f\n", zero)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
