package command

import (
	"errors"
	"fmt"
	"goproject/constants"
	"goproject/game/data"
	"goproject/structure"
	"math"
	"math/rand"
	"time"
)

func addItemToInventory(inventory *structure.Inventory, itemType constants.ItemType, itemNum int) {
	(*inventory)[itemType] += itemNum
}

func removeItemInInventory(inventory *structure.Inventory, itemType constants.ItemType, itemNum int) {
	(*inventory)[itemType] -= itemNum
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
		healPoint := 30
		status.Attribute.Health = int(math.Min(float64(data.MaxHealth), float64(status.Attribute.Health+healPoint)))
		removeItemInInventory(status.Inventory, constants.HealPotion, 1)
	case constants.Box:
		boxItemType, boxItemNum := GetItemByPercentage(&data.BoxDropItems)
		addItemToInventory(status.Inventory, boxItemType, boxItemNum)
		removeItemInInventory(status.Inventory, constants.Box, 1)
		fmt.Printf(constants.GetItem, constants.ItemTypeStringMap[boxItemType], boxItemNum)
	default:
		return
	}
}

func UseItemToDoorByName(room *structure.Room, inventory *structure.Inventory, itemName string, doorName string) {
	itemType := constants.StringItemTypeMap[itemName]

	if constants.Hammer == itemType {
		if constants.GlassDoor == constants.StringDoorTypeMap[doorName] {
			breakGlassDoor(room, inventory)
			removeItemInInventory(inventory, constants.Hammer, 1)
			return
		}
	}

	if constants.Key == itemType {
		if constants.LockedDoor == constants.StringDoorTypeMap[doorName] {
			unlockLockedDoor(room, inventory)
			removeItemInInventory(inventory, constants.Key, 1)
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

func ValidateItemUsability(inventory *structure.Inventory, itemType constants.ItemType, includeTargetFlag bool) error {
	if !hasItemInInventory(inventory, itemType) {
		return errors.New(constants.NoItemInInventory)
	} else if !isUsableItem(itemType) {
		return errors.New(constants.CanNotUseSuchItem)
	} else if !includeTargetFlag && data.ItemTypeTargetNeededMap[itemType] {
		return errors.New(constants.NoSpecificTargetForItem)
	}
	return nil
}
