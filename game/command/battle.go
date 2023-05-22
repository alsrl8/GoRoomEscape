package command

import (
	"goproject/structure"
	"math"
)

func reduceHealth(attribute *structure.Attribute, attack int) {
	attribute.Health -= int(math.Max(0, float64(attribute.Defense-attack)))
}
