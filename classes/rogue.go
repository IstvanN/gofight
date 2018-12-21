package classes

import "fmt"

type rogue struct {
	name                      string
	maxhp, hp, ap, armor, dex int
	stunned                   bool
}

// CreateRogue is the constructor of the Rogue struct
func CreateRogue(name string) *rogue {
	r := &rogue{
		name:  name,
		maxhp: 50 + Rolld(10),
		dex:   5 + Rolld(6),
	}
	r.ap += 10 + r.dex
	r.armor += Rolld(6) + r.dex
	r.hp = r.maxhp
	return r
}

func (r rogue) GetName() string {
	return r.name
}

func (r rogue) GetHP() int {
	return r.hp
}

func (r *rogue) Act(f Fighter) {
	if !r.IsDead() && !r.stunned {
		roll := Rolld(4)
		switch roll {
		case 1, 2:
			r.backstab(f)
		case 3:
			r.heal()
		case 4:
			r.dirt(f)
		}
	}
	if !r.IsDead() && r.stunned {
		fmt.Printf("%s is stunned, thus not acting this round.\n", r.name)
		r.stunned = false
	}
}

func (r *rogue) BeStunned() {
	r.stunned = true
}

func (r rogue) IsStunned() bool {
	return r.stunned
}

func (r *rogue) Suffer(dmg int) {
	actualDmg := dmg - r.armor
	if actualDmg <= 0 {
		fmt.Printf("Thanks to his armour, %s doesn't take any dmg!\n", r.name)
	} else {
		fmt.Printf("%s has taken %d damage.\n", r.name, actualDmg)
		r.hp -= actualDmg
	}
}

func (r rogue) IsDead() bool {
	return r.hp <= 0
}

func (r rogue) backstab(f Fighter) {
	dmg := r.ap + Rolld(10)
	if f.IsStunned() {
		dmg *= 2
	}
	fmt.Printf("%s is savagely backstabbing his opponent. ", r.name)
	f.Suffer(dmg)
}

func (r rogue) dirt(f Fighter) {
	fmt.Printf("%s throws dirt in the very eyes of %s, stunning him completely. Nasty move!\n", r.name, f.GetName())
	f.BeStunned()
}

func (r *rogue) heal() {
	healing := Rolld(10)

	if r.hp+healing >= r.maxhp {
		fmt.Printf("%s is popping up a healing potion, healing to max!\n", r.name)
		r.hp = r.maxhp
	} else {
		r.hp += healing
		fmt.Printf("%s is popping up a healing potion, healing %d health points.\n", r.name, healing)
	}
}
