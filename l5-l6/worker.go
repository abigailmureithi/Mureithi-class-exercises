package main
import (
	"strings"
)
func Map(key string, value string) []KVPair {

	output := make([]KVPair,0)

    // TODO: This loop iterates over each line of the "value" string
    // You will want to parse out the date and temperature from each line and add it to the "output" slice
	for _, line := range strings.Split(strings.TrimSuffix(value, "\n"), "\n") {
    	// fmt.Println(line)
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			date := strings.TrimSpace(parts[0])
			temp := strings.TrimSpace(parts[1])

			year := date
			if strings.Contains(date, "-") {
				year = strings.Split(date, "-")[0]
			}

	
			output = append(output, KVPair{key: year, value: temp})
	
		}
	}

	return output
}