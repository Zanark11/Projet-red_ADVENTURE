package main

import (
	"aventure/exploration"
	"aventure/inventaire"
	"aventure/menu_Principal"
	"aventure/personnage"
	"fmt"
)

func main() {
	var monPersonnage personnage.Character
	var personnagecree bool

	for {
		choix := menu_Principal.MenuPrincipal()

		switch choix {
		case 1:
			personnagecree = exploration.Exploration(&monPersonnage)
		case 2:
			if !personnagecree {
				fmt.Println("créer d'abord un personnage")
				continue
			}
			inventaire.AccessInventory(monPersonnage)
		case 3:
			fmt.Println("ENTRAINEMENT")
		case 4:
			fmt.Println("QUITTER")
			return

		}
	}
}
