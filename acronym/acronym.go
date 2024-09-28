// package acronym - is a library to generate acronim
package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

// Abbreviate return acronym
func Abbreviate(s string) string {
	re := regexp.MustCompile(`[a-zA-Z\']+`)
	var sb strings.Builder
	for _, v := range re.FindAllString(s, -1) {
		var c rune
		for _, r := range v {
			c = r
			break
		}
		sb.WriteRune(unicode.ToUpper(c))

	}
	return sb.String()
}
