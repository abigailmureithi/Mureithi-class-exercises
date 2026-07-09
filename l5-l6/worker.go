package main

import (
	"strconv"
	"strings"
)

func Map(key string, value string) []KVPair {

	output := make([]KVPair, 0)

	// TODO: This loop iterates over each line of the "value" string
	// You will want to parse out the date and temperature from each line and add it to the "output" slice
	for _, line := range strings.Split(strings.TrimSuffix(value, "\n"), "\n") {
		// fmt.Println(line)
		parts := strings.Split(line, ",")
		if len(parts) >= 3 {
			date := strings.TrimSpace(parts[1])
			temp := strings.TrimSpace(parts[2])

			year := date
			if strings.Contains(date, "-") {
				year = strings.Split(date, "-")[0]
			}

			output = append(output, KVPair{key: year, value: temp})

		}
	}

	return output
}

func Reduce(key string, value []string) float64 {
	// Converting from a string to float may be useful
	// val,err := strconv.ParseFloat(INPUT, 64)

	max := -999.0

	for _, valStr := range value {
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			continue
		}

		if val > max {
			max = val
		}
	}

	return max
}
