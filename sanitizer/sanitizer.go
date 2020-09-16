package sanitizer

import(
	"strings"
)

func
Sanitize(input []byte) []byte {
	input = []byte(strings.ReplaceAll(string(input), "\b", ""))
	return input
}
