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
