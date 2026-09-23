package personnage

import "fmt"

func CreationPersonnage() Character {
	var nom string
	var choixclasse string
	for {
		valide := true
		fmt.Println("======*CREATION DU PERSONNAGE*======")
		fmt.Println("Donne un nom à ton perso ? ")
		fmt.Scanln(&nom)

		if nom == "" {
			fmt.Println("ne nom ne pas peut pas être vide")
			continue
		}
		premierelettre := nom[0]

		for _, lettre := range nom {
			if lettre >= '0' && lettre <= '9' {
				valide = false
			}
		}
		if !valide {
			fmt.Println("votre ne peut contenir de chiffre")
			continue
		}
		if premierelettre >= 'a' && premierelettre <= 'z' {
			premierelettre = premierelettre - ('a' - 'A')
			nom = string(premierelettre) + nom[1:]
		}
		break
	}
	fmt.Println("choisir ta classe:")
	fmt.Println("1. Avocat")
	fmt.Println("2. Informaticien")

	fmt.Println("3. Medecin")

	fmt.Scanln(&choixclasse)

	character := InitCharacter(nom, choixclasse)
	DisplayInfo(character)

	return character
}

func DisplayInfo(character Character) {
	fmt.Println()
	fmt.Println("félicitation personnage crée !")
	fmt.Println("nom :", character.nom)
	fmt.Println("pointDeVieActuel:", character.pointDeVieActuel)
	fmt.Println("classe:", character.classe)
	fmt.Println("niveau:", character.niveau)
	fmt.Println("inventaire:", character.Inventaire)
	fmt.Println("xp:", character.xp)
	fmt.Println("mana:", character.mana)
	fmt.Println("Argent", character.Argent)
}

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
		Inventaire:       []string{"potion de vie", "potion de vie", "potion de vie", "", "", ""},
		xp:               0,
		mana:             100,
		Argent:           100,
	}
	return character
}

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

func JeterObjet(character *Character) {
	var choix int

	fmt.Println("Quel objet aimerais tu jeter ?")
	fmt.Scanln(&choix)
	if choix < 1 || choix > len(character.Inventaire) {
		fmt.Println("choix invalide")
		return
	}
	index := choix - 1
	if character.Inventaire[index] == "" {
		fmt.Println("case déja vide .")
		return
	}
	objet := character.Inventaire[index]

	character.Inventaire[index] = ""

	fmt.Println("vous avez jetez:", objet)
}

// PerdrePV retire des points de vie au personnage.
func PerdrePV(character *Character, degats int) {
	// On retire les dégâts aux PV actuels.
	character.pointDeVieActuel -= degats

	// Les PV ne peuvent pas descendre sous 0.
	if character.pointDeVieActuel < 0 {
		character.pointDeVieActuel = 0
	}
}

// AfficherPV affiche les points de vie du personnage.
func AfficherPV(character Character) {
	fmt.Println("PV du joueur :", character.pointDeVieActuel, "/", character.pointsDeVieMax)
}

// EstVaincu vérifie si le personnage n'a plus de PV.
func EstVaincu(character Character) bool {
	// Si les PV sont à 0 ou moins,
	// le personnage est vaincu.
	return character.pointDeVieActuel <= 0
}

