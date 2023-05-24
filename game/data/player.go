package data

import "goproject/structure"

var RunSuccessPercentage = 0.5

func GetAttribute() structure.Attribute {
	return structure.Attribute{
		Health:  50,
		Attack:  3,
		Defense: 3,
	}
}
