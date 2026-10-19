package main

import (
    "testing"
    
	ucl "app/src"
)

func TestHex(t *testing.T) {
    resultatObtenu := ucl.Hex("1E")
    resultatAttendu := "30"
    if resultatObtenu != resultatAttendu {
        t.Errorf("Erreur : j'attend %s, j'ai obtenue %s", resultatAttendu, resultatObtenu)
    }
}

func TestBin(t *testing.T) {
    resultatObtenu := ucl.Bin("10")
    resultatAttendu := "2"
    if resultatObtenu != resultatAttendu {
        t.Errorf("Erreur : j'attend %s, j'ai obtenue %s", resultatAttendu, resultatObtenu)
    }
}

func TestPonctuation(t *testing.T) {
    resultatObtenu := ucl.Ponctuation("Bonjour", "!")
    resultatAttendu := "Bonjour!"
    if resultatObtenu != resultatAttendu {
        t.Errorf("Erreur : j'attend %s, j'ai obtenue %s", resultatAttendu, resultatObtenu)
    }
}

func TestVerifvoyelleouh(t *testing.T) {
    resultatObtenu := ucl.Verifvoyelleouh("A")
    resultatAttendu := true
    if resultatObtenu != resultatAttendu {
        t.Errorf("Erreur : j'attend %t, j'ai obtenue %t", resultatAttendu, resultatObtenu)
    }
}

func TestCapitalize(t *testing.T) {
    resultatObtenu := ucl.Capitalize("it")
    resultatAttendu := "It"
    if resultatObtenu != resultatAttendu {
        t.Errorf("Erreur : j'attend %s, j'ai obtenue %s", resultatAttendu, resultatObtenu)
    }
}

func TestUpper(t *testing.T) {
    resultatObtenu := ucl.Upper("time")
    resultatAttendu := "TIME"
    if resultatObtenu != resultatAttendu {
        t.Errorf("Erreur : j'attend %s, j'ai obtenue %s", resultatAttendu, resultatObtenu)
    }
}


func TestLow(t *testing.T) {
    resultatObtenu := ucl.Low("SHOUTING")
    resultatAttendu := "shouting"
    if resultatObtenu != resultatAttendu {
        t.Errorf("Erreur : j'attend %s, j'ai obtenue %s", resultatAttendu, resultatObtenu)
    }
}

