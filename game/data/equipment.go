package data

import (
	"goproject/constants"
	"goproject/structure"
)

var EquipmentEffectMap = map[constants.ItemType]structure.Attribute{
	constants.WoodSword: {
		Health:  0,
		Attack:  3,
		Defense: 0,
	},
	constants.IronSword: {
		Health:  0,
		Attack:  10,
		Defense: 0,
	},
	constants.WoodShield: {
		Health:  0,
		Attack:  0,
		Defense: 0,
	},
	constants.LeatherHat: {
		Health:  0,
		Attack:  0,
		Defense: 0,
	},
	constants.LeatherCloth: {
		Health:  0,
		Attack:  0,
		Defense: 6,
	},
	constants.LeatherPants: {
		Health:  0,
		Attack:  0,
		Defense: 4,
	},
	constants.LeatherShoes: {
		Health:  0,
		Attack:  0,
		Defense: 3,
	},
}
