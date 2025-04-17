package core

import (
	"regexp"
	"testing"
)

var files = []string{
	"test.log",
	"test.tmp",
	"test",
	"test.txt",
	"node_modules/something/a",
	"node_modules/something/b",
	"node_modules/b/c",
	"dist/a/b/c",
	"src/a/b/c",
	"src/a/b/c/d",
	"src/a/b/c.log",
	".config/zsh/.zcompdump",
	".config/zsh/.zcompdump.langoliers-t",
	".config/zsh/.zcompdump.langoliers-t.123123",
	".config/zsh/.zcompdump.langoliers-t.123123.123123",
}

var ignorePatterns = []string{
	"*.log",
	"node_modules/",
	"*.tmp",
	"dist/",
	".zcompdump*",
}

func TestIgnore(t *testing.T) {
	patternStrings := generatePaternStrings(ignorePatterns)
	regexPatterns := make([]*regexp.Regexp, len(patternStrings))

	for i, pattern := range patternStrings {
		p, err := compilePattern(pattern)
		if err != nil {
			t.Fatal(err)
		}
		regexPatterns[i] = p
	}

	for _, file := range files {
		if matchesPattern(file, regexPatterns) {
			t.Logf("Excluding: %s", file)
		} else {
			t.Logf("Including: %s", file)
		}
	}
}
