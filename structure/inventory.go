package structure

import "goproject/constants"

func (inventory Inventory) GetItemNumber(itemType constants.ItemType) int {
	return inventory[itemType]
}
