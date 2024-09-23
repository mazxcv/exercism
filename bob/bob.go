package bob

import "strings"

// func  Hey imitated teenager answer
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	maybeScream := false
	for a := 'A'; a <= 'Z'; a++ {
		if strings.ContainsRune(remark, a) {
			maybeScream = true
			break
		}
	}

	scream := strings.ToUpper(remark) == remark && maybeScream
	question := remark[len(remark)-1] == '?'

	if scream && question {
		return "Calm down, I know what I'm doing!"
	}
	if scream && !question {
		return "Whoa, chill out!"
	}

	if !scream && question {
		return "Sure."
	}

	return "Whatever."

}
