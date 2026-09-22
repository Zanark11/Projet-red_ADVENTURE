package adversaire

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
		pointDeVieActuel: 100,
		pointDeVieMax:    200,
	}
	return adversaire
}
