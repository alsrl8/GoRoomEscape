package command

import (
	"fmt"
	"goproject/constants"
	"goproject/game/data"
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
	wearable := data.ItemTypeWearableMap[item]
	return wearable
}

func isBodyPartToWearExist(status *structure.Status, itemType constants.ItemType) bool {
	bodyParts := data.ItemTypeBodyPartMap[itemType]
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
	itemType := constants.StringItemTypeMap[itemName]
	if (*status.Inventory)[itemType] == 0 {
		fmt.Println(constants.NoItemInInventory, itemName)
		return
	} else if !isWearableItem(itemName) {
		fmt.Println(constants.CanNotWear, itemName)
		return
	} else if !isBodyPartToWearExist(status, itemType) {
		fmt.Println(constants.NoBodyPartToWear)
		return
	}

	for _, bodyPart := range data.ItemTypeBodyPartMap[itemType] {
		if !isBodyPartsEmpty(status, bodyPart) {
			continue
		}
		setEquipmentToBodyPart(status, bodyPart, itemType)
		applyEquipmentEffect(status, itemType)
		removeItemInInventory(status.Inventory, itemType)
		fmt.Printf("%s에 %s 장비\n", constants.BodyPartStringMap[bodyPart], constants.ItemTypeStringMap[itemType])
		return
	}
}

func setEquipmentToBodyPart(status *structure.Status, bodyPart constants.BodyPart, itemType constants.ItemType) {
	equipmentPart := getEquipmentPartByBodyPart(status.Equipment, bodyPart)
	*equipmentPart = itemType
}

func applyEquipmentEffect(status *structure.Status, itemType constants.ItemType) {
	status.Attribute.Health += data.EquipmentEffectMap[itemType].Health
	status.Attribute.Attack += data.EquipmentEffectMap[itemType].Attack
	status.Attribute.Defense += data.EquipmentEffectMap[itemType].Defense
}

func removeEquipmentEffect(status *structure.Status, itemType constants.ItemType) {
	status.Attribute.Attack -= data.EquipmentEffectMap[itemType].Attack
	status.Attribute.Health -= data.EquipmentEffectMap[itemType].Health
	status.Attribute.Defense -= data.EquipmentEffectMap[itemType].Defense
}

func Disarm(status *structure.Status, bodyPartName string) {
	bodyPart := constants.StringBodyPartMap[bodyPartName]
	equipmentPart := getEquipmentPartByBodyPart(status.Equipment, bodyPart)

	if *equipmentPart == constants.Nothing {
		fmt.Println(constants.NoEquipmentOnBodyPart, bodyPartName)
		return
	}

	itemType := *equipmentPart
	*equipmentPart = constants.Nothing
	addItemToInventory(status.Inventory, itemType)
	removeEquipmentEffect(status, itemType)
	fmt.Printf("%s에서 장비(%s)를 해제했습니다.\n", bodyPartName, constants.ItemTypeStringMap[itemType])
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

func IsAttackAble(status *structure.Status) bool {
	if status.GuardFlag {
		return false
	}
	return true
}

func IsGuardAble(status *structure.Status) bool {
	if status.GuardFlag {
		return false
	} else if !isSpecificItemEquipped(status.Equipment, constants.WoodShield) {
		return false
	}
	return true
}

func isSpecificItemEquipped(equipment *structure.Equipment, itemType constants.ItemType) bool {
	for _, bodyPart := range constants.BodyPartList {
		equipmentOnBody := getEquipmentPartByBodyPart(equipment, bodyPart)
		if *equipmentOnBody == itemType {
			return true
		}
	}
	return false
}
