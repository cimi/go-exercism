package grains

import (
	"errors"
	"math"
)

var board [65]uint64

func Square(pos int) (uint64, error) {
	if pos < 1 || pos > 64 {
		return 0, errors.New("Invalid position")
	}
	if board[pos] == 0 {
		board[pos] = uint64(math.Pow(2.0, float64(pos-1)))
	}
	return board[pos], nil
}

var total uint64 = 0

func Total() uint64 {
	if total == 0 {
		for i := 1; i <= 64; i++ {
			val, _ := Square(i)
			total += val
		}
	}
	return total
}
