package main

// ============================================================
// ÉQUIPEMENT
// Représente un équipement du joueur.
// ============================================================

type Equipment struct {
	Name    string
	Slot    string
	BonusHP int
}

// ============================================================
// EMPLACEMENTS D'ÉQUIPEMENT
// Le joueur peut avoir un équipement sur la tête,
// le torse et les pieds.
// ============================================================

type EquipmentSlots struct {
	Head  *Equipment
	Chest *Equipment
	Feet  *Equipment
}

// ============================================================
// ÉQUIPER UN OBJET
// Place l'équipement dans l'emplacement correspondant.
// Si un équipement est déjà présent, il est remplacé.
// ============================================================

func Equip(slots *EquipmentSlots, equipment Equipment) {
	switch equipment.Slot {
	case "Tête":
		slots.Head = &equipment

	case "Torse":
		slots.Chest = &equipment

	case "Pieds":
		slots.Feet = &equipment
	}
}

// ============================================================
// DÉSÉQUIPER UN OBJET
// Retire l'équipement de l'emplacement choisi.
// ============================================================

func Unequip(slots *EquipmentSlots, slot string) {
	switch slot {
	case "Tête":
		slots.Head = nil

	case "Torse":
		slots.Chest = nil

	case "Pieds":
		slots.Feet = nil
	}
}

// ============================================================
// CALCUL DU BONUS DE PV
// Additionne les bonus de PV de tous les équipements portés.
// ============================================================

func TotalBonusHP(slots *EquipmentSlots) int {
	total := 0

	if slots.Head != nil {
		total += slots.Head.BonusHP
	}

	if slots.Chest != nil {
		total += slots.Chest.BonusHP
	}

	if slots.Feet != nil {
		total += slots.Feet.BonusHP
	}

	return total
}