package classes

import (
	"fmt"
)

type druid struct {
	name                      string
	maxhp, hp, ap, armor, wis int
	stunned                   bool
}

// CreateDruid is the constructor of the druid struct
func CreateDruid(name string) *druid {
	d := &druid{
		name:  name,
		maxhp: 50 + Rolld(10),
		wis:   5 + Rolld(6),
	}
	d.ap += 10 + 2*d.wis
	d.armor += Rolld(6) + d.wis
	d.hp = d.maxhp
	return d
}

func (d druid) GetName() string {
	return d.name
}

func (d druid) GetHP() int {
	return d.hp
}

func (d druid) IsDead() bool {
	return d.hp <= 0
}

func (d druid) IsStunned() bool {
	return d.stunned
}

func (d *druid) BeStunned() {
	d.stunned = true
}

func (d *druid) Suffer(dmg int) {
	actualDmg := dmg - d.armor
	if actualDmg <= 0 {
		fmt.Printf("Thanks to his armour, %s doesn't take any dmg!\n", d.name)
	} else {
		fmt.Printf("%s has taken %d damage.\n", d.name, actualDmg)
		d.hp -= actualDmg
	}
}

func (d *druid) Act(f Fighter) {
	if !d.IsDead() && !d.stunned {
		roll := Rolld(4)
		switch roll {
		case 1, 2:
			d.moonbean(f)
		case 3:
			d.heal(f)
		case 4:
			d.thorns(f)
		}
	}
	if !d.IsDead() && d.stunned {
		fmt.Printf("%s is stunned, thus not acting this round.\n", d.name)
		d.stunned = false
	}
}

func (d druid) moonbean(f Fighter) {
	dmg := d.ap
	fmt.Printf("%s summons burning moonbean at his enemy. ", d.name)
	f.Suffer(dmg)
}

func (d druid) thorns(f Fighter) {
	fmt.Printf("%s grabs his enemy with thorns, completely stunning him.", d.name)
	f.BeStunned()
}

func (d *druid) heal(f Fighter) {
	healing := Rolld(10) + d.wis
	if f.IsStunned() {
		healing *= 2
	}

	if d.hp+healing >= d.maxhp {
		fmt.Printf("%s is blessed by the divines, healing him to the max!\n", d.name)
		d.hp = d.maxhp
	} else {
		d.hp += healing
		fmt.Printf("%s is blessed by the divines, healing him %d health points.\n", d.name, healing)
	}
}
