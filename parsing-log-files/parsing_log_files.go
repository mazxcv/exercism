package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\]`)
	indexes := re.FindStringIndex(text)
	return indexes != nil && indexes[0] == 0
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`[pP]{1}[aA]{1}[sS]{2}[wW]{1}[oO]{1}[rR]{1}[dD]{1}`)
	sum := 0
	for _, v := range lines {
		if IsValidLine(v) {
			if re.MatchString(v) {
				sum++
			}
		}
	}
	return sum
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+([\w\d]+)`)
	logs := []string{}
	for _, v := range lines {
		if IsValidLine(v) {
			user := re.FindStringSubmatch(v)
			if user != nil {
				logs = append(logs, "[USR] "+user[1]+" "+v)
			} else {
				logs = append(logs, v)
			}
		}
	}
	return logs
}
