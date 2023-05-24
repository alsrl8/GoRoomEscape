package command

import (
	"fmt"
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

func hasItemInInventory(inventory structure.Inventory, itemType constants.ItemType) bool {
	return inventory[itemType] > 0
}

func isUsableItem(itemType constants.ItemType) bool {
	return data.ItemTypeUsableMap[itemType]
}

// TODO UseItem에서 item validation 분리
func UseItem(inventory *structure.Inventory, itemName string) {
	itemType := constants.StringItemTypeMap[itemName]
	if !hasItemInInventory(*inventory, itemType) {
		fmt.Println(constants.NoSuchItem, itemName)
	} else if !isUsableItem(itemType) {
		fmt.Println(constants.CanNotUseSuchItem, itemName)
	}

	// TODO UseItem 상세 로직 추가
}

// TODO UseItemToDoor Validation 분리
func UseItemToDoor(room *structure.Room, inventory *structure.Inventory, itemName string, doorName string) {
	itemType := constants.StringItemTypeMap[itemName]
	if !hasItemInInventory(*inventory, itemType) {
		fmt.Println(constants.NoSuchItem, itemName)
	} else if !isUsableItem(itemType) {
		fmt.Println(constants.CanNotUseSuchItem, itemName)
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
