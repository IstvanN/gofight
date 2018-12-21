package classes

import "fmt"

type druid struct {
	name                      string
	maxhp, hp, ap, armor, wis int
	stunned                   bool
}

func CreateDruid(name string) *druid {
	d := &druid{
		name: name,
		maxhp: 50 + Rolld(10),
		wis:   5 + Rolld(6),
	}

	d.ap += 10 + 2*d.wis
	d.armor += Rolld(6) + d.wis
	d.hp = d.maxhp
	return d
	}
}