package structure

import (
	"goproject/constants"
)

type Wearable interface {
	Equip(itemType constants.ItemType)
	Disarm()
	IsWearable(itemType constants.ItemType) bool
}

func (armor *Armor) Equip(itemType constants.ItemType) {
	*armor = Armor{Item: Item{ItemType: itemType}, WearableItems: armor.WearableItems}
}

func (armor *Armor) Disarm() {
	*armor = Armor{Item: Item{ItemType: constants.Nothing}, WearableItems: armor.WearableItems}
}

func (armor *Armor) IsWearable(itemType constants.ItemType) bool {
	for i := 0; i < len(*armor.WearableItems); i++ {
		if (*armor.WearableItems)[i] == itemType {
			return true
		}
	}
	return false
}

func (weapon *Weapon) Equip(itemType constants.ItemType) {
	*weapon = Weapon{Item: Item{ItemType: itemType}, WearableItems: weapon.WearableItems}
}

func (weapon *Weapon) Disarm() {
	*weapon = Weapon{Item: Item{ItemType: constants.Nothing}, WearableItems: weapon.WearableItems}
}

func (weapon *Weapon) IsWearable(itemType constants.ItemType) bool {
	for i := 0; i < len(*weapon.WearableItems); i++ {
		if (*weapon.WearableItems)[i] == itemType {
			return true
		}
	}
	return false
}