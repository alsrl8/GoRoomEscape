package command

import (
	"fmt"
	"goproject/constants"
	"goproject/structure"
)

func getEquipmentList(status structure.Status) []constants.ItemType {
	return []constants.ItemType{
		status.Equipment.Top,
		status.Equipment.Pants,
		status.Equipment.Shoes,
		status.Equipment.LeftHand,
		status.Equipment.RightHand,
	}
}

func ShowEquipment(status structure.Status) {
	equipmentList := getEquipmentList(status)
	if len(equipmentList) != len(constants.BodyPartList) {
		panic("Length of equipment list and body part list must be same")
	}
	for i, equipment := range equipmentList {
		bodyPartName := constants.BodyPartStringMap[constants.BodyPartList[i]]
		fmt.Println(bodyPartName + ": " + constants.ItemTypeStringMap[equipment])
	}
}

func isBodyPartsEmpty(status *structure.Status, part constants.BodyPart) bool {
	equippedItem := constants.Nothing
	switch part {
	case constants.Top:
		equippedItem = status.Equipment.Top
	case constants.Pants:
		equippedItem = status.Equipment.Pants
	case constants.Shoes:
		equippedItem = status.Equipment.Shoes
	case constants.LeftHand:
		equippedItem = status.Equipment.LeftHand
	case constants.RightHand:
		equippedItem = status.Equipment.RightHand
	}
	return equippedItem == constants.Nothing
}

func isWearableItem(itemName string) bool {
	item := constants.StringItemTypeMap[itemName]
	wearable := constants.ItemTypeWearableMap[item]
	return wearable
}

func isBodyPartToWearExist(status *structure.Status, itemType constants.ItemType) bool {
	bodyParts := constants.ItemTypeBodyPartMap[itemType]
	if len(bodyParts) == 0 {
		return false
	}

	var emptyBodyParts []constants.BodyPart
	for _, part := range bodyParts {
		if !isBodyPartsEmpty(status, part) {
			continue
		}
		emptyBodyParts = append(emptyBodyParts, part)
	}
	if len(emptyBodyParts) == 0 {
		return false
	}

	return true
}

func Equip(status *structure.Status, itemName string) {
	wearable := isWearableItem(itemName)
	itemType := constants.StringItemTypeMap[itemName]
	if (*status.Inventory)[itemType] == 0 {
		fmt.Println(constants.NoItemInInventory)
		return
	} else if !wearable {
		fmt.Println(constants.CanNotWear)
		return
	}
	validBodyPartExist := isBodyPartToWearExist(status, itemType)
	if !validBodyPartExist {
		fmt.Println(constants.NoBodyPartToWear)
		return
	}

	for _, bodyPart := range constants.ItemTypeBodyPartMap[itemType] {
		if !isBodyPartsEmpty(status, bodyPart) {
			continue
		}
		setEquipmentToBodyPart(status, bodyPart, itemType)
		return
	}
}

func setEquipmentToBodyPart(status *structure.Status, bodyPart constants.BodyPart, itemType constants.ItemType) {
	switch bodyPart {
	case constants.Top:
		status.Equipment.Top = itemType
	case constants.Pants:
		status.Equipment.Pants = itemType
	case constants.Shoes:
		status.Equipment.Shoes = itemType
	case constants.LeftHand:
		status.Equipment.LeftHand = itemType
	case constants.RightHand:
		status.Equipment.RightHand = itemType
	}
}
