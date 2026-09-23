package combat

import (
	"fmt"
	"math/rand"

	"Projet-red_ADVENTURE/adversaire"
	"Projet-red_ADVENTURE/personnage"
)

type Question struct {
	Texte        string
	Reponses     [3]string
	BonneReponse int
}

func PileOuFace() string {
	if rand.Intn(2) == 0 {
		return "pile"
	}
	return "face"
}

func ChoisirReponseGrandSage() int {
	return rand.Intn(3) + 1
}

// =====================
// QUESTIONS COMBAT 1
// =====================

var Serie1 = [9]Question{
	{"Quelle est la planète la plus proche du Soleil ?", [3]string{"Mars", "Mercure", "Venus"}, 2},
	{"Combien y a-t-il de continents sur Terre ?", [3]string{"5", "6", "7"}, 3},
	{"Quel animal est souvent appelé le roi de la jungle ?", [3]string{"Le lion", "Le tigre", "Le gorille"}, 1},
	{"Quelle est la capitale de la France ?", [3]string{"Lyon", "Paris", "Marseille"}, 2},
	{"Combien font 5 x 5 ?", [3]string{"10", "20", "25"}, 3},
	{"Quelle couleur obtient-on en mélangeant du bleu et du jaune ?", [3]string{"Vert", "Orange", "Violet"}, 1},
	{"Quel est le plus grand océan du monde ?", [3]string{"Océan Atlantique", "Océan Pacifique", "Océan Indien"}, 2},
	{"Combien de jours compte une semaine ?", [3]string{"5", "7", "10"}, 2},
	{"Quel langage utilisons-nous pour développer Red Adventure ?", [3]string{"Python", "Java", "Go"}, 3},
}

// =====================
// COMBAT 1
// =====================

func JouerSerie1(joueur *personnage.Character) {

	gobelin := adversaire.CreerAdversaire()
	debut := 0
	tourJoueur := PileOuFace() == "pile"

	fmt.Println("========== COMBAT 1 ==========")

	for !personnage.EstVaincu(*joueur) && !adversaire.EstVaincu(gobelin) {

		if tourJoueur {

			fmt.Println()
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

			fmt.Println("Question :", question.Texte)

			for j := 0; j < 3; j++ {
				fmt.Println(j+1, "-", question.Reponses[j])
			}

			reponse := ChoisirReponseGrandSage()

			if reponse == question.BonneReponse {
				fmt.Println("Bonne réponse !")
				personnage.PerdrePV(joueur, 40)
			} else {
				fmt.Println("Mauvaise réponse !")
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
				adversaire.PerdrePV(&gobelin, 40)
			} else {
				fmt.Println("Mauvaise réponse !")
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

	if personnage.EstVaincu(*joueur) {
		fmt.Println()
		fmt.Println("========== DÉFAITE ==========")
	} else {
		fmt.Println()
		fmt.Println("========== VICTOIRE ==========")
		fmt.Println("Vous avez vaincu le Grand Sage !")

		// Récompenses
		personnage.GagnerCombat(joueur)

		// Passage au combat 2
		fmt.Println()
		fmt.Println("========== COMBAT 2 ==========")

		JouerSerie2(joueur)
	}
}

// =====================
// QUESTIONS COMBAT 2
// =====================

var Serie2 = [9]Question{
	{"Quelle est la plus grande planète du système solaire ?", [3]string{"Mars", "Jupiter", "Venus"}, 2},
	{"Combien font 10 + 5 ?", [3]string{"15", "20", "25"}, 1},
	{"Quel langage est utilisé dans Red Adventure ?", [3]string{"Go", "Java", "Python"}, 1},
	{"Combien y a-t-il de côtés sur un triangle ?", [3]string{"3", "4", "5"}, 1},
	{"Quelle est la capitale de l'Italie ?", [3]string{"Madrid", "Rome", "Berlin"}, 2},
	{"Combien font 6 x 6 ?", [3]string{"30", "36", "42"}, 2},
	{"Quel est le symbole chimique de l'eau ?", [3]string{"CO2", "H2O", "O2"}, 2},
	{"Combien y a-t-il d'heures dans une journée ?", [3]string{"12", "24", "48"}, 2},
	{"Quel système d'exploitation est développé par Microsoft ?", [3]string{"Windows", "Linux", "Android"}, 1},
}

// =====================
// COMBAT 2
// =====================

func JouerSerie2(joueur *personnage.Character) {

	gobelin := adversaire.CreerAdversaire()
	debut := 0
	tourJoueur := PileOuFace() == "pile"

	fmt.Println()
	fmt.Println("Le combat 2 commence !")

	for !personnage.EstVaincu(*joueur) && !adversaire.EstVaincu(gobelin) {

		if tourJoueur {

			fmt.Println()
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

	if personnage.EstVaincu(*joueur) {
		fmt.Println()
		fmt.Println("========== DÉFAITE ==========")
	} else {
		fmt.Println()
		fmt.Println("========== VICTOIRE ==========")
		fmt.Println("Vous avez gagné le combat 1 !")
        fmt.Println("Vous avez vaincu le Grand Sage !")
		fmt.Println("vous passez au combat 2 !")
		// Récompenses du combat 2
		personnage.GagnerCombat(joueur)
	}
}