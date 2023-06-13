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

func Equip(status *structure.Status, itemType constants.ItemType) {
	if !hasItemInInventory(status.Inventory, itemType) {
		fmt.Println(constants.NoItemInInventory, constants.ItemTypeStringMap[itemType])
		return
	}

	wearableBodyParts := getWearableBodyParts(status, itemType)
	if len(wearableBodyParts) == 0 {
		fmt.Println(constants.CanNotWear, constants.ItemTypeStringMap[itemType])
		return
	}

	wearableBodyParts[0].Equip(itemType) // 착용할 수 있는 부위 중 첫 번째 부위에 착용한다.
	applyEquipmentEffect(status, itemType)
	removeItemInInventory(status.Inventory, itemType, 1)
}

func Disarm(status *structure.Status, itemType constants.ItemType) {
	if itemType == constants.Nothing {
		fmt.Println(constants.NoSuchItem, constants.ItemTypeStringMap[itemType])
		return
	}

	bodyPartsWithEquipment := getBodyPartsWhereEquipmentIsWorn(status, itemType)
	if len(bodyPartsWithEquipment) == 0 {
		fmt.Println(constants.NoSuchEquipmentOnBody)
		return
	}

	bodyPartsWithEquipment[0].Disarm()
	removeEquipmentEffect(status, itemType)
	addItemToInventory(status.Inventory, itemType, 1)
}

func showBodyPartsForArmors(bodyParts *structure.BodyPartForArmor) {
	t := reflect.TypeOf(*bodyParts)
	v := reflect.ValueOf(*bodyParts)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		armor := v.Field(i).Interface().(*structure.Armor)
		fmt.Printf("%-10s : %s\n", field.Name, constants.ItemTypeStringMap[armor.ItemType])
	}
}

func showBodyPartsForWeapons(bodyParts *structure.BodyPartForWeapon) {
	t := reflect.TypeOf(*bodyParts)
	v := reflect.ValueOf(*bodyParts)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		weapon := v.Field(i).Interface().(*structure.Weapon)
		fmt.Printf("%-10s : %s\n", field.Name, constants.ItemTypeStringMap[weapon.ItemType])
	}
}

func isWearableItem(wearable structure.Wearable, itemType constants.ItemType) bool {
	return wearable.IsWearable(itemType)
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

func getBodyPartsWhereEquipmentIsWorn(status *structure.Status, itemType constants.ItemType) (ret []structure.Wearable) {
	bodyPartForArmor := status.BodyPartForArmor
	t := reflect.TypeOf(bodyPartForArmor)
	v := reflect.ValueOf(bodyPartForArmor)
	for i := 0; i < t.NumField(); i++ {
		armor := v.Field(i).Interface().(*structure.Armor)
		if armor.ItemType != itemType {
			continue
		}
		ret = append(ret, armor)
	}
	bodyPartForWeapon := status.BodyPartForWeapon
	t = reflect.TypeOf(bodyPartForWeapon)
	v = reflect.ValueOf(bodyPartForWeapon)
	for i := 0; i < t.NumField(); i++ {
		weapon := v.Field(i).Interface().(*structure.Weapon)
		if weapon.ItemType != itemType {
			continue
		}
		ret = append(ret, weapon)
	}
	return
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
