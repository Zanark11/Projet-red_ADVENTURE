package main 

// ============================================================
// ÉQUIPEMENT
// ============================================================

type Equipment struct {
	Name    string
	Slot    string
	BonusHP int
}

// ============================================================
// EMPLACEMENTS
// ============================================================

type EquipmentSlots struct {
	Head  *Equipment
	Chest *Equipment
	Feet  *Equipment
}

// ============================================================
// ÉQUIPER
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
// DÉSÉQUIPER
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
// BONUS TOTAL DE PV
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