package main

import (
	"regexp"
	"strings"
)

var sanitizeRegex = regexp.MustCompile("\\??\\.?\\!?")

func Sanitize(s string) string {
	return sanitizeRegex.ReplaceAllString(strings.TrimSpace(s), "")
}
