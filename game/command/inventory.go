package command

import (
	"goproject/constants"
	"goproject/structure"
)

func removeItemInInventory(inventory *structure.Inventory, itemType constants.ItemType) {
	(*inventory)[itemType] -= 1
}

func hasItemInInventory(inventory structure.Inventory, itemType constants.ItemType) bool {
	if inventory[itemType] > 0 {
		return false
	} else {
		return true
	}
}
