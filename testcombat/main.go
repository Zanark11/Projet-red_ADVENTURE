package main

import (
	"fmt"

	"Projet-red_ADVENTURE/combat"
)

func main() {
	// On affiche un message pour expliquer le test.
	fmt.Println("=== TEST DU CHRONO ===")
	fmt.Println()

	// Le niveau 1 donne 20 secondes pour répondre.
	fmt.Println("Tu as 20 secondes pour entrer un nombre.")
	fmt.Println()

	// On attend une réponse tout en surveillant le chrono.
	choix, aRepondu := combat.AttendreEntree(1)

	// Si le joueur a répondu avant la fin du chrono.
	if aRepondu {
		fmt.Println()
		fmt.Println("Tu as répondu :", choix)
		fmt.Println("Réponse donnée à temps !")
	} else {
		// Si le temps est dépassé.
		fmt.Println()
		fmt.Println("Temps écoulé !")
		fmt.Println("Le joueur aurait perdu 10 PV.")
	}
}
