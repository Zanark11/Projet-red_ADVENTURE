package inventaire

import (
	"aventure/personnage"
	"fmt"
)

func AccessInventory(player *personnage.Character) bool {
	var choix int
	for {
		fmt.Println("====Inventaire====")
		fmt.Println("1.Afficher mon inventaire")
		fmt.Println("2.Retour")
		fmt.Println("Que veux tu faire ?")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Println("====inventaire====")
			fmt.Println("pièces:", player.Argent)
			compteur := 0
			for i, item := range player.Inventaire {
				fmt.Println(i+1, "-", item)
				if item != "" {
					compteur++
				}
			}
			fmt.Println("1.Jeter un objet")
			fmt.Println("2.Utiliser un objet")
			fmt.Println("3.Forgéron")
			fmt.Println("4.Retour")
			fmt.Println("Que veux tu faire ?")
			fmt.Scanln(&choix)

			switch choix {
			case 1:
				personnage.JeterObjet(player)
			case 3:
				personnage.FabriqueObjet(player)
			case 2:
                personnage.UtiliserObjet(player)
		    case 4:
			    return false
		    default:
			    fmt.Println("choix invalide")
			}
		case 2:
			return false
		default:
			fmt.Println("choix invalide")
		}
	}
}



	
