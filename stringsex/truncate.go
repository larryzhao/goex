package stringsex

// Truncate truncate string s to length len from beginning
func Truncate(s string, length int) string {
	if len(s) == 0 || len(s) <= length {
		return s
	}

	rr := []rune(s)

	return string(rr[0:length])
}
