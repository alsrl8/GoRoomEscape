package command

import (
	"goproject/structure"
	"math"
)

func reduceHealth(attribute *structure.Attribute, attack int) {
	attribute.Health -= int(math.Max(0, float64(attack-attribute.Defense)))
}

func IsDead(attribute structure.Attribute) bool {
	return attribute.Health <= 0
}
