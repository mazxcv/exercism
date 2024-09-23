// Package pangram is a small library for checking if a phrase is a pangram
package pangram

import (
	"strings"
)

// IsPangram checks if a phrase is a pangram
func IsPangram(input string) bool {
	lookup := strings.ToLower(input)
	for i := 'a'; i <= 'z'; i++ {
		if !strings.ContainsRune(lookup, i) {
			return false
		}
	}
	return true
}
