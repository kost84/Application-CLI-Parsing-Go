# Application-CLI-Parsing-Go

## installation 

- 1 installer visual studio code
```bash
    https://code.visualstudio.com/download
```
- 2 installer GIT
```bash
    https://git-scm.com/
```
- 3 installer GO
```bash
    https://go.dev/
```
- 4 Clone le repository
```bash
git clone https://github.com/kost84/Application-CLI-Parsing-Go.git
```

## Description du projet

Ce programme lit un fichier texte contenant des mots, il applique une série de règles créé via différentes commandes permettant la correction de texte mot par mot, puis écrit le résultat nettoyé dans un second fichier.

Le projet a été crée en utilisant uniquement la bibliothèque standard de Go.

## Utilisation

Le programme s'exécute depuis le terminal et nécessite deux argument pour sont bon foctionnement :

- 1 . Le fichier contenant du texte brut  `sample.txt`
- 2 . Le fichier ou le texte brut est corrigé `result.txt`

La commande permetant de lancer le main :

```bash
go run . sample.txt result.txt
```
## Règles du programme

Le programme détecte et applique automatiquement les règles suivantes de gauche à droite :

### 1. Conversions hex ---> décimal / Bin ----> décimal

* `(hex)` : Remplace le mot précédent (considéré comme hexadécimal) par sa valeur décimale. `1E (hex)` ➔ `30`

* `(bin)` : Remplace le mot précédent (considéré comme binaire) par sa valeur décimale. `10 (bin)` ➔ `2`

### 2. Gestion de la casse

* `(up)` : Convertit le mot précédent en majuscules. `go (up)` ➔ `GO`

* `(low)` : Convertit le mot précédent en minuscules. * `SHOUTING (low)` ➔ `shouting`

* `(cap)` : Ajoute une majuscule initiale au mot précédent.  `bridge (cap)` ➔ `Bridge`

**Modificateurs multiples :**
Il est possible d'appliquer ces changements à plusieurs mots d'un coup en précisant un nombre `n` : `(up, n)`, `(low, n)`, `(cap, n)`.
`This is so exciting (up, 2)` ➔ `This is SO EXCITING`

### 3. Correction de Grammaire 
* **A / An :** Chaque occurrence isolée de `a` ou `A` est transformée en `an` ou `An` si le mot suivant commence par une voyelle `a`, `e`, `i`, `o`, `u`ou un `h`.  `A amazing rock!` ➔ `An amazing rock!`

### 4. Ponctuation et espacement

* **Signes simples :** Les signes `.`, `,`, `!`, `?`, `:` et `;` sont systématiquement collés au mot précédent et séparés du mot suivant par un espace.

* **Signes groupés :** Les groupes de ponctuation comme `...` ou `!?` sont respectés et formatés ensemble selon la même règle que les signes simples. `I was thinking ... You were right` ➔ `I was thinking... You were right`

* **Apostrophes :** Le signe `'` est traité par paires. Les deux signes viennent encadrer le ou les mots situés entre eux, sans laisser d'espace à l'intérieur.  `' awesome '` ➔ `'awesome'`

## Structure des fichiers

Le code est divisé en plusieurs fichiers pour une meilleure lisibilité et modularité :


```text
.
├── main.go               # Fichier principal 
├── fichiers_test.go      # Fichier regroupant les tests unitaires
├── sample.txt            # Fichier contenant du texte brut
├── result.txt            # Fichier généré après l'exécution
└── src/                  # Dossier contenant les utilitaires 
    ├── hexbin.go         # Fonctions de conversion Hex, Bin
    ├── up-cap-low.go     # Fonctions de Upper, Low, Capitalize
    ├── voyelle.go        # Fonctions de vérification des voyelles/h
    └── ponctuation.go    # Fonction de collage de la ponctuation
```

## Tests Unitaires

Conformément à l'énoncé, le projet inclut un test unitaire par règle de formatage. Les fonctions du package `src` sont testées de manière isolée en utilisant les exemples exacts fournis dans le sujet du projet.

Pour exécuter l'ensemble de la suite de tests :

```bash
go test
```

### Auteur : 
Lioris Racheter