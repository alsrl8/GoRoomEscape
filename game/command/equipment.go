package command

import (
	"fmt"
	"goproject/constants"
	"goproject/game/data"
	"goproject/structure"
	"reflect"
)

func ShowBodyParts(status structure.Status) {
	showBodyPartsForArmors(status.BodyPartForArmor)
	showBodyPartsForWeapons(status.BodyPartForWeapon)
}

func showBodyPartsForArmors(bodyParts structure.BodyPartForArmor) {
	t := reflect.TypeOf(bodyParts)
	v := reflect.ValueOf(bodyParts)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		armor := v.Field(i).Interface().(structure.Armor)
		fmt.Printf("%10s : %s\n", field.Name, constants.ItemTypeStringMap[armor.ItemType])
	}
}

func showBodyPartsForWeapons(bodyParts structure.BodyPartForWeapon) {
	t := reflect.TypeOf(bodyParts)
	v := reflect.ValueOf(bodyParts)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		weapon := v.Field(i).Interface().(structure.Weapon)
		fmt.Printf("%10s : %s\n", field.Name, constants.ItemTypeStringMap[weapon.ItemType])
	}
}

func isWearableItem(wearable structure.Wearable, itemType constants.ItemType) bool {
	return wearable.IsWearable(itemType)
}

func Equip(status *structure.Status, itemName string) {
	itemType := constants.StringItemTypeMap[itemName]
	if !hasItemInInventory(status.Inventory, itemType) {
		fmt.Println(constants.NoItemInInventory, itemName)
		return
	}
	//else if !isWearableItem(itemType) {
	//	fmt.Println(constants.CanNotWear, itemName)
	//	return
	//}

	//parts := data.ItemTypeBodyPartMap[itemType]
	//for _ := range parts {
	//	status.LeftHand.Equip(itemType)
	//}
	//
	//for _, bodyPart := range data.ItemTypeBodyPartMap[itemType] {
	//	setEquipmentToBodyPart(status, bodyPart, itemType)
	//	applyEquipmentEffect(status, itemType)
	//	removeItemInInventory(status.Inventory, itemType, 1)
	//	fmt.Printf(constants.EquipEquipment, constants.BodyPartStringMap[bodyPart], constants.ItemTypeStringMap[itemType])
	//	return
	//}

	wearableBodyParts := getWearableBodyParts(status, itemType)
	fmt.Println(wearableBodyParts)
}

func getWearableBodyParts(status *structure.Status, itemType constants.ItemType) (ret []structure.Wearable) {
	bodyPartForArmor := status.BodyPartForArmor
	t := reflect.TypeOf(bodyPartForArmor)
	v := reflect.ValueOf(bodyPartForArmor)
	for i := 0; i < t.NumField(); i++ {
		wearable := v.Field(i).Interface().(structure.Wearable)
		if !isWearableItem(wearable, itemType) {
			continue
		}
		ret = append(ret, wearable)
	}
	bodyPartForWeapon := status.BodyPartForWeapon
	t = reflect.TypeOf(bodyPartForWeapon)
	v = reflect.ValueOf(bodyPartForWeapon)
	for i := 0; i < t.NumField(); i++ {
		wearable := v.Field(i).Interface().(structure.Wearable)
		if !isWearableItem(wearable, itemType) {
			continue
		}
		ret = append(ret, wearable)
	}
	return
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
	addItemToInventory(status.Inventory, itemType, 1)
	removeEquipmentEffect(status, itemType)
	fmt.Printf(constants.DisarmEquipment, bodyPartName, constants.ItemTypeStringMap[itemType])
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
