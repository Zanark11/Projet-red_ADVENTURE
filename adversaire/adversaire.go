package adversaire

import "fmt"

type Adversaire struct {
	nom              string
	pointDeVieActuel int
	niveau           int
	pointDeVieMax    int
}

func CreerAdversaire() Adversaire {
	adversaire := Adversaire{
		nom:              "Grandmage",
		niveau:           1,
		pointDeVieActuel: 115,
		pointDeVieMax:    200,
	}
	return adversaire
}
func Degats(ennemi *Adversaire, degats int) {
	ennemi.pointDeVieActuel -= degats
	if ennemi.pointDeVieActuel < 0 {
		ennemi.pointDeVieActuel = 0
	}
}

func Adversairevivant(ennemi Adversaire) bool {
	return ennemi.pointDeVieActuel > 0
}

func Affichevie(ennemi Adversaire) {
	fmt.Println("vie du", ennemi.nom, ":", ennemi.pointDeVieActuel, "/", ennemi.pointDeVieMax)
}

