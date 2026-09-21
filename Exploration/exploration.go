package exploration

import (
	"aventure/personnage"
	"fmt"
)

func Exploration(perso *personnage.Character) bool {
	personnageCree := false
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
			personnageCree = true
		case 2:
			fmt.Println("l'exploration commence !")
		case 3:
			return personnageCree
		default:
			fmt.Println("choix invalide")
			return false
		}
	}
}
