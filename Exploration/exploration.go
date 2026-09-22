package Exploration

import (
	"aventure/menuexploration"
	"aventure/personnage"
	"fmt"
)

func Exploration(perso *personnage.Character, personnageCree *bool) {

	for {
		var choix int

		fmt.Println("====EXPLORATION DE LA NATURE====")
		fmt.Println("1.Création de personnage")
		fmt.Println("2.Explorer")
		fmt.Println("3.Retour")

		fmt.Scanln(&choix)
		switch choix {
		case 1:
			*perso = personnage.CreationPersonnage()
			*personnageCree = true
		case 2:
			if !*personnageCree {
				fmt.Println("vous devez créer un personnage d'abord")
				continue
			}
			menuexploration.Explorer(perso)
		case 3:
			return
		default:
			fmt.Println("choix invalide")
			return
		}
	}
}
