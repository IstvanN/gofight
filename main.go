package main

import (
	"fmt"
	"strings"

	cl "github.com/IstvanN/gofight/classes"
)

func main() {
	fmt.Println("Welcome to GoFight!")
	var f1 cl.Fighter = cl.CreateRogue("Wolfgraff")
	var f2 cl.Fighter = cl.CreateDruid("Cenarius")
	arena(f1, f2)
}

func arena(f1 cl.Fighter, f2 cl.Fighter) {
	fmt.Printf("Our fighters tonight are %s and %s!\n", f1.GetName(), f2.GetName())
	for !eitherOneDead(f1, f2) {
		if cl.Rolld(2) == 1 {
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

func eitherOneDead(f1 cl.Fighter, f2 cl.Fighter) bool {
	return f1.IsDead() || f2.IsDead()
}

func announceWinner(f1 cl.Fighter, f2 cl.Fighter) {
	if f1.IsDead() {
		fmt.Printf("Ladies and Gentlemen, the winner is %s!", strings.ToUpper(f2.GetName()))
	} else {
		fmt.Printf("Ladies and Gentlemen, the winner is %s!", strings.ToUpper(f1.GetName()))
	}
}

func status(f1 cl.Fighter, f2 cl.Fighter) {
	fmt.Printf("%s has %d hp left.\n%s has %d hp left.\n", f1.GetName(), f1.GetHP(), f2.GetName(), f2.GetHP())
}
