package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Measurement struct {
	Name  string
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("measurements.txt")

	if err != nil {
		panic(err)
	}

	defer measurements.Close() // Close the file descriptor when we're done

	scanner := bufio.NewScanner(measurements)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolonIndex := strings.Index(rawData, ";")

		location := rawData[:semicolonIndex]
		temp := rawData[semicolonIndex+1:]

		fmt.Println(location, temp)
	}
}
