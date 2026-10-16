package src

import (
	"strings"
	"unicode"
)

func Capitalize(word string)string {
	if len(word) == 0 {
		return word
	}
	rune := []rune(word)
	rune[0] = unicode.ToUpper(rune[0])
	return string(rune)
}

func Upper(word string)string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word)
}

func Low(word string)string{
	if len(word) == 0 {
		return word 
	}
	return strings.ToLower(word)

}