package grains

import (
	"fmt"
	"math"
)

func Square(number int) (uint64, error) {
	if number > 64 || number <= 0 {
		return 0, fmt.Errorf("64 squares on a chessboard (")
	}
	return uint64(math.Pow(float64(2), float64(number-1))), nil
}

func Total() uint64 {
	return 65536*65536*65536*65536 - 1
}
