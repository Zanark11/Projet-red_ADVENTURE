package personnage

import (
	"aventure/adversaire"
	"fmt"
	"time"
)

func CreationPersonnage() Character {

	var nom string
	var choixclasse string
	for {
		valide := true
		fmt.Println("======*CREATION DU PERSONNAGE*======")
		fmt.Println("Donne un nom à ton perso ? ")
		fmt.Scanln(&nom)

		if nom == "" {
			fmt.Println("le nom ne peut pas être vide")
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
	fmt.Println("Equipement", character.equipement)
}
func InitCharacter(nom string, choixclasse string) Character {
	pointDeVieActuel := 100

	switch choixclasse {
	case "1":
		pointDeVieActuel = 200
	case "2":
		pointDeVieActuel = 180
	case "3":
		pointDeVieActuel = 300
	}

	equipement := Porterequipement()
	character := Character{
		nom:              nom,
		pointDeVieActuel: pointDeVieActuel,
		classe:           choixclasse,
		niveau:           1,
		pointsDeVieMax:   400,
		Inventaire:       []string{"potion de vie", "potion de vie", "potion de vie", "", "", ""},
		xp:               0,
		mana:             100,
		Argent:           500,
		equipement:       equipement,
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
	equipement       Equipement
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

func Affichevie(perso *Character) {
	fmt.Println("vie", perso.pointDeVieActuel, "/", perso.pointsDeVieMax)
}

func Playervivant(perso *Character) bool {
	return perso.pointDeVieActuel > 0
}

func Degats(perso *Character, degats int) {
	perso.pointDeVieActuel -= degats
}

type Equipement struct {
	Tête  string
	Torse string
	Pieds string
}

func Porterequipement() Equipement {
	var choix int
	var equipement Equipement

	fmt.Println("===Equipement pour l'aventure===")
	fmt.Println()
	fmt.Println("choisir ton vêtement pour le Torse")
	fmt.Println("1.Un débardeur")
	fmt.Println("2.Un tee-shirt")
	fmt.Println("3.Aucun")

	fmt.Scanln(&choix)

	switch choix {
	case 1:
		equipement.Torse = "Un débardeur"
	case 2:
		equipement.Torse = "Un tee-shirt"
	case 3:
		equipement.Torse = "Aucun"
	}

	fmt.Println("choisir ton vêtement pour la tête")
	fmt.Println("1.Une casquette")
	fmt.Println("2.un bornet")
	fmt.Println("3.Aucun")

	fmt.Scanln(&choix)
	switch choix {
	case 1:
		equipement.Tête = "Une casquette"
	case 2:
		equipement.Tête = "Un bornet"
	case 3:
		equipement.Tête = "Aucun"
	}

	fmt.Println("choisir tes chaussures")
	fmt.Println("1.Des basquettes")
	fmt.Println("2.Des sandales")
	fmt.Println("3.Aucun")

	fmt.Scanln(&choix)

	switch choix {
	case 1:
		equipement.Pieds = "Des basquettes"
	case 2:
		equipement.Pieds = "Des sandales"
	case 3:
		equipement.Pieds = "Aucun"
	}
	return equipement
}

func FabriqueObjet(player *Character) {
	var choix int

	fmt.Println("====Fabrique d'objet====")
	fmt.Println()
	fmt.Println("1.Châpeau de l'aventurier ( Plume de Corbeau et Cuir de Sanglier)")
	fmt.Println("2.Tunique de l'aventurier ( Fourrures de Loup et Peau de Troll)")
	fmt.Println("3.Bottes de l'aventurier (Fourrure de Loup et Cuir de Sanglier)")
	fmt.Println("4.Retour")

	fmt.Scanln(&choix)
	switch choix {
	case 1:
		plume_de_corbeau := false
		cuir_de_sanglier := false
		for _, objet := range player.Inventaire {
			if objet == "plume de corbeau" {
				plume_de_corbeau = true

			}
			if objet == "cuir du sanglier" {
				cuir_de_sanglier = true

			}
		}
		if !plume_de_corbeau || !cuir_de_sanglier {
			fmt.Println("Tu n'as pas les ressources nécéssaires. Vous pouvez l'acheter chez le marchand")
			return
		}
	case 2:
		fourrure_de_loup := false
		peau_de_troll := false
		for i, objet := range player.Inventaire {
			if objet == "fourrure de loup" {
				fourrure_de_loup = true
				player.Inventaire[i] = ""
				break
			}
			if objet == "peau de troll" {
				peau_de_troll = true
				player.Inventaire[i] = ""
				break

			}
		}
		if !fourrure_de_loup || !peau_de_troll {
			fmt.Println("Tu n'as pas les ressources nécéssaire. Vous pouvez l'acheter chez le marchand")
			return
		}
	case 3:
		fourrure_de_loup := false
		cuir_de_sanglier := false
		for i, objet := range player.Inventaire {
			if objet == "fourrure de loup" {
				fourrure_de_loup = true
				player.Inventaire[i] = ""
				break
			}
			if objet == "cuir du sanglier" {
				cuir_de_sanglier = true
				player.Inventaire[i] = ""
				break
			}
		}
		if !cuir_de_sanglier || !fourrure_de_loup {
			fmt.Println("Tu n'as pas les ressources nécéssaires. Vous pouvez l'acheter chez le marchand")
		}
	case 4:
		return
	default:
		fmt.Println("choix invalide")
	}
}

func UtiliserObjet(player *Character, ennemi *adversaire.Adversaire) {
	var choix int
	fmt.Println("===Utiliser un objet===")
	for i, objet := range player.Inventaire {
		fmt.Println(i+1, "-", objet)
	}
	fmt.Println("Quel objet veux-tu utiliser ?")
	fmt.Scanln(&choix)

	if choix < 1 || choix > len(player.Inventaire) {
		fmt.Println("choix invalide")
		return
	}
	index := choix - 1

	objet := player.Inventaire[index]

	if objet == "" {
		fmt.Println("Cette case est vide")
		return
	}
	switch objet {
	case "potion de vie":
		player.pointDeVieActuel += 40
		fmt.Println("Tu récupères 40 points de vie ")
		fmt.Println("Ta vie est de :", player.pointDeVieActuel, "/", player.pointsDeVieMax)
		player.Inventaire[index] = ""
	case "livre de sort":
		if player.mana < 50 {
			fmt.Println("tu ne peux utiliser le livre de sort car pas assez de mana. Tu peux l'acheter chez")
			fmt.Println("le marchand")
			return
		}
		player.mana -= 50
		fmt.Println("mana restant :", player.mana)

		adversaire.Degats(ennemi, 20)
		fmt.Println("le livre de sort inflige 20 dégats sur la vie du Grandmage!")
		adversaire.Affichevie(*ennemi)
		player.Inventaire[index] = ""
	case "potion de mana":
		player.mana += 50
		fmt.Println("potion de mana utilisée")
		fmt.Println("mana actuel:", player.mana)
		player.Inventaire[index] = ""
	case "châpeau de l'aventurier":
		Augmenterviemax(player, 20)
		fmt.Println("Tu utilises le châpeau de l'aventurier !")
		fmt.Println("Ta vie max augmente de +20")
		fmt.Println("vie max:", player.pointsDeVieMax)
		player.Inventaire[index] = ""
	case "Tunique de l'aventurier":
		Augmenterviemax(player, 20)
		fmt.Println("Tu utilises Tunique de l'aventurier !")
		fmt.Println("Ta vie max augmente de +20")
		fmt.Println("vie max:", player.pointsDeVieMax)
		player.Inventaire[index] = ""
	case "Bottes de l'aventurier":
		Augmenterviemax(player, 20)
		fmt.Println("Tu utilises Bottes de l'aventurier !")
		fmt.Println("Ta vie max augmente de +20")
		fmt.Println("vie max:", player.pointsDeVieMax)
		player.Inventaire[index] = ""
	case "poison":
		poison(ennemi)
		player.Inventaire[index] = ""
	}
}
func Augmenterviemax(player *Character, bonnus int) {
	player.pointsDeVieMax += bonnus
}

func AjouterXp(player *Character, xpGagnes int) {
	player.xp += xpGagnes
	fmt.Println("Tu gagnes", xpGagnes, "xp")
	fmt.Println("xp:", player.xp)
}
func poison(ennemi *adversaire.Adversaire) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		adversaire.Degats(ennemi, 3)
		fmt.Println("le poison inflige des dégats pendant 3s")
		adversaire.Affichevie(*ennemi)
	}
}
