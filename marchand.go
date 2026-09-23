package marchand

import (
    "aventure/personnage"
    "fmt"
)

const prixAugmentationInventaire = 80

type Item struct {
    nom  string
    prix int
}

type Shop struct {
    Items []Item
}

// Vérifie si le joueur a assez d'argent
func HasEnoughMoney(player personnage.Character, item Item) bool {
    return player.Argent >= item.prix
}

func findItem(shopShop, nom string) (Item, bool) {
    for _, item := range shop.Items {
        if item.nom == nom {
            return item, true
        }
    }
    return Item{}, false
}
func HasInventorySpace(player personnage.Character) bool {
    for _, objet := range player.Inventaire {
        if objet == "" {
            return true
        }
    }
    return false
}

// Ajoute un objet dans l'inventaire
func AddItem(playerpersonnage.Character, item Item) bool {
    for i, objet := range player.Inventaire {
        if objet == "" {
            player.Inventaire[i] = item.nom
            return true
        }
    }
    return false
}

// Effectue l'achat complet
func Buy(player *personnage.Character, item Item) bool {
    // Vérifie l'argent
    if !HasEnoughMoney(player, item) {
        return false
    }

    // Vérifie la place
    if !HasInventorySpace(player) {
        return false
    }
    player.Argent -= item.prix
    return AddItem(player, item)
}

var maBoutique = Shop{
    Items: []Item{
        {nom: "potion de vie", prix: 0,},
        {nom: "livre de connaissance", prix: 50},
        {nom: "poison ", prix: 40},
        {nom: "potion de mana", prix: 40},
        {nom: "cuir du sanglier", prix: 80},
        {nom: "fourrure de loup", prix: 70},
        {nom: "plume de corbeau", prix: 60},
        {nom: "peau de troll", prix: 70},
    },
}
func Afficherobjet(player *personnage.Character) {
    for {
        var choix int
        fmt.Println("====OBJET DISPONIBLE====")
        for i, item := range maBoutique.Items {
            fmt.Println(i+1, "-", item.nom, "-", item.prix, "pièce")
        }
        fmt.Println()
        fmt.Println("1.Acheter un objet")
        fmt.Println("2.Retour")
        fmt.Scanln(&choix)

        switch choix {
        case 1:
            var choixObjet int
            fmt.Println("Quel objet veux tu acheter ?")
            fmt.Scanln(&choixObjet)
            if choixObjet < 1 || choixObjet > len(maBoutique.Items) {
                fmt.Println("choix invalide")
                continue
            }
            objet := maBoutique.Items[choixObjet-1]
            fmt.Println("tu as choisi :", objet.nom)
            if Buy(player, objet) {
                fmt.Println("Achat réussi !")
            } else {
                fmt.Println("Achat impossible")
            }
        case 2:
            return
        default:
            fmt.Println("choix invalide")
        }
    }
}
func AmeliorationInventaire(player *personnage.Character) bool {
    for {
        var choix int

        fmt.Println("====AMELIORER SON INVENTAIRE====")
        fmt.Println()
        fmt.Println("capacité actuelle:", len(player.Inventaire))
        fmt.Println("prix:", prixAugmentationInventaire, "pièce pour augmenter votre inventaire")
        fmt.Println("1.Acheter une augmentation de capacité d'inventaire")
        fmt.Println("2.Retour")

        fmt.Scanln(&choix)
        switch choix {
        case 1:
            if AugmenterInventaire(player) {
                fmt.Println("Amélioration achétée !")
                fmt.Println("Nouvelle capacité:", len(player.Inventaire))
            } else {
                fmt.Println("Tu n'as pas assez de pièce d'argent")
            }
        case 2:
            return true

        default:
            fmt.Println("choix invalide")
        }
    }
}