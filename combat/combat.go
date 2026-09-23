package combat

import (
	"fmt"
	"math/rand"

	"Projet-red_ADVENTURE/adversaire"
	"Projet-red_ADVENTURE/personnage"
)

// Question
type Question struct {
	Texte        string
	Reponses     [3]string
	BonneReponse int
}

// Pile ou face
func PileOuFace() string {
	if rand.Intn(2) == 0 {
		return "pile"
	}
	return "face"
}

// Réponse du Grand Sage
func ChoisirReponseGrandSage() int {
	return rand.Intn(3) + 1
}

// Vérifie si le joueur possède un objet
func PossedeObjet(joueur *personnage.Character, objet string) bool {
	for i := 0; i < len(joueur.Inventaire); i++ {
		if joueur.Inventaire[i] == objet {
			return true
		}
	}
	return false
}

// Menu pour utiliser un objet
func MenuObjetCombat(
	joueur *personnage.Character,
	gobelin *adversaire.Adversaire,
) bool {

	for {
		fmt.Println()
		fmt.Println("========== OBJETS ==========")
		fmt.Println("1 - Potion de vie")
		fmt.Println("2 - Livre de sort")
		fmt.Println("3 - Retour")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		// Potion
		if choix == 1 {
			if !PossedeObjet(joueur, "potion de vie") {
				fmt.Println("Vous n'avez pas de potion de vie.")
				continue
			}

			personnage.TakePot(joueur)
			return true
		}

		// Livre de sort
		if choix == 2 {
			for i := 0; i < len(joueur.Inventaire); i++ {
				if joueur.Inventaire[i] == "livre de sort" {
					joueur.Inventaire[i] = ""

					adversaire.PerdrePV(gobelin, 40)

					fmt.Println("Livre de sort utilisé !")
					fmt.Println("Le Grand Sage perd 40 PV.")

					return true
				}
			}

			fmt.Println("Vous n'avez pas de livre de sort.")
			continue
		}

		// Retour
		if choix == 3 {
			return false
		}

		fmt.Println("Choix invalide.")
	}
}

// ============================================================
// COMBAT 1
// ============================================================

var Serie1 = [9]Question{
	{
		"Quelle est la planète la plus proche du Soleil ?",
		[3]string{"Mars", "Mercure", "Venus"},
		2,
	},
	{
		"Combien y a-t-il de continents sur Terre ?",
		[3]string{"5", "6", "7"},
		3,
	},
	{
		"Quel animal est souvent appelé le roi de la jungle ?",
		[3]string{"Le lion", "Le tigre", "Le gorille"},
		1,
	},
	{
		"Quelle est la capitale de la France ?",
		[3]string{"Lyon", "Paris", "Marseille"},
		2,
	},
	{
		"Combien font 5 x 5 ?",
		[3]string{"10", "20", "25"},
		3,
	},
	{
		"Quelle couleur obtient-on en mélangeant du bleu et du jaune ?",
		[3]string{"Vert", "Orange", "Violet"},
		1,
	},
	{
		"Quel est le plus grand océan du monde ?",
		[3]string{"Atlantique", "Pacifique", "Indien"},
		2,
	},
	{
		"Combien de jours compte une semaine ?",
		[3]string{"5", "7", "10"},
		2,
	},
	{
		"Quel langage utilisons-nous pour développer Red Adventure ?",
		[3]string{"Python", "Java", "Go"},
		3,
	},
}

