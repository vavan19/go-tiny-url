package utils

import (
	"strings"
)

const (
	base62    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base      = len(base62)
)

// EncodeToBase62 takes an ID and encodes it to a base62 string.
func EncodeToBase62(id int) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(7) // Fixed length for tiny URL

	for id > 0 {
		encodedBuilder.WriteByte(base62[id%base])
		id = id / base
	}
	// Reverse and fill to ensure the string is 7 characters long
	str := encodedBuilder.String()
	reversed := reverseString(str)
	return fillString(reversed, 7)
}

// reverseString reverses the input string.
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// fillString fills the string to the desired length with the first character of the base.
func fillString(str string, length int) string {
	for len(str) < length {
		str = string(base62[0]) + str // Prepend to keep the most significant digit
	}
	return str
}