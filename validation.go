package main

import (
	"strings"

	"github.com/agnivade/levenshtein"
)

func IsCorrect(expect, actual string) bool {
	expect = Sanitize(strings.ToLower(expect))
	actual = strings.ToLower(actual)

	dist := levenshtein.ComputeDistance(expect, actual)

	return dist <= 2
}
