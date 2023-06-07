package structure

import "goproject/constants"

type Weapon struct {
	Item
	WearableItems []constants.ItemType
}

func (weapon *Weapon) Equip(itemType constants.ItemType) {
	(*weapon).Item = Item{ItemType: itemType}
}

func (weapon *Weapon) Disarm() {
	(*weapon).Item = Item{ItemType: constants.Nothing}
}

func (weapon *Weapon) IsWearable(itemType constants.ItemType) bool {
	for i := 0; i < len(weapon.WearableItems); i++ {
		if (weapon.WearableItems)[i] == itemType {
			return true
		}
	}
	return false
}
