package main

func main() {
	player := NewPlayer()

	inventory := []Item{}

	item := Item{
		Name:  "Potion",
		Price: 30,
	}

	maxSlots := 5

	println("Pièces au départ :", player.Gold)

	if Buy(&player, &inventory, item, maxSlots) {
		println("Achat réussi !")
	} else {
		println("Achat impossible !")
	}

	println("Pièces restantes :", player.Gold)
	println("Objets dans l'inventaire :", len(inventory))
}