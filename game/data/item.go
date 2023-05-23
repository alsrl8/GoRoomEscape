package data

import "goproject/constants"

var ItemTypeUsableMap = map[constants.ItemType]bool{
	constants.Hammer:     true,
	constants.Key:        true,
	constants.Box:        true,
	constants.HealPotion: true,
}
