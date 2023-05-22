package command

import (
	"fmt"
	"goproject/constants"
	"goproject/structure"
	"math/rand"
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

func UseItem(itemName string) {
	// TODO 아이템 사용
}

func UseItemToDoor(room *structure.Room, inventory *structure.Inventory, itemName string, doorName string) {
	if !hasItemInInventory(*inventory, constants.StringItemTypeMap[itemName]) {
		fmt.Println(constants.NoSuchItem, itemName)
	} else if findDoorByName(room, doorName) == nil {
		fmt.Println(constants.NoSuchDoor, doorName)
	}

	if constants.Hammer == constants.StringItemTypeMap[itemName] {
		if constants.GlassDoor == constants.StringDoorTypeMap[doorName] {
			breakGlassDoorAndReduceHammer(room, inventory)
			return
		}
	}

	if constants.Key == constants.StringItemTypeMap[itemName] {
		if constants.LockedDoor == constants.StringDoorTypeMap[doorName] {
			unlockLockedDoorAndReduceKey(room, inventory)
			return
		}
	}
}

func GetItemByPercentage(dropItems *[]structure.DropItem) (constants.ItemType, int) {
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
