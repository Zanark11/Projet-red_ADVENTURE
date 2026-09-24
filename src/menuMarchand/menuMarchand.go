package menuMarchand

import (
	"aventure/marchand"
	"aventure/personnage" 
	"fmt"
)

func MenuMarchand(player *personnage.Character) {
for{
	var choix int

	fmt.Println("====Marchand====")
	fmt.Println("1.voir les objets disponibles")
	fmt.Println("2.Augmentation de l'inventaire")
	fmt.Println("3.Retour")

	fmt.Scanln(&choix)

	switch choix {
	case 1:
		marchand.Afficherobjet(player)
	case 2:
	    marchand.AmeliorationInventaire(player) 
	case 3:
		return
	default:
		fmt.Println("choix invalide")
	}
  }
}
