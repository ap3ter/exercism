package grains

import (
	"fmt"
	"math"
)

// Square provides the number of grains in a sqaure provide the number  of
// square. Given number shoul be between 1 and 64
// For a given square n number of grains can be calculated as
// 2**(n-1)
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("Incorrect sqaure number %v", n)
	}

	return uint64(math.Pow(2, float64(n-1))), nil
}

// Total finds total number of grains that can be in a chess board
func Total() uint64 {
	var t uint64
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		t = t + s
	}
	return t
}
