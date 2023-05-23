package data

import (
	"goproject/constants"
	"goproject/structure"
)

var ItemTypeWearableMap = map[constants.ItemType]bool{
	constants.WoodSword:    true,
	constants.IronSword:    true,
	constants.WoodShield:   true,
	constants.LeatherHat:   true,
	constants.LeatherCloth: true,
	constants.LeatherPants: true,
	constants.LeatherShoes: true,
}

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

var ItemTypeBodyPartMap = map[constants.ItemType][]constants.BodyPart{
	constants.WoodSword:    {constants.LeftHand, constants.RightHand},
	constants.IronSword:    {constants.LeftHand, constants.RightHand},
	constants.WoodShield:   {constants.LeftHand, constants.RightHand},
	constants.LeatherCloth: {constants.Top},
	constants.LeatherPants: {constants.Pants},
	constants.LeatherShoes: {constants.Shoes},
}
