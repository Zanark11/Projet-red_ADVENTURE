package main

import (
    "Projet-red_ADVENTURE/combat"
    "Projet-red_ADVENTURE/personnage"
)

func main() {
    // On crée un personnage pour tester le combat.
    joueur := personnage.InitCharacter("Test", "1")

    // On lance le combat.
    combat.JouerSerie1(&joueur)
}