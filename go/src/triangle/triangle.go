// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle implements checking what kind of triangle value of three
// sides create.
package triangle

import (
	"math"
	"sort"
)

// Kind is type for traingle type
type Kind int8

const (
	// NaT : not a triangle
	NaT Kind = iota
	// Equ : equilateral
	Equ
	// Iso : isosceles
	Iso
	// Sca : scalene
	Sca
)

// KindFromSides finds out what kind of triangle three sides form.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case isNaT(a, b, c):
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	case a != b && b != c && a != c:
		k = Sca
	default:
		k = NaT
	}
	return k
}

// isNaT checks if triangle is a indded a triangle
func isNaT(a float64, b float64, c float64) bool {
	return (a == 0 && b == 0 && c == 0) ||
		(a < 0 || b < 0 || c < 0) ||
		inEqT(a, b, c) || mathCheck(a, b, c)

}

// inEqT is a helper function to check for inequality condition
func inEqT(a float64, b float64, c float64) bool {
	tmp := [...]float64{a, b, c}
	sort.Float64s(tmp[:])
	return (tmp[0]+tmp[1] < tmp[2])

}

// mathCheck checks for NaN and infinity values
func mathCheck(a float64, b float64, c float64) bool {
	neg, pos := -1, +1
	return (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)) ||
		(math.IsInf(a, neg) || math.IsInf(b, neg) || math.IsInf(c, neg)) ||
		(math.IsInf(a, pos) || math.IsInf(b, pos) || math.IsInf(c, pos))
}
