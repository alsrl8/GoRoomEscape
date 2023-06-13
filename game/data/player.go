package data

import "goproject/structure"

var RunSuccessPercentage = 0.5
var MaxHealth = 50

func GetAttribute() *structure.Attribute {
	return &structure.Attribute{
		Health:  MaxHealth,
		Attack:  3,
		Defense: 3,
	}
}
