package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var builder strings.Builder
	builder.Grow(len(character) * repeatCount)
	for range repeatCount {
		builder.WriteString(character)
	}
	return builder.String()
}
