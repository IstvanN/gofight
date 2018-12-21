package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to GoFight!")
	var f1 fighter = createRogue("Wolfgraff")
	var f2 fighter = createRogue("Grunt")
	fmt.Println(f1, f2)
	f1.act(f2)
	f2.act(f1)
	fmt.Println(f1, f2)
}

func rolld(dice int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(dice) + 1
	return r
}

type fighter interface {
	act(fighter)
	suffer(dmg int)
	heal()
}

type druid struct {
	name                      string
	maxhp, hp, ap, armor, wis int
}

func createDruid(name string) *druid {
	d := &druid{
		name:  name,
		maxhp: 50 + rolld(10),
		wis:   5 + rolld(6),
	}
	d.ap += rolld(10) + d.wis
	d.armor += rolld(10) + d.wis
	return d
}

func (d *druid) act(f fighter) {

}

func (d *druid) smite(f fighter) {
	dmg := d.ap + rolld(10)
	f.suffer(dmg)
}

func (d *druid) suffer(dmg int) {
	d.hp -= dmg
}
