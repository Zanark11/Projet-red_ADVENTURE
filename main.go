package main

import (
	"aventure/entrainement"
	"aventure/Exploration"
	"aventure/inventaire"
	"aventure/menuMarchand"
	"aventure/menu_Principal"
	"aventure/personnage"
	"fmt"
)

func main() {
	var monPersonnage personnage.Character
	var personnageCree bool

	for {
		choix := menu_Principal.MenuPrincipal()

		switch choix {
		case 1:
			Exploration.Exploration(&monPersonnage, &personnageCree)
		case 2:
			if !personnageCree {
				fmt.Println("créer d'abord un personnage")
				continue
			}
			inventaire.AccessInventory(&monPersonnage)
		case 3:
			if !personnageCree {
				fmt.Println("créer d'abord un personnage")
				continue
			}
			menuMarchand.MenuMarchand(&monPersonnage)
		case 4:
			if !personnageCree {
				fmt.Println("créer un personnage d'abord")
				continue
			}
			entrainement.Menuentrainement(&monPersonnage)
		case 5:
			fmt.Println("QUITTER")
			return
		default:
			fmt.Println("chopix invalide")
		}
	}
}
