package classes

import "fmt"

type rogue struct {
	name                      string
	maxhp, hp, ap, armor, dex int
}

// CreateRogue is the constructor of the Rogue class
func CreateRogue(name string) *rogue {
	r := &rogue{
		name:  name,
		maxhp: 50 + Rolld(10),
		dex:   5 + Rolld(6),
	}
	r.ap += Rolld(10) + r.dex
	r.armor += Rolld(6) + r.dex
	r.hp = r.maxhp
	return r
}

func (r rogue) GetName() string {
	return r.name
}

func (r *rogue) Act(f Fighter) {
	if !r.IsDead() {
		if Rolld(2) == 1 {
			r.Backstab(f)
		} else {
			r.Heal()
		}
	}
}

func (r rogue) Backstab(f Fighter) {
	dmg := r.ap + Rolld(10)
	fmt.Printf("%v is savagely backstabbing his opponent. ", r.name)
	f.Suffer(dmg)
}

func (r *rogue) Heal() {
	healing := Rolld(10)

	if r.hp+healing >= r.maxhp {
		fmt.Printf("%v is popping up a healing potion, healing to max!\n", r.name)
		r.hp = r.maxhp
	} else {
		r.hp += healing
		fmt.Printf("%v is popping up a healing potion, healing %d health points.\n", r.name, healing)
	}

}

func (r *rogue) Suffer(dmg int) {
	actualDmg := dmg - r.armor
	if actualDmg <= 0 {
		fmt.Printf("Thanks to his armour, %v doesn't take any dmg!\n", r.name)
	} else {
		fmt.Printf("%v has taken %d damage.\n", r.name, actualDmg)
		r.hp -= actualDmg
	}
}

func (r rogue) IsDead() bool {
	return r.hp <= 0
}
