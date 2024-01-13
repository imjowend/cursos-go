package main

import "fmt"

const (
	uphill   = 'U'
	downhill = 'D'
	seaLevel = 0
)

func main() {

	fmt.Println(countingValleys(8, "DDUUUUDD"))
}

func countingValleys(steps int32, path string) int32 {
	altitude := 0
	valleyCounter := 0

	for i := 0; i < int(steps); i++ {
		actualLevel := altitude
		if path[i] == uphill {
			altitude++
		}
		if path[i] == downhill {
			altitude--
		}

		if altitude < seaLevel && actualLevel == seaLevel {
			valleyCounter++
		}
	}
	return int32(valleyCounter)
}

// si estuvo negativo (abajo del mar)  y luego sube a 0 cuenta como que haya estado a un valle
