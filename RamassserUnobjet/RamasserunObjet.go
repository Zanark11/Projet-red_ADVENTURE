package ramassserunobjet

import (
	marchand "aventure/Marchand"
	"aventure/personnage"
	"fmt"
)

/*code pour l'appeler dans explorer pour que le joueur ait la possibilitée de ramaser comme de réfuser*/
func RamasserUnobjet(player *personnage.Character, objet string) {
	var choix int
	for i, objetInventaire := range player.Inventaire {
		if objetInventaire == "" {
			player.Inventaire[i] = objet
			return
		}
	}
	fmt.Println("votre invenraire est plein")
	fmt.Println("voulez vous augmenter votre inventaire ?")
	fmt.Println()
	fmt.Println("1.Oui")
	fmt.Println("2.Non")

	fmt.Scanln(&choix)
	switch choix {
	case 1:
		if marchand.AugmenterInventaire(player) {
			fmt.Println("votre inventaire a augmenté !")
			for i, objetInventaire := range player.Inventaire {
				if objetInventaire == "" {
					player.Inventaire[i] = objet
					break
				}
			}

		} else {
			fmt.Println("vous n'avez pas assez d'argent")
		}
	case 2:
		fmt.Println("vous laissez l'objet !")
		return
	default:
		fmt.Println("choix invalide")
	}
}
