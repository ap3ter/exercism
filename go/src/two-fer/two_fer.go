// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer implements function ShareWith
package twofer

import "fmt"

// ShareWith retunrs string of form "One for X, one for me."
// When X is a name or "you".
// If the given name is "Alice", the result should be "One for Alice, one for me."
// If no name is given, the result should be "One for you, one for me."
func ShareWith(name string) string {
	var x string
	if name == "" {
		x = "you"
	} else {
		x = name
	}
	return fmt.Sprintf("One for %s, one for me.", x)
}
