package combat

import (
	"fmt"
	"math/rand"
	"time"

	"Projet-red_ADVENTURE/adversaire"
	"Projet-red_ADVENTURE/personnage"
)

// Question représente une question du combat.
type Question struct {
	Texte        string
	Reponses     [3]string
	BonneReponse int
}

// PileOuFace choisit au hasard qui commence le combat.
func PileOuFace() string {
	// rand.Intn(2) donne soit 0, soit 1.
	resultat := rand.Intn(2)

	if resultat == 0 {
		return "pile"
	}

	return "face"
}

// ChoisirReponseGrandSage choisit une réponse au hasard.
func ChoisirReponseGrandSage() int {
	// On choisit un nombre entre 1 et 3.
	return rand.Intn(3) + 1
}

// PossedeObjet vérifie si le joueur possède un objet.
func PossedeObjet(joueur *personnage.Character, objet string) bool {
	for i := 0; i < len(joueur.Inventaire); i++ {
		if joueur.Inventaire[i] == objet {
			return true
		}
	}

	return false
}

// UtiliserPoison inflige 10 dégâts au Grand Sage
// pendant 3 secondes.
func UtiliserPoison(gobelin *adversaire.Adversaire) {
	for i := 0; i < 3; i++ {
		adversaire.PerdrePV(gobelin, 10)

		fmt.Println("Poison !")
		fmt.Println("Le Grand Sage perd 10 PV.")
		adversaire.AfficherPV(*gobelin)

		// On attend 1 seconde avant le prochain dégât.
		if i < 2 {
			time.Sleep(1 * time.Second)
		}
	}
}

// UtiliserChronoboros utilise le sort Chronoboros.
// Le sort coûte 40 mana et réduit le temps de réponse de 5 secondes.
func UtiliserChronoboros(joueur *personnage.Character) bool {
	// On vérifie si le joueur possède assez de mana.
	if personnage.GetMana(*joueur) < 40 {
		fmt.Println("Vous n'avez pas assez de mana.")
		fmt.Println("Il faut 40 mana pour utiliser Chronoboros.")
		return false
	}

	// On retire 40 mana au joueur.
	personnage.DepenserMana(joueur, 40)

	// On réduit le temps de réponse de 5 secondes.
	joueur.ReductionTemps(5)

	fmt.Println("Chronoboros utilisé !")
	fmt.Println("40 mana consommés.")
	fmt.Println("Votre temps de réponse est réduit de 5 secondes.")

	return true
}

// MenuObjetCombat permet au joueur d'utiliser un objet pendant le combat.
// La fonction retourne true si l'objet utilisé termine le tour.
// Elle retourne false si le joueur choisit de revenir au combat.
func MenuObjetCombat(
	joueur *personnage.Character,
	gobelin *adversaire.Adversaire,
) bool {
	for {
		fmt.Println()
		fmt.Println("========== OBJETS ==========")
		fmt.Println("1 - Potion de vie")
		fmt.Println("2 - Potion de poison")
		fmt.Println("3 - Potion de mana")
		fmt.Println("4 - Chronoboros")
		fmt.Println("5 - Retour")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			// On vérifie si le joueur possède une potion de vie.
			if !PossedeObjet(joueur, "potion de vie") {
				fmt.Println("Vous n'avez pas de potion de vie.")
				continue
			}

			// La potion de vie rend 50 PV.
			personnage.TakePot(joueur)

			// La potion de vie utilise le tour du joueur.
			return true

		case 2:
			// On vérifie si le joueur possède une potion de poison.
			if !PossedeObjet(joueur, "potion de poison") {
				fmt.Println("Vous n'avez pas de potion de poison.")
				continue
			}

			// On retire la potion de l'inventaire.
			personnage.TakePoisonPot(joueur)

			// Le poison attaque uniquement le Grand Sage.
			UtiliserPoison(gobelin)

			// La potion de poison utilise le tour.
			return true

		case 3:
			// On vérifie si le joueur possède une potion de mana.
			if !PossedeObjet(joueur, "potion de mana") {
				fmt.Println("Vous n'avez pas de potion de mana.")
				continue
			}

			// On utilise la potion de mana.
			personnage.TakeManaPot(joueur)

			// La potion de mana ne termine pas le tour.
			// Le joueur revient donc au menu des objets.
			continue

		case 4:
			// On vérifie si le joueur possède le livre de sort.
			if !PossedeObjet(joueur, "livre de sort : chronoboros") {
				fmt.Println("Vous ne possédez pas le sort Chronoboros.")
				continue
			}

			// On utilise Chronoboros.
			utilise := UtiliserChronoboros(joueur)

			// Si Chronoboros est utilisé, le tour est terminé.
			if utilise {
				return true
			}

			// Si le joueur n'a pas assez de mana,
			// il reste dans le menu.
			continue

		case 5:
			// Le joueur revient au menu du combat.
			return false
		}
	}
}

