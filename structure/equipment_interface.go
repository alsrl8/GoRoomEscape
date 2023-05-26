package structure

import (
	"goproject/constants"
)

type Wearable interface {
	Equip(itemType constants.ItemType)
	Disarm()
}

func (armor *Armor) Equip(itemType constants.ItemType) {
	*armor = Armor{Item: Item{ItemType: itemType}}
}

func (armor *Armor) Disarm() {
	*armor = Armor{Item: Item{ItemType: constants.Nothing}}
}

func (weapon *Weapon) Equip(itemType constants.ItemType) {
	*weapon = Weapon{Item: Item{ItemType: itemType}}
}

func (weapon *Weapon) Disarm() {
	*weapon = Weapon{Item: Item{ItemType: constants.Nothing}}
}