// Combat 1
func JouerSerie1(joueur *personnage.Character) {

	gobelin := adversaire.CreerAdversaire()

	debut := 0
	tourJoueur := PileOuFace() == "pile"

	fmt.Println()
	fmt.Println("========== COMBAT 1 ==========")

	for !personnage.EstVaincu(*joueur) &&
		!adversaire.EstVaincu(gobelin) {

		// Option toujours disponible avant l'action
		fmt.Println()
		fmt.Println("========== ACTION ==========")
		fmt.Println("1 - Continuer le combat")
		fmt.Println("2 - Utiliser un objet")
		fmt.Print("Votre choix : ")

		var action int
		fmt.Scanln(&action)

		// Utiliser un objet
		if action == 2 {

			utilise := MenuObjetCombat(joueur, &gobelin)

			if utilise {
				personnage.AfficherPV(*joueur)
				adversaire.AfficherPV(gobelin)

				// L'objet utilise le tour
				tourJoueur = !tourJoueur
			}

			continue
		}

		if action != 1 {
			fmt.Println("Choix invalide.")
			continue
		}

		// ========================================================
		// TOUR DU JOUEUR
		// ========================================================

		if tourJoueur {

			fmt.Println()
			fmt.Println("========== TOUR DU JOUEUR ==========")
			fmt.Println("Choisissez une question :")

			for j := 0; j < 3; j++ {
				fmt.Println(j+1, "-", Serie1[debut+j].Texte)
			}

			fmt.Print("Votre choix : ")

			var choix int
			fmt.Scanln(&choix)

			if choix < 1 || choix > 3 {
				fmt.Println("Choix invalide.")
				continue
			}

			question := Serie1[debut+choix-1]

			fmt.Println()
			fmt.Println("Question :", question.Texte)

			for j := 0; j < 3; j++ {
				fmt.Println(j+1, "-", question.Reponses[j])
			}

			reponse := ChoisirReponseGrandSage()

			fmt.Println()
			fmt.Println("Le Grand Sage choisit :", reponse)

			if reponse == question.BonneReponse {
				fmt.Println("Bonne réponse !")
				fmt.Println("Le joueur perd 40 PV.")

				personnage.PerdrePV(joueur, 40)
			} else {
				fmt.Println("Mauvaise réponse !")
				fmt.Println("Le Grand Sage perd 40 PV.")

				adversaire.PerdrePV(&gobelin, 40)
			}

			personnage.AfficherPV(*joueur)
			adversaire.AfficherPV(gobelin)

			debut += 3

			if debut >= len(Serie1) {
				debut = 0
			}

			tourJoueur = false

		} else {

			// ====================================================
			// TOUR DU GRAND SAGE
			// ====================================================

			fmt.Println()
			fmt.Println("========== TOUR DU GRAND SAGE ==========")

			question := Serie1[debut+rand.Intn(3)]

			fmt.Println("Question :", question.Texte)

			for j := 0; j < 3; j++ {
				fmt.Println(j+1, "-", question.Reponses[j])
			}

			fmt.Print("Votre réponse : ")

			var reponse int
			fmt.Scanln(&reponse)

			if reponse == question.BonneReponse {
				fmt.Println("Bonne réponse !")
				fmt.Println("Le Grand Sage perd 40 PV.")

				adversaire.PerdrePV(&gobelin, 40)
			} else {
				fmt.Println("Mauvaise réponse !")
				fmt.Println("Le joueur perd 40 PV.")

				personnage.PerdrePV(joueur, 40)
			}

			personnage.AfficherPV(*joueur)
			adversaire.AfficherPV(gobelin)

			debut += 3

			if debut >= len(Serie1) {
				debut = 0
			}

			tourJoueur = true
		}
	}

	// ========================================================
	// FIN COMBAT 1
	// ========================================================

	if personnage.EstVaincu(*joueur) {

		fmt.Println()
		fmt.Println("========== DEFAITE ==========")
		fmt.Println("Vous avez perdu le combat 1.")

	} else {

		fmt.Println()
		fmt.Println("========== VICTOIRE ==========")
		fmt.Println("Vous avez vaincu le Grand Sage !")

		personnage.GagnerCombat(joueur)

		fmt.Println()
		fmt.Println("========== COMBAT 2 ==========")

		// Passage au combat 2
		JouerSerie2(joueur)
	}
}

// ============================================================
// COMBAT 2
// ============================================================

