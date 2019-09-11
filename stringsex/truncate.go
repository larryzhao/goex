package stringsex

import "fmt"

// Truncate truncate string s to length len from beginning
func Truncate(s string, length int, suffix string) string {
	if len(s) == 0 || len(s) <= length {
		return s
	}

	rr := []rune(s)

	suffixLen := len(suffix)
	return fmt.Sprintf("%s%s", string(rr[0:length-suffixLen]), suffix)
}
