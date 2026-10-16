package main

import (
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
    // permet à capitalize de fonctioner
		case "(cap)" : 
		if i > 0 {
			words[i-1] = ucl.Capitalize(words[i-1])
			words[i] = ""
		}
	// permet à Upper de fonctionner 
		case "(up)" : 
		if i > 0 {
			words[i-1] = ucl.Upper(words[i-1])
			words[i] = ""
		}
	// permet à lower de fonctionner 
		case "(low)":
		if i > 0 {
			words[i-1] = ucl.Low(words[i-1])
			words[i] = ""
		}
	// permet à la conversion hexa en decimal
		case "(hex)":
			if i > 0 {
				words[i-1] = ucl.Hex(words[i-1])
				words[i] = ""
			}
		// permet de convertir binaire en decimal
		case "(bin)":
		if i > 0 {
			words[i-1] = ucl.Bin(words[i-1])
			words[i] = ""
		}
	// permet de régler les probleme de ponctuation 
		case ",", ".", "!", "?", "...":
		corection := i - 1
		for corection >= 0 && words[corection] == "" {
			corection--
		}
		if i >= 0 {
			words[corection] = ucl.Ponctuation(words[corection], words[i])
			words[i] = ""
		}

	}

	// Création d'une nouvelle boucle afin de pouvoir corrigée l'index (à plus touchée)
	var indexcorrigée []string

	for i := 0; i < len(words); i++{
		 if words[i] != ""{
			indexcorrigée = append(indexcorrigée, words[i])
		}
	}
	// permet l'écriture de result après le passage de chaque fonction 
		result := strings.Join(indexcorrigée, " " ) 
		result = strings.TrimSpace(result)	

			err = os.WriteFile("result.txt", []byte(result), 0644)
				if err 	!= nil {
					fmt.Println("Erreur d'écriture ", err)
					return
				}
			}
		}	
	