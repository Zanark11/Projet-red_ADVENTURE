package personnage

import "fmt"

func CreationPersonnage() Character {

	var nom string
	var choixclasse string

	fmt.Println("======*CREATION DU PERSONNAGE*======")
	fmt.Println("Donne un nom à ton perso ? ")
	fmt.Scanln(&nom)

	fmt.Println("choisir ta classe:")
	fmt.Println("1. Guerrier")
	fmt.Println("2. Mage")
	fmt.Println("3. Archer")

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
}
func InitCharacter(nom string, choixclasse string) Character {
	pointDeVieActuel := 100

	switch choixclasse {
	case "1":
		pointDeVieActuel = 90
	case "2":
		pointDeVieActuel = 60
	case "3":
		pointDeVieActuel = 70
	}

	character := Character{
		nom:              nom,
		pointDeVieActuel: pointDeVieActuel,
		classe:           choixclasse,
		niveau:           1,
		pointsDeVieMax:   100,
		Inventaire:       [6]string{"potion", "potion", "potion", "", "", ""},
		xp:               0,
	}
	return character
}

type Character struct {
	nom              string
	classe           string
	niveau           int
	pointsDeVieMax   int
	pointDeVieActuel int
	Inventaire       [6]string
	xp               int
}

func JeterObjet(character *Character) {
	var choix int

	fmt.Println("Quel objet aimerais tu jeter ?")
	fmt.Scanln(&choix)
	if choix < 1 || choix > 6 {
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
