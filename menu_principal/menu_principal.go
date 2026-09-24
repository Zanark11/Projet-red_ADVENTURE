package menu_Principal

import "fmt"

func MenuPrincipal() int {
	var choix int
	fmt.Println("====*MENU PRINCIPAL*====")
	fmt.Println("1.EXPLORATION DE LA NATURE")
	fmt.Println("2.INVENTAIRE")
	fmt.Println("3.MARCHAND")
	fmt.Println("4.ENTRAINEMENT")
	fmt.Println("5.Qui sont-ils ?")
	fmt.Println("6.QUITTER")
	fmt.Println("Que veux tu faire ?")
	fmt.Scanln(&choix)
	return choix
}

func Information() {
	for {
		var choix int

		fmt.Println("====Qui sont-ils====")
		fmt.Println()
		fmt.Println("1.Homme mystérieur")
		fmt.Println("2.Grandmage")
		fmt.Println("3.Retour")

		fmt.Scanln(&choix)
		switch choix {
		case 1:
			fmt.Println("====CHANAN RUFO====")
			fmt.Println()
			fmt.Println("Il fut autrefois un habitant du village qui ne cessait de prendre soin des villageois.")
			fmt.Println("Il était également un auteur à succès.")
			fmt.Println("En l'an 1455, il écrivit un livre intitulé<<les lamentations des larmes>> rempli de proverbe")
			fmt.Println("Il décède en l'an 1500")
			fmt.Println("Malgré sa mort, son fantôme revient observer les villageois.")
			fmt.Println()
			fmt.Println("C'est lui qui a envoyé la lettre au jeune avanturier.")
			fmt.Println()
		case 2:
			fmt.Println("====LE GRANDMAGE====")
			fmt.Println()
			fmt.Println("Le Grandmage est un homme intéllectuel au service d'un souverain.")
			fmt.Println("Il fut envoyé afin de stopper quiconque tenterait de faire échouer les plans de son souverain.")
			fmt.Println("Après avoir appris que les des personnes cherchaient à retrouver les villageois")
			fmt.Println("disparus, le Grandmage fut envoyé pour les arrêter.")
			fmt.Println()
			fmt.Println()
		case 3:
			return
		default:
			fmt.Println("choix invalide")

		}
	}

}
