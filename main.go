package main

import (
	"fmt"

	. "github.com/IstvanN/gofight/classes"
)

func main() {
	fmt.Println("Welcome to GoFight!")
	var f1 Fighter = CreateRogue("Wolfgraff")
	var f2 Fighter = CreateRogue("Grunt")
	arena(f1, f2)
}

func arena(f1 Fighter, f2 Fighter) {
	fmt.Printf("Our fighters tonight are %s and %s!", f1.GetName(), f2.GetName())
	for {
		if Rolld(2) == 1 {
			f1.Act(f2)
			f2.Act(f1)
		}
	}
}
