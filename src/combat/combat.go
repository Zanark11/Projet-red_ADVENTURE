package combat

import (
	"fmt"
	"math/rand"

	"aventure/adversaire"
	"aventure/personnage"
)

type Question struct {
	Texte        string
	Reponses     [3]string
	BonneReponse int
}

func PileOuFace() string {
	var choix int

	for {
		fmt.Println("Choisissez")
		fmt.Println("1.Pile")
		fmt.Println("2.Face")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			return "pile"
		case 2:
			return "face"
		default:
			fmt.Println("choix invalide")
		}
	}
}

func ChoisirReponseGrandSage() int {

	return rand.Intn(3) + 1
}

var Serie1 = []Question{
	{
		Texte: "Je peux être cassé sans être touchée. Qui suis je ?",
		Reponses: [3]string{
			"une promesse",
			"une pierre",
			"une porte",
		},
		BonneReponse: 1,
	},
	{
		Texte: "Plus je sèche, plus je suis mouillée. Qui suis je ?",
		Reponses: [3]string{
			"une éponge",
			"une serviette",
			"une éponge de mer",
		},
		BonneReponse: 2,
	},
	{
		Texte: "Quel animal est souvent appelé le roi de la jungle ?",
		Reponses: [3]string{
			"Le lion",
			"Le tigre",
			"Le gorille",
		},
		BonneReponse: 1,
	},
	{
		Texte: "j'ai des aiguilles mais je ne peux pas coudre. Qui suis je ?",
		Reponses: [3]string{
			"une horloge",
			"un hérisson",
			"un sapin",
		},
		BonneReponse: 1,
	},
	{
		Texte: "Combien font 5 x 5 ?",
		Reponses: [3]string{
			"10",
			"20",
			"25",
		},
		BonneReponse: 3,
	},
	{
		Texte: "Quelle couleur obtient-on en mélangeant du bleu et du jaune ?",
		Reponses: [3]string{
			"Vert",
			"Orange",
			"Violet",
		},
		BonneReponse: 1,
	},
	{
		Texte: "Quel est le plus grand océan du monde ?",
		Reponses: [3]string{
			"Océan Atlantique",
			"Océan Pacifique",
			"Océan Indien",
		},
		BonneReponse: 2,
	},
	{
		Texte: "je grandis sans être vivant et je meurs sans avoir vécu. Qui suis je ?",
		Reponses: [3]string{
			"une plante",
			"un feu",
			"un nuage",
		},
		BonneReponse: 2,
	},
	{
		Texte: "plus on m'enlève, plus je deviens grand. Qui suis je ?",
		Reponses: [3]string{
			"un trou",
			"un arbre",
			"une montagne",
		},
		BonneReponse: 1,
	},
	{
		Texte: "Qu'est ce qui a des dents mais ne peux pas manger ?",
		Reponses: [3]string{
			"un crocodile",
			"un peigne",
			"une fourchette",
		},
		BonneReponse: 2,
	},
	{
		Texte: "Qu'est ce qui monte et descend sans jamais bouger ?",
		Reponses: [3]string{
			"un ascenseur",
			"un escalier",
			"un balon",
		},
		BonneReponse: 2,
	},
	{
		Texte: "Qu'est ce qui a un cou mais pas de tête ?",
		Reponses: [3]string{
			"une bouteille",
			"une girafe",
			"une chemise",
		},
		BonneReponse: 1,
	},
	{
		Texte: "Je n'ai pas de jambes, mais je peux courir. je n'ai pas de bouche,mais je peux murmurer. Qui suis je ?",
		Reponses: [3]string{
			"le vent",
			"une rivière",
			"un nuage",
		},
		BonneReponse: 2,
	},
}

