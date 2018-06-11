package raindrops

import "fmt"

func Convert(input int) string {
	output := ""

	if input%3 == 0 {
		output += "Pling"
	}

	if input%5 == 0 {
		output += "Plang"
	}

	if input%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		output = fmt.Sprintf("%d", input)
	}

	return output
}
