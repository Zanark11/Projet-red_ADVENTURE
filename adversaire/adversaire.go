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
		nom:              "Grand sage",
		niveau:           1,
		pointDeVieActuel: 100,
		pointDeVieMax:    100,
	}
	return adversaire
}

// PerdrePV retire des points de vie à l'adversaire.
func PerdrePV(adversaire *Adversaire, degats int) {
	// On retire les dégâts aux points de vie actuels.
	adversaire.pointDeVieActuel -= degats

	// Les points de vie ne peuvent pas être négatifs.
	if adversaire.pointDeVieActuel < 0 {
		adversaire.pointDeVieActuel = 0
	}
}

// AfficherPV affiche les points de vie actuels de l'adversaire.
func AfficherPV(adversaire Adversaire) {
	fmt.Println("PV de", adversaire.nom, ":", adversaire.pointDeVieActuel, "/", adversaire.pointDeVieMax)
}

// EstVaincu vérifie si l'adversaire n'a plus de PV.
func EstVaincu(adversaire Adversaire) bool {
	return adversaire.pointDeVieActuel <= 0
}
