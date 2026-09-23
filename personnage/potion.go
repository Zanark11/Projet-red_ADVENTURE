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

// TakePoisonPot utilise une potion de poison dans l'inventaire.
// La fonction retourne true si la potion a bien été utilisée.
func TakePoisonPot(character *Character) bool {
	for i := 0; i < len(character.Inventaire); i++ {
		if character.Inventaire[i] == "potion de poison" {
			// On retire la potion de l'inventaire.
			character.Inventaire[i] = ""

			// Les dégâts seront appliqués au Grand Sage
			// par le système de combat.
			fmt.Println("Potion de poison utilisée !")

			return true
		}
	}

	fmt.Println("Vous n'avez pas de potion de poison.")
	return false
}

// TakeManaPot utilise une potion de mana.
// La potion rend 40 mana sans dépasser 100.
// La fonction retourne true si la potion a été utilisée.
func TakeManaPot(character *Character) bool {
	for i := 0; i < len(character.Inventaire); i++ {
		if character.Inventaire[i] == "potion de mana" {
			// On retire la potion de l'inventaire.
			character.Inventaire[i] = ""

			// On ajoute 40 mana.
			character.mana += 40

			// Le mana maximum est de 100.
			if character.mana > 100 {
				character.mana = 100
			}

			fmt.Println("Potion de mana utilisée !")
			fmt.Println("Mana :", character.mana, "/ 100")

			return true
		}
	}

	fmt.Println("Vous n'avez pas de potion de mana.")
	return false
}
