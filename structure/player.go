package structure

import (
	"goproject/constants"
)

type Status struct {
	Name      string
	Room      *Room
	Inventory *Inventory
	Equipment *Equipment
	BodyPartForArmor
	BodyPartForWeapon
	Attribute Attribute
}

type Attribute struct {
	Health  int
	Attack  int
	Defense int
}

type Equipment struct {
	Top       constants.ItemType
	Pants     constants.ItemType
	Shoes     constants.ItemType
	LeftHand  constants.ItemType
	RightHand constants.ItemType
}

type Inventory map[constants.ItemType]int
