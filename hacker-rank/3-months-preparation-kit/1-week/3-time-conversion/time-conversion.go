package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here
	var hour24 string
	if strings.Contains(s, "PM") {
		if !(s[0] == 1 && s[1] == 2) {
			hour := strings.TrimSpace(s[0:2])
			intHour, err := strconv.Atoi(hour)
			if err != nil {
				return ""
			}
			intHour += 12
			newHour := strconv.Itoa(intHour)

			hourChanged := strings.Replace(s, hour, newHour, 1)
			militaryHour := strings.Trim(hourChanged, "PM")
			hour24 = militaryHour
		} else {
			militaryHour := strings.Trim(s, "PM")
			hour24 = militaryHour
		}
	} else if strings.Contains(s, "AM") {
		if s[0] == 1 && s[1] == 2 {
			hour := strings.TrimSpace(s[0:2])
			intHour, err := strconv.Atoi(hour)
			if err != nil {
				return ""
			}
			intHour -= 12
			newHour := strconv.Itoa(intHour)

			hourChanged := strings.Replace(s, hour, newHour, 1)
			militaryHour := strings.Trim(hourChanged, "AM")
			fmt.Println(militaryHour)
			hour24 = militaryHour
		} else {
			militaryHour := strings.Trim(s, "AM")
			hour24 = militaryHour
		}
	}
	return hour24
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
