package src

import (
	"strings"
	"unicode"
)

func Verifvoyelleouh(word string) bool {
	if word == "" {
		return false
	}
	firstRune := []rune(word)[0]
	firstRune = unicode.ToLower(firstRune)
	return strings.ContainsRune(" a , e , i , o , u", firstRune)
}