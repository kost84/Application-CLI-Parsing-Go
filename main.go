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

	siok := false
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
		corectionponc := i - 1
		for corectionponc >= 0 && words[corectionponc] == "" {
			corectionponc--
		}
		if i >= 0 {
			words[corectionponc] = ucl.Ponctuation(words[corectionponc], words[i])
			words[i] = ""
		}
	// permet de Coller les ' ' autour du ou des mots encadrés
		case "'" :
		
		if siok == false {
			words[i+1] = "'" + words[i+1]
			words[i] = ""
			siok = true
		}else{
		
			words[i-1] = words[i-1] + "'"
			words[i] = ""
			siok = false
		}

		case "a" , "A":
		if i+1  < len(words) {
			if ucl.Verifvoyelleouh(words[i+1]) {
				if words[i] == "A"{
					words[i] = "An"
				}else{
					words[i] = "an"
				}
			}
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
	