func JouerSerie1(joueur *personnage.Character) {

	gobelin := adversaire.CreerAdversaire()

	resultat := PileOuFace()

	debut := 0

	tourJoueur := resultat == "pile"

	fmt.Println("Pile ou face...")

	if resultat == "pile" {
		fmt.Println("Résultat : pile")
		fmt.Println("Le joueur commence !")
	} else {
		fmt.Println("Résultat : face")
		fmt.Println("Le Grandmage commence !")
	}

	fmt.Println()

	fmt.Println("========== COMBAT - TOUR PAR TOUR ==========")
	fmt.Println()

	for personnage.Playervivant(joueur) && adversaire.Adversairevivant(gobelin) {

		if tourJoueur {
			var action int
			fmt.Println("Choisissez une question :")
			fmt.Println()
			fmt.Println("1.Choisir une question")
			fmt.Println("2.Utiliser un objet")
			fmt.Println("Quelle action choissisez vous ?")
			fmt.Scanln(&action)

			if action == 2 {
				personnage.UtiliserObjet(joueur, &gobelin)
				tourJoueur = false
				continue
			}
			if action != 1 {
				fmt.Println("choix invalide")
				continue
			}
			fmt.Println()
			fmt.Println("choissisez une question :")
			fmt.Println()

			for j := 0; j < 3 && debut+j <len(Serie1); j++ {
				fmt.Println(j+1, "-", Serie1[debut+j].Texte)
			}

			fmt.Print("\nVotre choix : ")

			var choix int
			fmt.Scanln(&choix)

			if choix < 1 || choix > 3 {
				fmt.Println("Choix invalide.")
				continue
			}

			question := Serie1[debut+choix-1]

			fmt.Println()
			fmt.Println("Question :", question.Texte)
			fmt.Println()

			for j := 0; j < len(question.Reponses); j++ {
				fmt.Println(j+1, "-", question.Reponses[j])
			}

			reponse := ChoisirReponseGrandSage()

			fmt.Println("Le Grandmage choisit la réponse :", reponse)

			fmt.Println()

			if reponse == question.BonneReponse {
				fmt.Println("Bonne réponse !")
				fmt.Println("Le joueur perd 40 PV.")

				personnage.Degats(joueur, 40)
			} else {
				fmt.Println("Mauvaise réponse !")
				fmt.Println("Le Grandmage perd 40 PV.")

				adversaire.Degats(&gobelin, 40)
			}

			personnage.Affichevie(joueur)
			adversaire.Affichevie(gobelin)

			fmt.Println()

			debut += 3

			if debut >= len(Serie1) {
				debut = 0
			}

			tourJoueur = false

		} else {

			fmt.Println("========== TOUR DU GRANDMAGE ==========")
			fmt.Println()

			indiceQuestion := rand.Intn(3)
			if debut+indiceQuestion>=len(Serie1){
				indiceQuestion=0
			}
			question := Serie1[debut+indiceQuestion]

			fmt.Println("Le Grandmage choisit une question :")
			fmt.Println()
			fmt.Println("Question :", question.Texte)
			fmt.Println()

			for j := 0; j < len(question.Reponses); j++ {
				fmt.Println(j+1, "-", question.Reponses[j])
			}

			fmt.Println()
			fmt.Print("Votre réponse : ")

			var reponseJoueur int
			fmt.Scanln(&reponseJoueur)

			if reponseJoueur == question.BonneReponse {
				fmt.Println("Bonne réponse !")
				fmt.Println("Le Grandmage perd 40 PV.")

				adversaire.Degats(&gobelin, 40)
			} else {
				fmt.Println("Mauvaise réponse !")
				fmt.Println("Le joueur perd 40 PV.")

				personnage.Degats(joueur, 40)
			}

			fmt.Println()
			personnage.Affichevie(joueur)
			adversaire.Affichevie(gobelin)
			fmt.Println()

			debut += 3

			if debut >= len(Serie1) {
				debut = 0
			}

			tourJoueur = true
		}
	}
	if !personnage.Playervivant(joueur) {

		fmt.Println("================================")
		fmt.Println("DÉFAITE !")
		fmt.Println("Le Grandmage a gagné le combat.")
		fmt.Println("================================")

	} else if !adversaire.Adversairevivant(gobelin) {

		fmt.Println("================================")
		fmt.Println("VICTOIRE !")
		fmt.Println("Vous avez vaincu le Grandmage !")
		fmt.Println("================================")

		fmt.Println()
		fmt.Println("===Victoire===")
		fmt.Println("vous avez vaincu le Grandmage. Félicitation !")
		fmt.Println("vous remportez le combat")
	}
}
