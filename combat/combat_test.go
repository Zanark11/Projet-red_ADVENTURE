package combat

import (
    "testing"

    "Projet-red_ADVENTURE/personnage"
)

func TestJouerSerie1(t *testing.T) {
    // On crée un personnage pour le test.
    joueur := personnage.InitCharacter("Test", "1")

    // On lance le combat avec ce personnage.
    JouerSerie1(&joueur)
}