// Serie1 contient les 3 questions du premier combat.
var Serie1 = [9]Question{
	{
		Texte: "Quelle est la planète la plus proche du Soleil ?",
		Reponses: [3]string{
			"Mars",
			"Mercure",
			"Venus",
		},
		BonneReponse: 2,
	},
	{
		Texte: "Combien y a-t-il de continents sur Terre ?",
		Reponses: [3]string{
			"5",
			"6",
			"7",
		},
		BonneReponse: 3,
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
		Texte: "Quelle est la capitale de la France ?",
		Reponses: [3]string{
			"Lyon",
			"Paris",
			"Marseille",
		},
		BonneReponse: 2,
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
		Texte: "Combien de jours compte une semaine ?",
		Reponses: [3]string{
			"5",
			"7",
			"10",
		},
		BonneReponse: 2,
	},
	{
		Texte: "Quel langage utilisons-nous pour développer Red Adventure ?",
		Reponses: [3]string{
			"Python",
			"Java",
			"Go",
		},
		BonneReponse: 3,
	},
}

// JouerSerie1 lance le combat de la série 1.
func JouerSerie1(joueur *personnage.Character) {
	// On crée le Grand Sage.
	gobelin := adversaire.CreerAdversaire()

	// On lance le pile ou face pour savoir qui commence.
	resultat := PileOuFace()

	debut := 0

	// true = le joueur commence.
	// false = le Grand Sage commence.
	tourJoueur := resultat == "pile"

	fmt.Println("Pile ou face...")

	if resultat == "pile" {
		fmt.Println("Résultat : PILE")
		fmt.Println("Le joueur commence !")
	} else {
		fmt.Println("Résultat : FACE")
		fmt.Println("Le Grand Sage commence !")
	}

	fmt.Println()

	fmt.Println("========== COMBAT - SÉRIE 1 ==========")
	fmt.Println()

	// On fait les 3 questions de la série.
	for !personnage.EstVaincu(*joueur) && !adversaire.EstVaincu(gobelin) {
		// Tour du joueur.
		if tourJoueur {
			fmt.Println("Choisissez une question :")
			fmt.Println()

			// Le joueur peut choisir d'utiliser un objet
			// avant de répondre à une question.
			fmt.Println()
			fmt.Println("========== ACTION ==========")
			fmt.Println("1 - Continuer le combat")
			fmt.Println("2 - Utiliser un objet")
			fmt.Print("Votre choix : ")

			var action int
			fmt.Scanln(&action)

			if action == 2 {
				// On ouvre le menu des objets.
				utilise := MenuObjetCombat(joueur, &gobelin)

				// Si un objet a été utilisé et qu'il termine le tour,
				// on passe directement au tour du Grand Sage.
				if utilise {
					personnage.AfficherPV(*joueur)
					adversaire.AfficherPV(gobelin)

					tourJoueur = false
					continue
				}

				// Si le joueur choisit "Retour",
				// on revient au choix de l'action.
				continue
			}

			if action != 1 {
				fmt.Println("Choix invalide.")
				continue
			}
			// On affiche seulement les 3 questions du lot actuel.
			for j := 0; j < 3; j++ {
				fmt.Println(j+1, "-", Serie1[debut+j].Texte)
			}

			fmt.Print("\nVotre choix : ")

			fmt.Println()

			// Le joueur a un temps limité pour choisir une question.
			choix, aRepondu := AttendreEntree(personnage.TempsReponse(*joueur))

			// Si le joueur n'a pas répondu à temps.
			if !aRepondu {
				fmt.Println("Le temps est écoulé !")
				fmt.Println("Le joueur perd 10 PV.")

				personnage.PerdrePV(joueur, 10)

				personnage.AfficherPV(*joueur)
				adversaire.AfficherPV(gobelin)

				fmt.Println()

				// Le tour du joueur est terminé.
				tourJoueur = false

				continue
			}

			// On vérifie que le choix est entre 1 et 3.
			if choix < 1 || choix > 3 {
				fmt.Println("Choix invalide.")
				continue
			}
			// On récupère la question choisie.
			question := Serie1[debut+choix-1]

			fmt.Println()
			fmt.Println("Question :", question.Texte)
			fmt.Println()

			// On affiche les 3 réponses.
			for j := 0; j < len(question.Reponses); j++ {
				fmt.Println(j+1, "-", question.Reponses[j])
			}

			// Le Grand Sage choisit automatiquement une réponse.
			reponse := ChoisirReponseGrandSage()

			fmt.Println("Le Grand Sage choisit la réponse :", reponse)

			fmt.Println()

			// On vérifie la réponse.
			if reponse == question.BonneReponse {
				fmt.Println("Bonne réponse !")
				fmt.Println("Le joueur perd 40 PV.")

				personnage.PerdrePV(joueur, 40)
			} else {
				fmt.Println("Mauvaise réponse !")
				fmt.Println("Le Grand Sage perd 40 PV.")

				adversaire.PerdrePV(&gobelin, 40)
			}

			// On affiche les PV du Gobelin.
			personnage.AfficherPV(*joueur)
			adversaire.AfficherPV(gobelin)

			fmt.Println()

			// On passe au lot de 3 questions suivant.
			debut += 3

			// Si on arrive après la dernière question,
			// on recommence avec le premier lot.
			if debut >= len(Serie1) {
				debut = 0
			}

			// Le tour du joueur est terminé.
			tourJoueur = false

		} else {
			// Tour du Grand Sage.
			fmt.Println("========== TOUR DU GRAND SAGE ==========")
			fmt.Println()

			// Le Grand Sage choisit au hasard une des 3 questions du lot.
			indiceQuestion := rand.Intn(3)
			question := Serie1[debut+indiceQuestion]

			fmt.Println("Le Grand Sage choisit une question :")
			fmt.Println()
			fmt.Println("Question :", question.Texte)
			fmt.Println()

			// Le joueur doit choisir une réponse parmi les 3.
			for j := 0; j < len(question.Reponses); j++ {
				fmt.Println(j+1, "-", question.Reponses[j])
			}

			fmt.Println()
			fmt.Print("Votre réponse : ")

			var reponseJoueur int
			fmt.Scanln(&reponseJoueur)

			// On vérifie si la réponse du joueur est correcte.
			if reponseJoueur == question.BonneReponse {
				fmt.Println("Bonne réponse !")
				fmt.Println("Le Grand Sage perd 40 PV.")

				adversaire.PerdrePV(&gobelin, 40)
			} else {
				fmt.Println("Mauvaise réponse !")
				fmt.Println("Le joueur perd 40 PV.")

				personnage.PerdrePV(joueur, 40)
			}

			fmt.Println()
			personnage.AfficherPV(*joueur)
			adversaire.AfficherPV(gobelin)
			fmt.Println()

			// On passe au lot de 3 questions suivant.
			debut += 3

			// Si on arrive après la dernière question,
			// on recommence avec le premier lot.
			if debut >= len(Serie1) {
				debut = 0
			}

			// Le tour du Grand Sage est terminé.
			tourJoueur = true
		}
	}
	// On vérifie qui a gagné le combat.
	if personnage.EstVaincu(*joueur) {
		fmt.Println("================================")
		fmt.Println("DÉFAITE !")
		fmt.Println("Le Grand Sage a gagné le combat.")
		fmt.Println("================================")
	} else if adversaire.EstVaincu(gobelin) {
		fmt.Println("================================")
		fmt.Println("VICTOIRE !")
		fmt.Println("Vous avez vaincu le Grand Sage !")
		personnage.GagnerCombat(joueur)
		fmt.Println("================================")
	}
}
