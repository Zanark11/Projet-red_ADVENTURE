package main

func main() {
	player := NewPlayer()

	inventory := []Item{}

	item := Item{
		Name:  "Potion",
		Price: 30,
	}

	maxSlots := 5

	// ============================================================
	// MARCHAND
	// ============================================================

	println("=== MARCHAND ===")
	println("Pièces au départ :", player.Gold)

	if Buy(&player, &inventory, item, maxSlots) {
		println("Achat réussi !")
	} else {
		println("Achat impossible !")
	}

	println("Pièces restantes :", player.Gold)
	println("Objets dans l'inventaire :", len(inventory))

	// ============================================================
	// ÉQUIPEMENTS
	// ============================================================

	slots := EquipmentSlots{}

	head := Equipment{
		Name:    "Chapeau de l'aventurier",
		Slot:    "Tête",
		BonusHP: 10,
	}

	chest := Equipment{
		Name:    "Tunique de l'aventurier",
		Slot:    "Torse",
		BonusHP: 20,
	}

	feet := Equipment{
		Name:    "Bottes de l'aventurier",
		Slot:    "Pieds",
		BonusHP: 10,
	}

	// Équiper les objets
	Equip(&slots, head)
	Equip(&slots, chest)
	Equip(&slots, feet)

	println()
	println("=== ÉQUIPEMENTS ===")

	println("Tête :", slots.Head.Name)
	println("Torse :", slots.Chest.Name)
	println("Pieds :", slots.Feet.Name)

	println("Bonus total de PV :", TotalBonusHP(&slots))
}
