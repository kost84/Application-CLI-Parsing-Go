package main

import (
	//"fmt"
	"os"
	"strings"
	ucl  "app/src"
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
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = ucl.Capitalize(words[i-1]) 
			words[i] = ""
		}

		if words[i] =="(up)" && i > 0 {
			words[i-1] = ucl.Upper(words[12])
			words[i] = ""
		}

		if words[i] =="(low)" && i > 0 {
			words[i-1] = ucl.Low(words[20])
			words[i] = ""
		}
	
	result := strings.Join(words, " ")
	

			err = os.WriteFile("result.txt", []byte(result), 0644)
				if err 	!= nil {
					return
				}
			}
		}	