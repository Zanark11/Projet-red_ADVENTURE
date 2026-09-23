package menu_Principal

import "fmt"

func MenuPrincipal() int {
	var choix int
	fmt.Println("====*MENU PRINCIPAL*====")
	fmt.Println("1.EXPLORATION DE LA NATURE")
	fmt.Println("2.INVENTAIRE")
	fmt.Println("3.MARCHAND")
	fmt.Println("4.ENTRAINEMENT")
	fmt.Println("5.QUITTER")
	fmt.Println("Que veux tu faire ?")
	fmt.Scanln(&choix)
	return choix
}
