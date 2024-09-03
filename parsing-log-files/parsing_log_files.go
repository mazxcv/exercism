package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]+>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`[pP]?[aA]?[sS]{2}[wW]?[oO]?[rR]?[dD]?`)
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
				logs = append(logs, "[USR] "+user[1]+v)
			}
		}
	}
	return logs
}
