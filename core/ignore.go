package core

import (
	"regexp"
	"strings"
)

func padBeginning(str string) (string, bool) {
	if str[0:1] == "*" {
		return str, false
	}

	return "*/" + str, true
}

func padEnd(str string) (string, bool) {
	lastChar := str[len(str)-1:]

	if lastChar == "*" {
		return str, false
	}

	if lastChar == "/" {
		return str + "*", true
	}

	return str + "/*", true
}

func generatePaternStrings(ignorePatterns []string) []string {
	allPatterns := []string{}

	for _, pattern := range ignorePatterns {
		allPatterns = append(allPatterns, pattern)
		padB, ok := padBeginning(pattern)
		if ok {
			allPatterns = append(allPatterns, padB)
			padBE, ok := padEnd(padB)
			if ok {
				allPatterns = append(allPatterns, padBE)
			}
		}
		padE, ok := padEnd(pattern)
		if ok {
			allPatterns = append(allPatterns, padE)
		}
	}
	return allPatterns
}

func compilePattern(pattern string) (*regexp.Regexp, error) {
	// Escape special characters and replace '*' with '.*'
	regexPattern := regexp.QuoteMeta(pattern)
	regexPattern = strings.ReplaceAll(regexPattern, `\*`, `.*`)

	regexPattern = "^" + regexPattern + "$"

	return regexp.Compile(regexPattern)
}

func matchesPattern(path string, patterns []*regexp.Regexp) bool {
	for _, re := range patterns {
		if re.MatchString(path) {
			return true
		}
	}
	return false
}
