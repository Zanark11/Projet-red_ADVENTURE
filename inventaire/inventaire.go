package inventaire

import ("fmt"
"aventure/personnage"
)
func AccessInventory(character personnage.Character) bool {
	var choix int
	for{
	fmt.Println("====Inventaire====")
	fmt.Println("1.Afficher mon inventaire")
	fmt.Println("2.Retour")
	fmt.Println("Que veux tu faire ?")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
        fmt.Println("====inventaire====")
	compteur := 0
	for i, item := range character.Inventaire {
		fmt.Println(i+1, "-", item)
		if item != "" {
			compteur++
		}
	}
	fmt.Println("1.Jeter un objet")
	fmt.Println("2.Fabriquer un objet")
	fmt.Println("3.Retour")
	fmt.Println("Que veux tu faire ?")
	fmt.Scanln(&choix)

	switch choix{
	case 1:
		personnage.JeterObjet(&character)
	case 2:
		/*mettre le code pour fabriquer un objet */
	}
    case 2:
		return false
	default:
		fmt.Println("choix invalide")
        }
	
    }
}



