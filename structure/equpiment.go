package structure

type Armor struct {
	Item
}

type BodyPartForArmor struct {
	Top   Armor
	Pants Armor
	Shoes Armor
}

type Weapon struct {
	Item
}

type BodyPartForWeapon struct {
	LeftHand  *Weapon
	RightHand *Weapon
}
