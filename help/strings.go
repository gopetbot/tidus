package help

import "strings"

func IsEmpty(str string) bool {
	if len(strings.TrimSpace(str)) == 0 {
		return true
	}
	return false
}

func Empty() string {
	return ""
}
