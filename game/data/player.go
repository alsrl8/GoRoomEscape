package data

import "goproject/structure"

func GetAttribute() structure.Attribute {
	return structure.Attribute{
		Health:  50,
		Attack:  100,
		Defense: 3,
	}
}
