package structure

type BodyPartForArmor struct {
	Top   *Armor
	Pants *Armor
	Shoes *Armor
}

type BodyPartForWeapon struct {
	LeftHand  *Weapon
	RightHand *Weapon
}
