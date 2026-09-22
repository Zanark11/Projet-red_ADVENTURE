package entrainement

import (
	"aventure/personnage"
	"fmt"
)


type adversaire struct {
	nom              string
	attaque          int
	pointDeVieMax    int
	pointDeVieActuel int
}

func initGrandmage() adversaire {
	Grandmage := adversaire{
		nom:              "Grandmage d'entrainement",
		attaque:          20,
		pointDeVieActuel: 100,
		pointDeVieMax:    100,
	}
	return Grandmage
}
func Grandmagepattern(perso *personnage.Character) {
	
	Grandmage := initGrandmage()
	tour := 1
	for {
		degats := Grandmage.attaque
		if tour %3==0 {
			degats = Grandmage.attaque * 2
		}
		personnage.Degats(perso, degats)
		fmt.Println(Grandmage.nom,  "inflige", degats, "de dégât")
		tour++
	}
}
