package menuexploration

import (
	"aventure/personnage"
	"aventure/RamasserunObjet"
	"aventure/combat"
	"fmt"
)

func Explorer(perso *personnage.Character) {
	
	niveau2 := false

	for {
		var choix int

		fmt.Println("=====EXPLORATION====")
		fmt.Println()
		fmt.Println("1.Niveau n°1")
		fmt.Println("2.Niveau n°2")
		fmt.Println("3.Retour")

		fmt.Scanln(&choix)
		switch choix {
		case 1:
			var choix int
			fmt.Println("====NIVEAU 1====")
			fmt.Println()
			fmt.Println("BRHI- le mystère de la forêt")
			fmt.Println()
			fmt.Println("Vous arrivez dans une forêt inconnue après avoir reçu une étrange lettre")
			fmt.Println("lettre:<< si tu veux découvrir ce qui est arrivé au village de BRHI, rends-toi dans la forêt. Mais méfie-toi:")
			fmt.Println("les réponses sont parfois plus importantes que la force.>>")
			fmt.Println()
			fmt.Println("vous êtes dans la forêt. Votre objectif initial est de trouver le village BRHI.")
			fmt.Println("Pendant que vous avancez sur le chemin, vous voyez une maison abandonnée et un puits")
			fmt.Println()
			fmt.Println("1.Aller vers la maison abandonnée")
			fmt.Println("2.Explorer le puits")
			fmt.Println("3.Continuer sur le chemin")
			fmt.Println("4.retour")
			fmt.Scanln(&choix)
			switch choix {
			case 1:
				porteOuverte := false
				/*énigme de la porte */
				for !porteOuverte {
					var choix int
					fmt.Println("Vous arrivez devant la maison.")
					fmt.Println("vous entrez dans la maison et vous constatez qu'elle est dans un état insalubre et décidez")
					fmt.Println("quand même de fouiller. Durant votre fouille vous constatez que l'une des portes est fermée.")
					fmt.Println("avec une inscription, << pour ouvrir la porte resolvez cette équation 2x + 2 = 0 >>")
					fmt.Println()
					fmt.Println("1.la réponse est 0")
					fmt.Println("2.la réponse est -1")
					fmt.Println("3.la réponse est 1")
					fmt.Println("4.Abandonée")
					fmt.Scanln(&choix)
					/*choix de la reponse*/
					switch choix {
					case 1:
						fmt.Println("Mauvaise réponse !")
					case 2:
						porteOuverte = true
						fmt.Println("Bonne réponse! la porte est ouverte")
						fmt.Println()
						fmt.Println("vous entrez dans la pièce et tout est bien rangé. intrigué que cette pièce soit propre,")
						fmt.Println("vous regardez autour de vous et voyez des photos de famille, et surtout une photo d'une jeune fille avec une inscription derrière")
						fmt.Println("écrire avec du sang : Au secour!")
						fmt.Println("vous gardez la photo, plus déterminé que jamais à retrouver les villageois")
						fmt.Println("vous continuez de fouiller et vous tomber sur des pièces.")
						fmt.Println()
						/*choix de prendre les pièce*/
						fmt.Println("1.vous ramasser les 70 pièces")
						fmt.Println("2.vous ne voulez pas rammasser les 70 pièces")
						fmt.Scanln(&choix)
						switch choix {
						case 1:
							perso.Argent += 70
							fmt.Println("Vous avez maintenant", perso.Argent, "pièces")
						case 2:
							fmt.Println("vous laissez les pièces")
						default:
							fmt.Println("choix invalide")
						}
						/*boite d'encre*/
                        fmt.Println()
						fmt.Println("vous continuez votre fouille et vous tombez sur une boîte à encre couleur de  l'arc-en-ciel")
						fmt.Println()
						fmt.Println("1.prenez_vous cette boîte ?")
						fmt.Println("2.Refuser_vous cette boîte ?")
						fmt.Scanln(&choix)
						switch choix {
						case 1:
							// appel la fonction pour ramasser les objets
							ramasserunobjet.RamasserUnobjet(perso, "boîte d'encre")
						case 2:
							fmt.Println("vous refusez la boîte")
						default:
							fmt.Println("choix invalide")
						}
			        case 3:
						fmt.Println("Mauvaise réponse !")
					case 4:
						return
					default:
						fmt.Println("choix invalide")
					   }
					if porteOuverte {
						fmt.Println(" vous sortez de la maison.")
					} else {
						fmt.Println("la porte reste fermée.")
					}
				    }
			case 2:
                fmt.Println("vous vous approchez du puits et il vous semble très ancien.")
			    fmt.Println("vous regardez à l'intérieur....")
			    fmt.Println("il fait complètement noir.")
				fmt.Println()
				fmt.Println("vous remarquer quelque chose accroché au bord du puits")
				fmt.Println("c'est une vieille corde")
				
			    fmt.Println("vous décidez de continuer...")
			case 3:
				fmt.Println()
				fmt.Println("Vous continuez sur le chemin.")
			case 4:
				return 
			default:
				fmt.Println("choix invalide")
				}
				fmt.Println("Après cela vous avancez  et vous remarquez une ombre . ")
				fmt.Println("vous décidez de continuer malgré cela et au fur et à mesure l'ombre devient claire")
				fmt.Println("vous êtes à 2 mètre de l'ombre et vous vous rendez compte qu'il s'agit d'une personne")
				fmt.Println("la personne est de dos , vous avancez à petits pas dans le but de l'immobiliser et savoir qui c'est.")
				fmt.Println("Mais vous marchez sur une branche qui fait du bruit et là, la personne se retourne et vous voit.")
				fmt.Println("Elle vous dit je suis <<le Grandmage>> et je t'attendais. si tu veux passer, montre moi que tu est un aventurier")
				fmt.Println("dont la force se trouve dans la tête")
				combat.JouerSerie1(perso)
				fmt.Println()
				fmt.Println("vous avez battu le Grandmage .")
				fmt.Println("vous constatez une clé à coté du Grandmage. vous prenez la clé.")
				fmt.Println("vous continuez et sur votre chemin vous croisez une personne qui vous dit:")
				fmt.Println("<<N'oublie pas la route est longue, aventurier...")
				fmt.Println("mais souviens-toi: qui va doucement va sûrement. surtout ne te précipite pas. Observe,")
				fmt.Println("réfléchis...et tu trouveras ce que tu cherches.")
				fmt.Println("vous vous approchez afin de savoir qui est ce ? Mais il disparaît.")
				fmt.Println("vous continuer votre chemin et vous vous retrouvez face à une grande porte ne sachant comment l'ouvrir")
				fmt.Println("et là vous vous rappelez de la clé du Grandmage.")
				fmt.Println("vous l'utiliser pour ouvrir la grande porte")
				fmt.Println("vous entrez et quelques mètre plus loin  vous voyez une pancarte avec l'enseigne << VILLAGE BRIH>>")
				fmt.Println("félicitation car vous avez retrouver le village. Rendez vous au niveau 2 pour retrouver les villageois")

	    case 2:
	        if !niveau2{
				fmt.Println("le niveau 2 sortira le 30/09/2026")
				fmt.Println("Avec plus de fonctionnalité")
				continue
			}
		case 3:
			return
		default:
			fmt.Println("choix invalide")
	    }
    }
}