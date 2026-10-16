package main

import (
	//"fmt"
	ucl "app/src"
	"fmt"
	"os"
	"strings"
)


func main() {
	if len(os.Args) < 2  {
		os.Exit(1)
	}	
	data, err := os.ReadFile(os.Args[1])
	if err != nil{
		return
	}
	words := strings.Fields(string(data))

	for i := 0; i < len(words); i++{

		switch words[i]{

		case "(cap)" : 
		if i > 0 {
			words[i-1] = ucl.Capitalize(words[i-1])
			words[i] = ""
		}
		case "(up)" : 
		if i > 0 {
			words[i-1] = ucl.Upper(words[i-1])
			words[i] = ""
		}
		case "(low)":
		if i > 0 {
			words[i-1] = ucl.Low(words[i-1])
			words[i] = ""
		}
		case "(hex)":
			if i > 0 {
				words[i-1] = ucl.Hex(words[i-1])
				words[i] = ""
			}
		case "(bin)":
		if i > 0 {
			words[i-1] = ucl.Bin(words[i-1])
			words[i] = " "
		}
	}
	
		result := strings.Join(words, " ", ) 
		result = strings.TrimSpace(result)	

			err = os.WriteFile("result.txt", []byte(result), 0644)
				if err 	!= nil {
					fmt.Println("Erreur d'écriture ", err)
					return
				}
			}
		}	
	