// util.go
package alc_shared

import "strings"

// ReverseString reverses a given string.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ToUpperCase converts a string to uppercase.
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}
