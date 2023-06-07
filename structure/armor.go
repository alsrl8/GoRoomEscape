package structure

import "goproject/constants"

type Armor struct {
	Item
	WearableItems []constants.ItemType
}

func (armor *Armor) Equip(itemType constants.ItemType) {
	(*armor).Item = Item{ItemType: itemType}
}

func (armor *Armor) Disarm() {
	(*armor).Item = Item{ItemType: constants.Nothing}
}

func (armor *Armor) IsWearable(itemType constants.ItemType) bool {
	for i := 0; i < len(armor.WearableItems); i++ {
		if (armor.WearableItems)[i] == itemType {
			return true
		}
	}
	return false
}
