// Package isogram is a small library for determine isogram
package isogram

import (
	"strings"
)

// IsIsogram determine and return if is isogram
func IsIsogram(word string) bool {
	testIsogram := map[string]int{}
	for _, v := range strings.Split(strings.ToLower(word), "") {
		_, ok := testIsogram[v]
		if !ok {
			testIsogram[v]++
		} else {
			if v != " " && v != "-" {
				return false
			}
		}
	}
	return true
}
