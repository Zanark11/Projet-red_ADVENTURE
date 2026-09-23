package main

type Shop struct {
	Items []Item
}
var shop = Shop{
	Items: []Item{
		{Name: "Potion de vie", Price: 20},
		{Name: "Livre de sort", Price: 50},
		{Name: "Poison", Price: 30},
		{Name: "Potion de mana", Price: 30},
	},
}

// Vérifie si le joueur a assez d'argent
func HasEnoughMoney(player *Player, item Item) bool {
	return player.Gold >= item.Price
}

// Vérifie si l'inventaire possède encore une place
func HasInventorySpace(player *Player, maxSlots int, inventory []Item) bool {
	return len(inventory) < maxSlots
}

// Ajoute un objet dans l'inventaire
func AddItem(inventory *[]Item, item Item, maxSlots int) bool {
	if len(*inventory) >= maxSlots {
		return false
	}

	*inventory = append(*inventory, item)
	return true
}

// Effectue l'achat complet
func Buy(player *Player, inventory *[]Item, item Item, maxSlots int) bool {
	// Vérifie l'argent
	if !HasEnoughMoney(player, item) {
		return false
	}

	// Vérifie la place
	if !HasInventorySpace(player, maxSlots, *inventory) {
		return false
	}

	// Retire l'argent
	if !BuyItem(player, item) {
		return false
	}

	// Ajoute l'objet
	if !AddItem(inventory, item, maxSlots) {
		return false
	}

	return true
}

// Cherche un objet dans la boutique
func FindItem(shop *Shop, name string) (Item, bool) {
	for _, item := range shop.Items {
		if item.Name == name {
			return item, true
		}
	}

	return Item{}, false
}
