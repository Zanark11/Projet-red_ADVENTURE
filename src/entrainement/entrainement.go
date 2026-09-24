package entrainement

import (
	"aventure/personnage"
	"fmt"
)
func Menuentrainement(perso *personnage.Character){

for{
    var choix int 

    fmt.Println("====Entrainement====")
    fmt.Println()
    fmt.Println("1.Combatre le Grandmage")
    fmt.Println("2.Retour")
    fmt.Scanln(&choix)

    switch choix {
    case 1:
	    Grandmagepattern(perso)
    case 2:
	    return
    default:
	    fmt.Println("choix invalide")
    }
  }	
}

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
		if tour%3 == 0 {
			degats = Grandmage.attaque * 2
		}
		personnage.Degats(perso, degats)

		fmt.Println(Grandmage.nom, "inflige", degats, "de dégât")

		personnage.Affichevie(perso)
		if !personnage.Playervivant(perso) {
			fmt.Println("ton personnage est mort")
			break
		}
		tour++
	}
}
