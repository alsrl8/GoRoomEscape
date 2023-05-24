package command

import (
	"errors"
	"goproject/constants"
	"goproject/game/data"
	"goproject/structure"
	"math/rand"
	"time"
)

func addItemToInventory(inventory *structure.Inventory, itemType constants.ItemType) {
	(*inventory)[itemType] += 1
}

func removeItemInInventory(inventory *structure.Inventory, itemType constants.ItemType) {
	(*inventory)[itemType] -= 1
}

func hasItemInInventory(inventory *structure.Inventory, itemType constants.ItemType) bool {
	return (*inventory)[itemType] > 0
}

func isUsableItem(itemType constants.ItemType) bool {
	return data.ItemTypeUsableMap[itemType]
}

func UseItemByName(status *structure.Status, itemName string) {
	itemType := constants.StringItemTypeMap[itemName]

	switch itemType {
	case constants.HealPotion:
		status.Attribute.Health += 30
		removeItemInInventory(status.Inventory, constants.HealPotion)
	default:
		return
	}
}

func UseItemToDoorByName(room *structure.Room, inventory *structure.Inventory, itemName string, doorName string) {
	itemType := constants.StringItemTypeMap[itemName]

	if constants.Hammer == itemType {
		if constants.GlassDoor == constants.StringDoorTypeMap[doorName] {
			breakGlassDoor(room, inventory)
			removeItemInInventory(inventory, constants.Hammer)
			return
		}
	}

	if constants.Key == itemType {
		if constants.LockedDoor == constants.StringDoorTypeMap[doorName] {
			unlockLockedDoor(room, inventory)
			removeItemInInventory(inventory, constants.Key)
			return
		}
	}
}

func GetItemByPercentage(dropItems *[]structure.DropItem) (constants.ItemType, int) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Float64()
	totalProbability := 0.0

	for _, dropItem := range *dropItems {
		totalProbability += dropItem.DropPercentage.Percentage
		if totalProbability >= randomNum {
			return dropItem.ItemType, dropItem.DropPercentage.Num
		}
	}
	return constants.Nothing, 0
}

func ValidateItemUsability(inventory *structure.Inventory, itemType constants.ItemType) error {
	if hasItemInInventory(inventory, itemType) {
		return errors.New(constants.NoItemInInventory)
	} else if isUsableItem(itemType) {
		return errors.New(constants.CanNotUseSuchItem)
	}
	return nil
}
