package personnage

import (

	"fmt"
)

func CreationPersonnage() Character {

	var nom string
	var choixclasse string
	for {
		valide := true
		fmt.Println("======*CREATION DU PERSONNAGE*======")
		fmt.Println("Donne un nom à ton perso ? ")
		fmt.Scanln(&nom)

		if nom == "" {
			fmt.Println("le nom ne peut pas être vide")
			continue
		}
		premierelettre := nom[0]

		for _, lettre := range nom {
			if lettre >= '0' && lettre <= '9' {
				valide = false
			}
		}
		if !valide {
			fmt.Println("votre ne peut contenir de chiffre")
			continue
		}
		if premierelettre >= 'a' && premierelettre <= 'z' {
			premierelettre = premierelettre - ('a' - 'A')
			nom = string(premierelettre) + nom[1:]
		}
		break
	}
	fmt.Println("choisir ta classe:")
	fmt.Println("1. Avocat")
	fmt.Println("2. Informaticien")

	fmt.Println("3. Medecin")

	fmt.Scanln(&choixclasse)

	character := InitCharacter(nom, choixclasse)
	DisplayInfo(character)

	return character
}

func DisplayInfo(character Character) {
	fmt.Println()
	fmt.Println("félicitation personnage crée !")
	fmt.Println("nom :", character.nom)
	fmt.Println("pointDeVieActuel:", character.pointDeVieActuel)
	fmt.Println("classe:", character.classe)
	fmt.Println("niveau:", character.niveau)
	fmt.Println("inventaire:", character.Inventaire)
	fmt.Println("xp:", character.xp)
	fmt.Println("mana:", character.mana)
	fmt.Println("Argent", character.Argent)
	fmt.Println("Equipement", character.equipement)
}
func InitCharacter(nom string, choixclasse string) Character {
	pointDeVieActuel := 100

	switch choixclasse {
	case "1":
		pointDeVieActuel = 120
	case "2":
		pointDeVieActuel = 110
	case "3":
		pointDeVieActuel = 100
	}


	equipement := Porterequipement()
	character := Character{
		nom:              nom,
		pointDeVieActuel: pointDeVieActuel,
		classe:           choixclasse,
		niveau:           1,
		pointsDeVieMax:   500,
		Inventaire:       []string{"potion de vie,", "potion de vie,", "potion de vie,", "", "", ""},
		xp:               0,
		mana:            100,
		Argent:           100,
		equipement:       equipement,
		
	}
	return character
}

type Character struct {
	nom              string
	classe           string
	niveau           int
	pointsDeVieMax   int
	pointDeVieActuel int
	Inventaire       []string
	xp               int
	mana             int
	Argent           int
	equipement       Equipement
	
}

func JeterObjet(character *Character) {
	var choix int

	fmt.Println("Quel objet aimerais tu jeter ?")
	fmt.Scanln(&choix)
	if choix < 1 || choix > len(character.Inventaire) {
		fmt.Println("choix invalide")
		return
	}
	index := choix - 1
	if character.Inventaire[index] == "" {
		fmt.Println("case déja vide .")
		return
	}
	objet := character.Inventaire[index]

	character.Inventaire[index] = ""

	fmt.Println("vous avez jetez:", objet)
}

func Degats(perso *Character, degats int){
   perso.pointDeVieActuel -= degats 
}

func Affichevie(perso *Character){
	fmt.Println("vie",perso.pointDeVieActuel, "/", perso.pointsDeVieMax,)
}

func Playervivant(perso *Character) bool {
	return perso.pointDeVieActuel > 0
}

type Equipement struct{
	Tête  string
	Torse  string
	Pieds  string
}

func Porterequipement() Equipement{
	var choix int
	var equipement Equipement

	fmt.Println("===Equipement pour l'aventure===")
	fmt.Println()
	fmt.Println("choisir ton vêtement pour le Torse")
	fmt.Println("1.Un débardeur")
	fmt.Println("2.Un tee-shirt")
	fmt.Println("3.Aucun")

	fmt.Scanln(&choix)

	switch choix {
	case 1:
		equipement.Torse="Un débardeur"
	case 2:
		equipement.Torse="Un tee-shirt"
	case 3:
		equipement.Torse="Aucun"
	}


	fmt.Println("choisir ton vêtement pour la tête")
	fmt.Println("1.Une casquette")
	fmt.Println("2.un bornet")
	fmt.Println("3.Aucun")

	fmt.Scanln(&choix)
	switch choix{
	case 1:
		equipement.Tête="Une casquette"
	case 2:
		equipement.Tête="Un bornet"
	case 3:
		equipement.Tête="Aucun"
	}

	fmt.Println("choisir tes chaussures")
	fmt.Println("1.Des basquette")
	fmt.Println("2.Des sandale")
	fmt.Println("3.Aucun")

    fmt.Scanln(&choix)

    switch choix{
	case 1:
		equipement.Pieds="Des basquette"
	case 2:
		equipement.Pieds="Des sandale"
	case 3:
		equipement.Pieds="Aucun"
	}
	return equipement
}

