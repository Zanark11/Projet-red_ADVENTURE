package personnage

import "fmt"

// TakePot utilise une potion de vie.
func TakePot(character *Character) {
	for i := 0; i < len(character.Inventaire); i++ {

		if character.Inventaire[i] == "potion de vie" {

			// On retire la potion de l'inventaire.
			character.Inventaire[i] = ""

			// La potion rend 50 PV.
			character.pointDeVieActuel += 50

			// On ne peut pas dépasser les PV maximum.
			if character.pointDeVieActuel > character.pointsDeVieMax {
				character.pointDeVieActuel = character.pointsDeVieMax
			}

			fmt.Println("Potion de vie utilisée !")
			fmt.Println("PV :", character.pointDeVieActuel, "/", character.pointsDeVieMax)

			return
		}
	}

	fmt.Println("Vous n'avez pas de potion de vie.")
}

// TakePoisonPot utilise une potion de poison.
func TakePoisonPot(character *Character) bool {
	for i := 0; i < len(character.Inventaire); i++ {

		if character.Inventaire[i] == "potion de poison" {

			// On retire la potion de l'inventaire.
			character.Inventaire[i] = ""

			fmt.Println("Potion de poison utilisée sur le Grand Sage !")

			return true
		}
	}

	fmt.Println("Vous n'avez pas de potion de poison.")
	return false
}