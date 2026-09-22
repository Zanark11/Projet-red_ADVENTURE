package menuexploration

import (
	"aventure/personnage"
	"fmt"
)

func Explorer(perso *personnage.Character) {
	niveau2 := false
	niveau3 := false
for{
var choix int

	fmt.Println("=====EXPLORATION====")
	fmt.Println()
	fmt.Println("1.Niveau n°1")
	fmt.Println("2.Niveau n°2")
	fmt.Println("3.Niveau n°3")
	fmt.Println("4.Retour")

	fmt.Scanln(&choix)
	switch choix{
	case 1:
		fmt.Println("Bienvenu dans l'aventure.")
	case 2:
		if !niveau2 {
			fmt.Println("Débloquer le niveau 1 pour passer au niveau 2")
			continue
		}
	case 3:
		if !niveau3 {
			fmt.Println("jouer le niveau 1 et niveau 2 pour débloquer le niveau 3")
				continue 
			}
	case 4:
          return
	default:
		fmt.Println("choix invalide")
		}
    }
}
