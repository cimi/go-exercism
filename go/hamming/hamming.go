package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("Strands have different lengths %d %d", len(a), len(b))
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance += 1
		}
	}
	return distance, nil
}
