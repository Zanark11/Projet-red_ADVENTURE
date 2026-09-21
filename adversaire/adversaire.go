package adversaire


type Adversaire struct {
	nom                   string
	pointDeVieActuel      int
	niveau                int
	pointDeVieMax         int
}

func CreerAdversaire() Adversaire{
	adversaire := Adversaire{
    nom:"Gobelin",
	niveau:1,
	pointDeVieActuel:100,
	pointDeVieMax : 100,
	}
	return adversaire
}