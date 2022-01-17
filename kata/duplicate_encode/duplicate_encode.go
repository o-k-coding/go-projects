package duplicate_encode

import (
	"strings"
)

func DuplicateEncode(arg string) string {
	var sb strings.Builder
	cache := make(map[rune]bool)
	arg = strings.ToLower(arg)
	for _, char := range arg {
		if (cache[char] || strings.Count(arg, string(char)) > 1) {
			sb.WriteRune(')')
			cache[char] = true
		} else {
			sb.WriteRune('(')
		}
	}
	return sb.String()
}
