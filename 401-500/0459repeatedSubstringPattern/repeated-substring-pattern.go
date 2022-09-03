package repeatedSubstringPattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	return strings.Index((s + s)[1:], s) != len(s)-1
}
