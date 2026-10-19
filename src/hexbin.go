package src

import (
	"fmt"
	"strconv"
)

func Hex(word string)string {

	if len(word) == 0 {
		return word
	}
		decimalvalue, err := strconv.ParseInt(word, 16, 64)
	if err == nil {
		strBase10 := strconv.FormatInt(decimalvalue, 10)
		fmt.Println("Base 10", strBase10)
		fmt.Printf("Hex %s is decimal %d\n", word, decimalvalue)
		return strBase10
	}
	return word
}

func Bin(word string)string {
	if len(word) == 0 {
		return word
	}
	binairevalue, err := strconv.ParseInt(word, 2, 64)
	if err == nil {
		strBase2 := strconv.FormatInt(binairevalue, 10)
		fmt.Println(strBase2)
		
		return strBase2
	}
	return word
}

