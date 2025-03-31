package iteration

import "strings"

func Repeat(chars string, n int) string {
	var res strings.Builder
	for i := 0; i < n; i++ {
		res.WriteString(chars)
	}
	return res.String()
	
}
