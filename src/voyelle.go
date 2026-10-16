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

func ajoutvoyelle(voyelle string) string {
	words := strings.Fields(voyelle)
	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			if Verifvoyelleouh(words[i+1]) {
				if words[i] == "A" {
						words[i] = "An"
				}else{
					words[i] = "an"
				}
			}
		}
	}
	return strings.Join(words, " ")
}