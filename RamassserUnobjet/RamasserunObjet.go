package ramassserunobjet

import (
	"aventure/personnage"
	"fmt"
)

func RamasserUnobjet(player *personnage.Character, objet string ) {
	if HasInventorySpace(player){
		AddItem(player, Item{nom:objet})
	}else {
		fmt.Println("votre invenraire est plein")
	}
}