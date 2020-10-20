package str

func
Chop(s string, n int) string {
	l := len(s)
	if l<n {
		return ""
	}
	return s[:l-n]
}
