package command

import (
	"fmt"
	"goproject/constants"
	"goproject/structure"
)

var itemTypeWearableMap = map[constants.ItemType]bool{
	constants.WoodSword:    true,
	constants.IronSword:    true,
	constants.WoodShield:   true,
	constants.LeatherHat:   true,
	constants.LeatherCloth: true,
	constants.LeatherPants: true,
	constants.LeatherShoes: true,
}

var equipmentEffectMap = map[constants.ItemType]structure.Attribute{
	constants.WoodSword: {
		Health:  0,
		Attack:  3,
		Defense: 0,
	},
	constants.IronSword: {
		Health:  0,
		Attack:  3,
		Defense: 0,
	},
	constants.WoodShield: {
		Health:  0,
		Attack:  0,
		Defense: 0,
	},
	constants.LeatherHat: {
		Health:  0,
		Attack:  0,
		Defense: 0,
	},
	constants.LeatherCloth: {
		Health:  0,
		Attack:  0,
		Defense: 6,
	},
	constants.LeatherPants: {
		Health:  0,
		Attack:  0,
		Defense: 4,
	},
	constants.LeatherShoes: {
		Health:  0,
		Attack:  0,
		Defense: 3,
	},
}

var itemTypeBodyPartMap = map[constants.ItemType][]constants.BodyPart{
	constants.WoodSword:    {constants.LeftHand, constants.RightHand},
	constants.IronSword:    {constants.LeftHand, constants.RightHand},
	constants.WoodShield:   {constants.LeftHand, constants.RightHand},
	constants.LeatherCloth: {constants.Top},
	constants.LeatherPants: {constants.Pants},
	constants.LeatherShoes: {constants.Shoes},
}

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
	wearable := itemTypeWearableMap[item]
	return wearable
}

func isBodyPartToWearExist(status *structure.Status, itemType constants.ItemType) bool {
	bodyParts := itemTypeBodyPartMap[itemType]
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

	for _, bodyPart := range itemTypeBodyPartMap[itemType] {
		if !isBodyPartsEmpty(status, bodyPart) {
			continue
		}
		setEquipmentToBodyPart(status, bodyPart, itemType)
		applyEquipmentEffect(status, itemType)
		removeItemInInventory(status.Inventory, itemType)
		fmt.Printf("%s에 장비(%s)를 장비했습니다.\n", constants.BodyPartStringMap[bodyPart], constants.ItemTypeStringMap[itemType])
		return
	}
}

func setEquipmentToBodyPart(status *structure.Status, bodyPart constants.BodyPart, itemType constants.ItemType) {
	equipmentPart := getEquipmentPartByBodyPart(status.Equipment, bodyPart)
	*equipmentPart = itemType
}

func applyEquipmentEffect(status *structure.Status, itemType constants.ItemType) {
	status.Attribute.Health += equipmentEffectMap[itemType].Health
	status.Attribute.Attack += equipmentEffectMap[itemType].Attack
	status.Attribute.Defense += equipmentEffectMap[itemType].Defense
}

func removeEquipmentEffect(status *structure.Status, itemType constants.ItemType) {
	status.Attribute.Health -= equipmentEffectMap[itemType].Health
	status.Attribute.Attack -= equipmentEffectMap[itemType].Attack
	status.Attribute.Defense -= equipmentEffectMap[itemType].Defense
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