var Serie2 = [9]Question{
	{
		"Quelle est la plus grande planète du système solaire ?",
		[3]string{"Mars", "Jupiter", "Venus"},
		2,
	},
	{
		"Combien font 10 + 5 ?",
		[3]string{"15", "20", "25"},
		1,
	},
	{
		"Quel langage est utilisé pour Red Adventure ?",
		[3]string{"Go", "Java", "Python"},
		1,
	},
	{
		"Combien de côtés possède un triangle ?",
		[3]string{"3", "4", "5"},
		1,
	},
	{
		"Quelle est la capitale de l'Italie ?",
		[3]string{"Madrid", "Rome", "Berlin"},
		2,
	},
	{
		"Combien font 6 x 6 ?",
		[3]string{"30", "36", "42"},
		2,
	},
	{
		"Quel est le symbole chimique de l'eau ?",
		[3]string{"CO2", "H2O", "O2"},
		2,
	},
	{
		"Combien y a-t-il d'heures dans une journée ?",
		[3]string{"12", "24", "48"},
		2,
	},
	{
		"Quel système d'exploitation est développé par Microsoft ?",
		[3]string{"Windows", "Linux", "Android"},
		1,
	},
}

// Combat 2
func JouerSerie2(joueur *personnage.Character) {

	gobelin := adversaire.CreerAdversaire()

	debut := 0
	tourJoueur := PileOuFace() == "pile"

	fmt.Println()
	fmt.Println("========== COMBAT 2 ==========")

	for !personnage.EstVaincu(*joueur) &&
		!adversaire.EstVaincu(gobelin) {

		// Option objet
		fmt.Println()
		fmt.Println("========== ACTION ==========")
		fmt.Println("1 - Continuer le combat")
		fmt.Println("2 - Utiliser un objet")
		fmt.Print("Votre choix : ")

		var action int
		fmt.Scanln(&action)

		if action == 2 {

			utilise := MenuObjetCombat(joueur, &gobelin)

			if utilise {
				personnage.AfficherPV(*joueur)
				adversaire.AfficherPV(gobelin)

				tourJoueur = !tourJoueur
			}

			continue
		}

		if action != 1 {
			fmt.Println("Choix invalide.")
			continue
		}

		// TOUR DU JOUEUR
		if tourJoueur {

			fmt.Println()
			fmt.Println("========== TOUR DU JOUEUR ==========")
			fmt.Println("Choisissez une question :")

			for j := 0; j < 3; j++ {
				fmt.Println(j+1, "-", Serie2[debut+j].Texte)
			}

			fmt.Print("Votre choix : ")

			var choix int
			fmt.Scanln(&choix)

			if choix < 1 || choix > 3 {
				fmt.Println("Choix invalide.")
				continue
			}

			question := Serie2[debut+choix-1]

			fmt.Println()
			fmt.Println("Question :", question.Texte)

			for j := 0; j < 3; j++ {
				fmt.Println(j+1, "-", question.Reponses[j])
			}

			reponse := ChoisirReponseGrandSage()

			if reponse == question.BonneReponse {
				personnage.PerdrePV(joueur, 40)
			} else {
				adversaire.PerdrePV(&gobelin, 40)
			}

			personnage.AfficherPV(*joueur)
			adversaire.AfficherPV(gobelin)

			debut += 3

			if debut >= len(Serie2) {
				debut = 0
			}

			tourJoueur = false

		} else {

			// TOUR DU GRAND SAGE
			fmt.Println()
			fmt.Println("========== TOUR DU GRAND SAGE ==========")

			question := Serie2[debut+rand.Intn(3)]

			fmt.Println("Question :", question.Texte)

			for j := 0; j < 3; j++ {
				fmt.Println(j+1, "-", question.Reponses[j])
			}

			fmt.Print("Votre réponse : ")

			var reponse int
			fmt.Scanln(&reponse)

			if reponse == question.BonneReponse {
				adversaire.PerdrePV(&gobelin, 40)
			} else {
				personnage.PerdrePV(joueur, 40)
			}

			personnage.AfficherPV(*joueur)
			adversaire.AfficherPV(gobelin)

			debut += 3

			if debut >= len(Serie2) {
				debut = 0
			}

			tourJoueur = true
		}
	}

	// FIN COMBAT 2
	if personnage.EstVaincu(*joueur) {

		fmt.Println()
		fmt.Println("========== DEFAITE ==========")
		fmt.Println("Vous avez perdu le combat 2.")

	} else {

		fmt.Println()
		fmt.Println("========== VICTOIRE ==========")
		fmt.Println("Vous avez gagné le combat 3 !")

		personnage.GagnerCombat(joueur)
	}
}