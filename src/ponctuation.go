package src

import "strings"

func Ponctuation(word, ponctuation string) string {
	word = strings.TrimSpace(word)
	ponctuation = strings.TrimSpace(ponctuation)

	if ponctuation == "" {
		return word
	}
	return word + ponctuation

}
