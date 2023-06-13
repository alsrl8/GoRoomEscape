package command

import (
	"errors"
	"fmt"
	"goproject/constants"
	"goproject/game/data"
	"goproject/structure"
	"math"
)

func addItemToInventory(inventory structure.Inventory, itemType constants.ItemType, itemNum int) {
	inventory[itemType] += itemNum
}

func removeItemInInventory(inventory structure.Inventory, itemType constants.ItemType, itemNum int) {
	inventory[itemType] -= itemNum
}

func hasItemInInventory(inventory structure.Inventory, itemType constants.ItemType, itemNum int) bool {
	return inventory[itemType] > itemNum
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
		var dropItems structure.DropItemSlice = data.BoxDropItems
		boxItemType, boxItemNum := dropItems.GetItemByPercentage()
		addItemToInventory(status.Inventory, boxItemType, boxItemNum)
		removeItemInInventory(status.Inventory, constants.Box, 1)
		fmt.Printf(constants.GetItem, constants.ItemTypeStringMap[boxItemType], boxItemNum)
	default:
		return
	}
}

func UseItemToDoorByName(room *structure.Room, inventory structure.Inventory, itemName string, doorName string) {
	itemType := constants.StringItemTypeMap[itemName]

	if constants.Hammer == itemType {
		if constants.GlassDoor == constants.StringDoorTypeMap[doorName] {
			breakGlassDoor(room)
			removeItemInInventory(inventory, constants.Hammer, 1)
			return
		}
	}

	if constants.Key == itemType {
		if constants.LockedDoor == constants.StringDoorTypeMap[doorName] {
			unlockLockedDoor(room)
			removeItemInInventory(inventory, constants.Key, 1)
			return
		}
	}
}

func ValidateItemExist(inventory structure.Inventory, itemType constants.ItemType, itemNum int) error {
	if !hasItemInInventory(inventory, itemType, itemNum) {
		return errors.New(constants.NoItemInInventory)
	}
	return nil
}

func ValidateItemUsability(inventory structure.Inventory, itemType constants.ItemType, includeTargetFlag bool) error {
	if !hasItemInInventory(inventory, itemType, 1) {
		return errors.New(constants.NoItemInInventory)
	} else if !isUsableItem(itemType) {
		return errors.New(constants.CanNotUseSuchItem)
	} else if !includeTargetFlag && data.ItemTypeTargetNeededMap[itemType] {
		return errors.New(constants.NoSpecificTargetForItem)
	}
	return nil
}

func ValidateItemDoorMatch(itemType constants.ItemType, doorType constants.DoorType) error {
	switch itemType {
	case constants.Hammer:
		if doorType == constants.GlassDoor {
			return nil
		} else {
			return errors.New(constants.NoMatchItemAndDoor)
		}
	case constants.Key:
		if doorType == constants.LockedDoor {
			return nil
		} else {
			return errors.New(constants.NoMatchItemAndDoor)
		}
	default:
		return errors.New(constants.NoMatchItemAndDoor)
	}
}

func PickUpItems(status *structure.Status, itemType constants.ItemType, itemNum int) {
	room := GetCurrentRoom(status)
	_itemNum := room.Items[itemType]
	if _itemNum == 0 {
		fmt.Println(constants.NoSuchItem, constants.ItemTypeStringMap[itemType])
		return
	} else if _itemNum < itemNum {
		fmt.Println(constants.NotEnoughItem, constants.ItemTypeStringMap[itemType])
		return
	}

	inventory := status.Inventory
	inventory[itemType] += itemNum
	room.Items[itemType] -= itemNum
}

func DropItems(status *structure.Status, itemType constants.ItemType, itemNum int) {
	inventory := status.Inventory
	_itemNum := inventory[itemType]
	if _itemNum == 0 {
		fmt.Println(constants.NoItemInInventory, constants.ItemTypeStringMap[itemType])
		return
	} else if _itemNum < itemNum {
		fmt.Println(constants.NotEnoughItem, constants.ItemTypeStringMap[itemType])
		return
	}

	room := GetCurrentRoom(status)
	room.Items[itemType] += itemNum
	inventory[itemType] -= itemNum
}

func DiscardItem(inventory structure.Inventory, itemType constants.ItemType, itemNum int) {
	inventory[itemType] -= itemNum
}
