package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'fizzBuzz' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

const (
	fizz       = "Fizz"
	fizzNumber = 3
	buzz       = "Buzz"
	buzzNumber = 5
)

func fizzBuzz(n int32) {
	// Write your code here

	for i := 1; i <= int(n); i++ {
		if !(i%fizzNumber == 0 || i%buzzNumber == 0) {
			fmt.Println(i)
		} else {
			var res string
			if i%fizzNumber == 0 {
				res += fizz
			}
			if i%buzzNumber == 0 {
				res += buzz
			}
			fmt.Println(res)
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	fizzBuzz(n)
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
