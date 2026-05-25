package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	pattern := `^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])\s+`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return regex.MatchString(text)
}

func SplitLogLine(text string) []string {
	pattern := `<[~\*=-]*>`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil
	}
	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var counter int
	pattern := `".*(?i:password).*"`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return 0
	}
	for line := range lines {
		if regex.MatchString(lines[line]) {
			counter++
		}
	}
	return counter
}
func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line\d+`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return ""
	}
	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `User\s+(\b[a-zA-Z0-9]+\b)`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil
	}
	for line := range lines {
		str := regex.FindStringSubmatch(lines[line])
		if len(str) > 0 {
			lines[line] = "[USR] " + str[1] + " " + lines[line]
		}
	}
	return lines
}
