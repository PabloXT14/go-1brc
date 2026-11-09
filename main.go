package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	start := time.Now()

	measurements, err := os.Open("measurements.txt")

	if err != nil {
		panic(err)
	}

	defer measurements.Close() // Close the file descriptor when we're done

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolonIndex := strings.Index(rawData, ";")

		location := rawData[:semicolonIndex]
		rawTemp := rawData[semicolonIndex+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurement, ok := data[location]

		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		data[location] = measurement
	}

	locations := make([]string, 0, len(data))
	for name := range data {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, name := range locations {
		measurement := data[name]

		fmt.Printf(
			"%s=%.1f/%.1f/%.1f}, ",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Printf("}\n")

	fmt.Println(time.Since(start))
}
