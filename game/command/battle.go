package command

import (
	"goproject/structure"
	"math"
)

func reduceHealth(attribute *structure.Attribute, attack int) int {
	damage := int(math.Max(0, float64(attack-attribute.Defense)))
	attribute.Health -= damage
	return damage
}

func IsDead(attribute structure.Attribute) bool {
	return attribute.Health <= 0
}
