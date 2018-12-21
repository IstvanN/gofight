package main

import (
	"fmt"
	"strings"

	. "github.com/IstvanN/gofight/classes"
)

func main() {
	fmt.Println("Welcome to GoFight!")
	var f1 Fighter = CreateRogue("Wolfgraff")
	var f2 Fighter = CreateRogue("Grunt")
	arena(f1, f2)
}

func arena(f1 Fighter, f2 Fighter) {
	fmt.Printf("Our fighters tonight are %s and %s!\n", f1.GetName(), f2.GetName())
	for !eitherOneDead(f1, f2) {
		if Rolld(2) == 1 {
			f1.Act(f2)
			f2.Act(f1)
		} else {
			f2.Act(f1)
			f1.Act(f2)
		}
		status(f1, f2)
	}
	announceWinner(f1, f2)
}

func eitherOneDead(f1 Fighter, f2 Fighter) bool {
	return f1.IsDead() || f2.IsDead()
}

func announceWinner(f1 Fighter, f2 Fighter) {
	if f1.IsDead() {
		fmt.Printf("Ladies and Gentlemen, the winner is %s!", strings.ToUpper(f2.GetName()))
	} else {
		fmt.Printf("Ladies and Gentlemen, the winner is %s!", strings.ToUpper(f1.GetName()))
	}
}

func status(f1 Fighter, f2 Fighter) {
	fmt.Printf("%s has %d hp left.\n%s has %d hp left.\n", f1.GetName(), f1.GetHP(), f2.GetName(), f2.GetHP())
}
