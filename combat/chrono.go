package combat

import (
	"fmt"
	"time"
)

// TempsDeReponse retourne le temps disponible selon le niveau du joueur.
func TempsDeReponse(niveau int) int {
	// Le niveau 1 donne 20 secondes.
	// Chaque niveau ajoute 5 secondes.
	temps := 20 + (niveau-1)*5

	// Le niveau maximum est 10.
	// Donc le temps maximum est 65 secondes.
	if niveau > 10 {
		temps = 65
	}

	return temps
}

// Chronometre lance un compte à rebours.
func Chronometre(niveau int) {
	temps := TempsDeReponse(niveau)

	fmt.Println("Vous avez", temps, "secondes pour répondre.")

	// On crée un compte à rebours.
	for i := temps; i > 0; i-- {
		fmt.Printf("\rTemps restant : %d secondes", i)

		// On attend une seconde avant de diminuer le temps.
		time.Sleep(1 * time.Second)
	}

	fmt.Println()
	fmt.Println("Temps écoulé !")
}

// AttendreEntree permet de récupérer une réponse du joueur
// tout en affichant le temps restant sur une seule ligne.
func AttendreEntree(temps int) (int, bool) {
	// On crée un channel pour recevoir la réponse.
	reponse := make(chan int, 1)

	// On lance une goroutine qui attend la saisie du joueur.
	go func() {
		var choix int
		fmt.Scanln(&choix)

		// On envoie le choix dans le channel.
		reponse <- choix
	}()

	// On crée un ticker qui déclenche une fois par seconde.
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// On affiche le premier message.
	fmt.Printf("\rTu as %d secondes pour entrer un nombre", temps)

	// On compte les secondes restantes.
	for secondes := temps - 1; secondes >= 0; secondes-- {
		select {
		case choix := <-reponse:
			// Le joueur a répondu avant la fin du chrono.
			fmt.Println()
			return choix, true

		case <-ticker.C:
			if secondes == 0 {
				fmt.Println()
				fmt.Println("Temps écoulé !")
				return 0, false
			}

			// \r revient au début de la même ligne.
			fmt.Printf("\rTu as %d secondes pour entrer un nombre", secondes)
		}
	}

	return 0, false
}
