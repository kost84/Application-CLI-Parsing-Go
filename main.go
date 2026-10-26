package main

import (
	ucl "app/src"
	"fmt"
	"os"
	"strconv"
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
	ligne := strings.Split(string(data), "\n")
	var findeligne []string

	for _, passageligne := range ligne {
		words := strings.Fields(passageligne)

	siok := false

	for i := 0; i < len(words); i++{

		switch words[i]{
    // permet à capitalize de fonctioner
		case "(cap)": 
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
		case ",", ".", "!", "?", "...", "!!":
		corectionponc := i - 1
		for corectionponc >= 0 && words[corectionponc] == "" {
			corectionponc--
		}
		if corectionponc >= 0 {
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
			corectionapo := i -1
			for corectionapo >= 0 && words[corectionapo] == "'" {
				corectionapo--
			}
			if corectionapo >= 0 {
				words[corectionapo] = words[corectionapo] + "'"
			}
			words[i] = ""
			siok = false
		}
	// permet de changer A ----> An Devant une voyelle ou un h
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
	// 	permet de changer 2 mot en même temps tout en prenant en compte Upper , Capitalize , Lower 
		case "(up,", "(low,", "(cap,":
		if i+1 < len(words) {
			doubletextmajuscule := strings.TrimSuffix(words[i+1], ")")
			chiffre, err := strconv.Atoi(doubletextmajuscule)
			if err == nil {
				lettremodifier := 0
				for j := i-1; j >= 0 && lettremodifier < chiffre; j-- {
					if words[j] != "" {
						if words[i] == "(up,"{
							words[j] = ucl.Upper(words[j])
						}else if words[i] == "(low,"{
							words[j] = ucl.Low(words[j]) 
						}else if words[i] == "(cap," {
							words[j] = ucl.Capitalize(words[j])
						}
						lettremodifier++
					}
				}
			}
			words[i] = ""
			words[i+1] = ""
		}
		
	}
}

	// Création d'une nouvelle boucle afin de pouvoir corrigée l'index (à ne plus touchée)
	var indexcorrigée []string

	for i := 0; i < len(words); i++{
		 if words[i] != ""{
			indexcorrigée = append(indexcorrigée, words[i])
		}
	}
	result := strings.Join(indexcorrigée, " " ) 
		result = strings.TrimSpace(result)	
		findeligne = append(findeligne, result)
	}	

	finalresult := strings.Join(findeligne, "\n")
	err = os.WriteFile("result.txt", []byte(finalresult), 0644)
				if err 	!= nil {
					fmt.Println("Erreur d'écriture ", err)
					return
				}
			}
		