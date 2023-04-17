package govkbot

import "strings"

// HasPrefix tests case insensitive whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.ToLower(s[0:len(prefix)]) == strings.ToLower(prefix)
}

func Equals(s1, s2 string) bool {
	return strings.ToLower(strings.TrimSpace(s1)) == strings.ToLower(strings.TrimSpace(s2))
}

// TrimPrefix returns s without the provided leading case insensitive prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s, prefix string) string {
	if HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}
