package personnage

import "fmt"

// TakePot utilise une potion de vie.
func TakePot(character *Character) {
	// On cherche une potion de vie dans l'inventaire.
	for i := 0; i < len(character.Inventaire); i++ {
		if character.Inventaire[i] == "potion de vie" {

			// La potion est consommée.
			character.Inventaire[i] = ""

			// La potion rend 50 PV.
			character.pointDeVieActuel += 50

			// Les PV ne peuvent pas dépasser le maximum.
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

// TakePoisonPot cherche et consomme une potion de poison.
//
// La fonction ne retire PAS de PV au joueur.
// Les dégâts seront appliqués au Grand Sage par le système de combat.
func TakePoisonPot(character *Character) bool {
	// On cherche une potion de poison dans l'inventaire.
	for i := 0; i < len(character.Inventaire); i++ {
		if character.Inventaire[i] == "potion de poison" {

			// La potion est consommée.
			character.Inventaire[i] = ""

			fmt.Println("Potion de poison utilisée sur le Grand Sage !")

			// true signifie que la potion a bien été utilisée.
			return true
		}
	}

	fmt.Println("Vous n'avez pas de potion de poison.")

	// false signifie qu'aucune potion n'a été trouvée.
	return false
}