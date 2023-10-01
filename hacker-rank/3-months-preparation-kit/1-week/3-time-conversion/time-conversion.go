package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here

	var hour24 string
	if strings.Contains(s, "PM") {
		if !(s[0] == 1 && s[1] == 2) {
			hour := strings.TrimSpace(s[0:2])
			fmt.Println("hour: ", hour)
			intHour, err := strconv.Atoi(hour)
			fmt.Println("intHour: ", intHour)
			if err != nil {
				return ""
			}
			intHour += 12
			newHour := strconv.Itoa(intHour)
			fmt.Println("newHour: ", newHour)
			hourChanged := strings.Replace(s, hour, newHour, 1)
			fmt.Println("hourChanged: ", hourChanged)
			militaryHour := strings.Trim(hourChanged, "PM")
			fmt.Println("militaryHour", militaryHour)
			hour24 = militaryHour
			fmt.Println("hour24:", hour24)

		} else {
			militaryHour := strings.Trim(s, "PM")
			fmt.Println("militaryHour", militaryHour)

			hour24 = militaryHour
			fmt.Println("hour24:", hour24)

		}
	} else if strings.Contains(s, "AM") {
		if s[0] == 1 && s[1] == 2 {
			hour := strings.TrimSpace(s[0:2])
			fmt.Println("hour: ", hour)

			intHour, err := strconv.Atoi(hour)
			fmt.Println("intHour: ", intHour)

			if err != nil {
				return ""
			}
			intHour -= 12
			newHour := strconv.Itoa(intHour)
			fmt.Println("newHour: ", newHour)

			hourChanged := strings.Replace(s, hour, newHour, 1)
			fmt.Println("hourChanged: ", hourChanged)

			militaryHour := strings.Trim(hourChanged, "AM")
			fmt.Println(militaryHour)
			hour24 = militaryHour
			fmt.Println("hour24", hour24)

		} else {
			militaryHour := strings.Trim(s, "AM")
			hour24 = militaryHour
			fmt.Println("hour24", hour24)

		}
	}
	return hour24
}

func main() {

	result := timeConversion("12:01:01PM")

	fmt.Println(result)

}
