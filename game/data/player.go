package data

import "goproject/structure"

func GetAttribute() structure.Attribute {
	return structure.Attribute{
		Health:  50,
		Attack:  3,
		Defense: 3,
	}
}
