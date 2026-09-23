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
		reductionTemps:   0,
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
	reductionTemps   int
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

// GagnerCombat donne les récompenses après avoir gagné un combat.
func GagnerCombat(character *Character) {
	// Le joueur gagne 30 pièces.
	character.Argent += 30

	// Le joueur gagne 1 XP et monte d'un niveau,
	// sauf s'il est déjà au niveau maximum.
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

// TempsReponse calcule le temps disponible pour répondre
// en fonction du niveau du joueur.
func TempsReponse(character Character) int {
	// Niveau 1 = 20 secondes.
	// Chaque niveau ajoute 5 secondes.
	temps := 20 + (character.niveau-1)*5

	// Chronoboros peut réduire le temps de réponse.
	temps -= character.reductionTemps

	// Le niveau maximum est 10,
	// donc le temps maximum est 65 secondes.
	if character.niveau > 10 {
		temps = 65
	}

	return temps
}

// GetMana retourne la quantité de mana actuelle du personnage.
func GetMana(character Character) int {
	return character.mana
}

// DepenserMana retire une quantité de mana au personnage.
// La fonction retourne true si le joueur avait assez de mana.
func DepenserMana(character *Character, cout int) bool {
	if character.mana < cout {
		return false
	}

	character.mana -= cout
	return true
}

// ReductionTemps réduit le temps de réponse du personnage.
func (character *Character) ReductionTemps(reduction int) {
	character.reductionTemps += reduction
}
