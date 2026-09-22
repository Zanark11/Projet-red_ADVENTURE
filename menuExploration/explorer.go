package menuexploration

import (
	"aventure/personnage"
	"fmt"
)

func Explorer(perso *personnage.Character) {
	niveau2 := false

	for {
		var choix int

		fmt.Println("=====EXPLORATION====")
		fmt.Println()
		fmt.Println("1.Niveau n°1")
		fmt.Println("2.Niveau n°2")
		fmt.Println("3.Retour")

		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Println("Bienvenu dans l'aventure.")
		case 2:
			if !niveau2 {
				fmt.Println("Débloquer le niveau 1 pour passer au niveau 2")
				continue
			}
		case 3:
			return
		default:
			fmt.Println("choix invalide")
		}
	}
}
