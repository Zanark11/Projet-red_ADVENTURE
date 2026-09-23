package personnage

import "testing"

// TestTakePot vérifie que la potion de vie fonctionne correctement.
func TestTakePot(t *testing.T) {
	// On crée un personnage.
	character := InitCharacter("Test", "2")

	// Le personnage commence avec 60 PV.
	if character.pointDeVieActuel != 60 {
		t.Errorf("PV incorrects : %d au lieu de 60", character.pointDeVieActuel)
	}

	// On utilise une potion.
	TakePot(&character)

	// La potion doit rendre 50 PV.
	// 60 + 50 = 110, mais le maximum est 100.
	// Le personnage doit donc avoir 100 PV.
	if character.pointDeVieActuel != 100 {
		t.Errorf("PV incorrects après la potion : %d au lieu de 100", character.pointDeVieActuel)
	}

	// On vérifie que la potion a été supprimée.
	if character.Inventaire[0] != "" {
		t.Error("La potion devrait avoir été supprimée de l'inventaire")
	}
}

// TestTakePoisonPot vérifie que la potion de poison est retirée
// de l'inventaire sans blesser le joueur.
func TestTakePoisonPot(t *testing.T) {
	character := InitCharacter("Test", "2")

	// On ajoute une potion de poison dans l'inventaire.
	character.Inventaire[0] = "potion de poison"

	// On utilise la potion.
	utilisee := TakePoisonPot(&character)

	// La potion doit avoir été utilisée.
	if !utilisee {
		t.Error("La potion de poison aurait dû être utilisée")
	}

	// Le joueur ne doit pas perdre de PV.
	if character.pointDeVieActuel != 60 {
		t.Errorf("Le joueur a perdu des PV : %d au lieu de 60", character.pointDeVieActuel)
	}

	// La potion doit avoir été supprimée.
	if character.Inventaire[0] != "" {
		t.Error("La potion de poison devrait avoir été supprimée")
	}
}
