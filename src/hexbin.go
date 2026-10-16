package src

import (
	"fmt"
	"strconv"
)

func Hex(word string)string {
hexValue := " "
	if len(word) == 0 {
		return word
	}
		decimalvalue, err := strconv.ParseInt(hexValue, 16, 64)
	if err == nil {
		fmt.Println("Hex %s is decimal %d\n", hexValue, decimalvalue)
	}
	return word
}

