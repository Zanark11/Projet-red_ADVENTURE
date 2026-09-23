package exploration

import (
	"aventure/personnage"
	"fmt"
)

// ============================================================
// OBJET
// ============================================================

type Item struct {
	Name string
}

// ============================================================
// OBJETS TROUVABLES
// ============================================================

var objetsExploration = []string{
	"Bois",
	"Pierre",
	"Bâton",
	"Encre",
	"Tube",
	"Lame",
	"Vis",
	"Verre",
	"Métal",
	"Plume magique",
	"Papier",
	"Encre arc-en-ciel",
}

// ============================================================
// EXPLORATION
// ============================================================

func Exploration(player *personnage.Character) bool {

	fmt.Println()
	fmt.Println("🌲 EXPLORATION")
	fmt.Println()

	fmt.Println("Tu explores une forêt...")
	fmt.Println()

	FindItem(player, "Bois")

	fmt.Println()
	fmt.Println("Tu continues ton exploration...")
	fmt.Println()

	FindItem(player, "Pierre")

	fmt.Println()
	fmt.Println("Tu continues à explorer...")
	fmt.Println()

	FindItem(player, "Bâton")

	fmt.Println()
	fmt.Println("Tu trouves quelque chose près d'un ancien bureau...")
	fmt.Println()

	FindItem(player, "Encre")

	fmt.Println()
	fmt.Println("Tu trouves un petit tube...")
	fmt.Println()

	FindItem(player, "Tube")

	fmt.Println()
	fmt.Println("Tu trouves une vieille lame...")
	fmt.Println()

	FindItem(player, "Lame")

	fmt.Println()
	fmt.Println("Tu trouves une vis...")
	fmt.Println()

	FindItem(player, "Vis")

	fmt.Println()
	fmt.Println("Tu trouves un morceau de verre...")
	fmt.Println()

	FindItem(player, "Verre")

	fmt.Println()
	fmt.Println("Tu trouves un morceau de métal...")
	fmt.Println()

	FindItem(player, "Métal")

	fmt.Println()
	fmt.Println("Tu trouves une plume étrange...")
	fmt.Println()

	FindItem(player, "Plume magique")

	fmt.Println()
	fmt.Println("Tu trouves une feuille de papier...")
	fmt.Println()

	FindItem(player, "Papier")

	fmt.Println()
	fmt.Println("Tu trouves une mystérieuse encre multicolore...")
	fmt.Println()

	FindItem(player, "Encre arc-en-ciel")

	fmt.Println()
	fmt.Println("✅ Exploration terminée !")

	return true
}

// ============================================================
// TROUVER UN OBJET
// ============================================================

func FindItem(player *personnage.Character, itemName string) {

	item := Item{
		Name: itemName,
	}

	player.Inventory = append(player.Inventory, item)

	fmt.Println("🔎 Tu trouves :", itemName)
	fmt.Println("✅", itemName, "a été ajouté à ton inventaire !")
}

// ============================================================
// ASSEMBLAGE
// ============================================================

func Assemble(
	player *personnage.Character,
	item1 string,
	item2 string,
	result string,
) {

	index1 := -1
	index2 := -1

	for i, item := range player.Inventory {

		if item.Name == item1 && index1 == -1 {
			index1 = i
			continue
		}

		if item.Name == item2 && index2 == -1 {
			index2 = i
		}
	}

	// Vérifier que les deux objets existent
	if index1 == -1 || index2 == -1 {
		fmt.Println()
		fmt.Println("❌ Tu n'as pas les objets nécessaires.")
		fmt.Println()
		return
	}

	// Retirer les deux objets
	newInventory := []Item{}

	for i, item := range player.Inventory {

		if i != index1 && i != index2 {
			newInventory = append(newInventory, item)
		}
	}

	// Ajouter le résultat
	newInventory = append(newInventory, Item{
		Name: result,
	})

	player.Inventory = newInventory

	fmt.Println()
	fmt.Println("🔧 ASSEMBLAGE RÉUSSI !")
	fmt.Println()
	fmt.Println(item1, "+", item2, "→", result)
	fmt.Println()
	fmt.Println("✨", result, "a été fabriqué !")
}

// ============================================================
// RECETTES
// ============================================================

func AssembleObject(player *personnage.Character, choice int) {

	switch choice {

	case 1:
		Assemble(
			player,
			"Bois",
			"Pierre",
			"🔥 Feu",
		)

	case 2:
		Assemble(
			player,
			"Bâton",
			"Pierre",
			"🔨 Marteau",
		)

	case 3:
		Assemble(
			player,
			"Encre",
			"Tube",
			"🖊️ Stylo",
		)

	case 4:
		Assemble(
			player,
			"Lame",
			"Vis",
			"✂️ Ciseaux",
		)

	case 5:
		Assemble(
			player,
			"Verre",
			"Métal",
			"🔍 Loupe",
		)

	case 6:
		Assemble(
			player,
			"Loupe",
			"Plume magique",
			"🔮 Loupe révélatrice",
		)

	case 7:
		Assemble(
			player,
			"Papier",
			"Encre arc-en-ciel",
			"🗺️ Carte magique",
		)

	default:
		fmt.Println("❌ Choix invalide.")
	}
}

// ============================================================
// MENU ASSEMBLAGE
// ============================================================

func MenuAssemblage(player *personnage.Character) {

	for {

		fmt.Println()
		fmt.Println("🔧 ASSEMBLAGE")
		fmt.Println()
		fmt.Println("1 - Bois + Pierre → 🔥 Feu")
		fmt.Println("2 - Bâton + Pierre → 🔨 Marteau")
		fmt.Println("3 - Encre + Tube → 🖊️ Stylo")
		fmt.Println("4 - Lame + Vis → ✂️ Ciseaux")
		fmt.Println("5 - Verre + Métal → 🔍 Loupe")
		fmt.Println("6 - Loupe + Plume magique → 🔮 Loupe révélatrice")
		fmt.Println("7 - Papier + Encre arc-en-ciel → 🗺️ Carte magique")
		fmt.Println("0 - Retour")
		fmt.Println()

		var choix int

		fmt.Print("Choisis une recette : ")
		fmt.Scanln(&choix)

		if choix == 0 {
			return
		}

		AssembleObject(player, choix)
	}
}

// ============================================================
// AFFICHER L'INVENTAIRE
// ============================================================

func ShowInventory(player *personnage.Character) {

	fmt.Println()
	fmt.Println("🎒 INVENTAIRE")
	fmt.Println()

	if len(player.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}

	for _, item := range player.Inventory {
		fmt.Println("-", item.Name)
	}
}