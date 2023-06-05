package structure

import "goproject/constants"

type Armor struct {
	Item
	WearableItems *[]constants.ItemType
}

type BodyPartForArmor struct {
	Top   *Armor
	Pants *Armor
	Shoes *Armor
}

type Weapon struct {
	Item
	WearableItems *[]constants.ItemType
}

type BodyPartForWeapon struct {
	LeftHand  *Weapon
	RightHand *Weapon
}
