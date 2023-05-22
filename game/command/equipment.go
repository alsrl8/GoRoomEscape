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
	itemOnBody := getEquipmentPartByBodyPart(status.Equipment, part)
	return *itemOnBody == constants.Nothing
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
		removeItemInInventory(status.Inventory, itemType)
		return
	}
}

func setEquipmentToBodyPart(status *structure.Status, bodyPart constants.BodyPart, itemType constants.ItemType) {
	equipmentPart := getEquipmentPartByBodyPart(status.Equipment, bodyPart)
	*equipmentPart = itemType
}

func Disarm(status *structure.Status, bodyPartName string) {
	bodyPart := constants.StringBodyPartMap[bodyPartName]
	equipmentPart := getEquipmentPartByBodyPart(status.Equipment, bodyPart)

	if *equipmentPart == constants.Nothing {
		fmt.Println(constants.NoEquipmentOnBodyPart)
		return
	}

	item := *equipmentPart
	*equipmentPart = constants.Nothing
	(*status.Inventory)[item] += 1
}

func getEquipmentPartByBodyPart(equipment *structure.Equipment, bodyPart constants.BodyPart) *constants.ItemType {
	switch bodyPart {
	case constants.Top:
		return &equipment.Top
	case constants.Pants:
		return &equipment.Pants
	case constants.Shoes:
		return &equipment.Shoes
	case constants.LeftHand:
		return &equipment.LeftHand
	case constants.RightHand:
		return &equipment.RightHand
	}
	return nil
}
