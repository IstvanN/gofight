package main

import (
	"fmt"

	"github.com/IstvanN/gofight/classes"
)

func main() {
	fmt.Println("Welcome to GoFight!")
	var f1 classes.Fighter = classes.CreateRogue("Wolfgraff")
	var f2 classes.Fighter = classes.CreateRogue("Grunt")
	fmt.Println(f1, f2)
	f1.Act(f2)
	f2.Act(f1)
	fmt.Println(f1, f2)
}

type druid struct {
	name                      string
	maxhp, hp, ap, armor, wis int
}

func createDruid(name string) *druid {
	d := &druid{
		name:  name,
		maxhp: 50 + Rolld(10),
		wis:   5 + Rolld(6),
	}
	d.ap += Rolld(10) + d.wis
	d.armor += Rolld(10) + d.wis
	return d
}

func (d *druid) Act(f Fighter) {

}

func (d *druid) smite(f Fighter) {
	dmg := d.ap + Rolld(10)
	f.Suffer(dmg)
}

func (d *druid) Suffer(dmg int) {
	d.hp -= dmg
}
