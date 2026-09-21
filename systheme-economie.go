package main

type Player struct {
	Gold int
}

type Item struct {
	Name  string
	Price int
}

// Crée un joueur avec 100 pièces d'or
func NewPlayer() Player {
	return Player{
		Gold: 100,
	}
}

// Vérifie si le joueur peut payer
func CanBuy(player Player, item Item) bool {
	return player.Gold >= item.Price
}

// Retire l'argent lors d'un achat
func BuyItem(player *Player, item Item) bool {
	if !CanBuy(*player, item) {
		return false
	}

	player.Gold -= item.Price
	return true
}