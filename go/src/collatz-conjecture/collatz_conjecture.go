package collatzconjecture

import (
	"errors"
)

// CollatzConjecture implements Collatz Conjecture or 3x+1 problem
// Function takes an integer as input and returns the number of steps
// taken to reach 1 using Collatz Conjecture
func CollatzConjecture(inp int) (int, error) {

	if inp <= 0 {
		return inp, errors.New("Invalid Input")
	}
	r := Rec(inp, 0)
	var n int
	for n = 0; inp > 1; n++ {
		if isEven(inp) {
			inp = inp / 2
		} else {
			inp = 3*inp + 1
		}
	}
	if n == r {
		return n, nil
	}
	return n, errors.New("Recursive and imperative not equall")
}

// isEven is a helper to check if a given number is even
func isEven(n int) bool {
	return n%2 == 0
}

// Rec implements Collatz Conjecture using recursion
func Rec(inp int, n int) int {
	if inp == 1 {
		return n
	}
	if isEven(inp) {
		inp = inp / 2
	} else {
		inp = 3*inp + 1
	}
	return Rec(inp, n+1)
}
