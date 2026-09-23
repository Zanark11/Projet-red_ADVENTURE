package personnage

import "fmt"

// Character représente le personnage du joueur.
type Character struct {
	nom              string
	classe           string
	niveau           int
	pointsDeVieMax   int
	pointDeVieActuel int
	Inventaire       []string
	xp               int
	mana             int
	Argent           int
}

// CreationPersonnage crée un nouveau personnage.
func CreationPersonnage() Character {
	var nom string
	var choixclasse string

	for {
		valide := true

		fmt.Println("====== CREATION DU PERSONNAGE ======")
		fmt.Println("Donne un nom à ton perso ?")
		fmt.Scanln(&nom)

		if nom == "" {
			fmt.Println("Le nom ne peut pas être vide.")
			continue
		}

		premierelettre := nom[0]

		for _, lettre := range nom {
			if lettre >= '0' && lettre <= '9' {
				valide = false
			}
		}

		if !valide {
			fmt.Println("Votre nom ne peut pas contenir de chiffre.")
			continue
		}

		if premierelettre >= 'a' && premierelettre <= 'z' {
			premierelettre = premierelettre - ('a' - 'A')
			nom = string(premierelettre) + nom[1:]
		}

		break
	}

	fmt.Println("Choisir ta classe :")
	fmt.Println("1. Avocat")
	fmt.Println("2. Informaticien")
	fmt.Println("3. Medecin")

	fmt.Scanln(&choixclasse)

	character := InitCharacter(nom, choixclasse)

	DisplayInfo(character)

	return character
}

// DisplayInfo affiche les informations du personnage.
func DisplayInfo(character Character) {
	fmt.Println()
	fmt.Println("Félicitations, personnage créé !")
	fmt.Println("Nom :", character.nom)
	fmt.Println("Points de vie :", character.pointDeVieActuel)
	fmt.Println("Classe :", character.classe)
	fmt.Println("Niveau :", character.niveau)
	fmt.Println("Inventaire :", character.Inventaire)
	fmt.Println("XP :", character.xp)
	fmt.Println("Mana :", character.mana)
	fmt.Println("Argent :", character.Argent)
}

// InitCharacter initialise un personnage.
func InitCharacter(nom string, choixclasse string) Character {
	pointDeVieActuel := 100

	switch choixclasse {
	case "1":
		pointDeVieActuel = 90
	case "2":
		pointDeVieActuel = 60
	case "3":
		pointDeVieActuel = 70
	}

	character := Character{
		nom:              nom,
		pointDeVieActuel: pointDeVieActuel,
		classe:           choixclasse,
		niveau:           1,
		pointsDeVieMax:   100,
		Inventaire: []string{
			"potion de vie",
			"potion de vie",
			"potion de vie",
			"",
			"",
			"",
		},
		xp:     0,
		mana:   100,
		Argent: 100,
	}

	return character
}

// JeterObjet permet de jeter un objet de l'inventaire.
func JeterObjet(character Character) {
	var choix int

	fmt.Println("Quel objet aimerais-tu jeter ?")
	fmt.Scanln(&choix)

	if choix < 1 || choix > len(character.Inventaire) {
		fmt.Println("Choix invalide.")
		return
	}

	index := choix - 1

	if character.Inventaire[index] == "" {
		fmt.Println("Case déjà vide.")
		return
	}

	objet := character.Inventaire[index]
	character.Inventaire[index] = ""

	fmt.Println("Vous avez jeté :", objet)
}

// PerdrePV retire des points de vie au personnage.
func PerdrePV(character *Character, degats int) {
	character.pointDeVieActuel -= degats

	if character.pointDeVieActuel < 0 {
		character.pointDeVieActuel = 0
	}
}

// AfficherPV affiche les points de vie du personnage.
func AfficherPV(character Character) {
	fmt.Println(
		"PV du joueur :",
		character.pointDeVieActuel,
		"/",
		character.pointsDeVieMax,
	)
}

// EstVaincu vérifie si le personnage n'a plus de PV.
func EstVaincu(character Character) bool {
	return character.pointDeVieActuel <= 0
}

// GagnerCombat donne les récompenses après une victoire.
func GagnerCombat(character *Character) {
	// Le joueur gagne 30 pièces.
	character.Argent += 30

	// Le joueur gagne 1 XP et 1 niveau,
	// jusqu'au niveau maximum 10.
	if character.niveau < 10 {
		character.xp++
		character.niveau++

		fmt.Println("+1 XP !")
		fmt.Println("Nouveau niveau :", character.niveau)
	} else {
		fmt.Println("Niveau maximum atteint !")
	}

	fmt.Println("+30 pièces !")
	fmt.Println("XP :", character.xp)
	fmt.Println("Argent :", character.Argent, "pièces")
}

// TempsReponse calcule le temps disponible pour répondre.
func TempsReponse(character Character) int {
	// Niveau 1 = 20 secondes
	// Chaque niveau ajoute 5 secondes.
	//
	// Niveau 1 = 20 secondes
	// Niveau 2 = 25 secondes
	// Niveau 3 = 30 secondes
	// ...
	// Niveau 10 = 65 secondes.

	temps := 20 + (character.niveau-1)*5

	return temps
}