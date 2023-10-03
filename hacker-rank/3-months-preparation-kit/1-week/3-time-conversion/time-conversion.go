package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	var hour24 string
	if strings.Contains(s, "PM") {
		if !(s[0:2] == "12") {
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
		if s[0:2] == "12" {
			hour := strings.TrimSpace(s[0:2])
			hourChanged := strings.Replace(s, hour, "00", 1)
			militaryHour := strings.Trim(hourChanged, "AM")
			hour24 = militaryHour
		} else {
			militaryHour := strings.Trim(s, "AM")
			hour24 = militaryHour
		}
	}
	return hour24
}

func main() {

	result := timeConversion("07:01:01AM")

	fmt.Println(result)

}
