// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob implements the lackadaisical teenager Bob
package bob

import (
	"strings"
)

// Hey implements the lackadaisical teenager Bob
// Bob answers 'Sure.' if you ask him a question.
// He answers 'Whoa, chill out!' if you yell at him.
// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
// He says 'Fine. Be that way!' if you address him without actually saying
// anything.
// He answers 'Whatever.' to anything else.
// Bob's conversational partner is a purist when it comes to written communication
// and always follows normal rules regarding sentence punctuation in English.
func Hey(remark string) string {
	var response string
	switch {
	case (isaquestion(remark) && (strings.ToUpper(remark) == remark) && !isnumberonly(remark)):
		response = "Calm down, I know what I'm doing!"
	case isaquestion(remark):
		response = "Sure."
	case strings.ToUpper(remark) == remark && !isnumberonly(remark):
		response = "Whoa, chill out!"
	case isempty(remark):
		response = "Fine. Be that way!"
	case !isnumberonly(remark):
		response = "Whatever."
	default:
		response = "Whatever."

	}
	return response
}

func isnumberonly(str string) bool {
	for _, r := range str {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return false
		}
	}
	return true
}

func isaquestion(str string) bool {
	if strings.HasSuffix(strings.Trim(str, " "), "?") {
		return true
	}
	return false
}

func isempty(str string) bool {
	trimmed := strings.Replace(
		strings.Replace(
			strings.Replace(
				strings.Replace(
					str, " ", "", -1), "\t", "", -1), "\n", "", -1), "\r", "", -1)
	if len(trimmed) == 0 {
		return true
	}
	return false
}
