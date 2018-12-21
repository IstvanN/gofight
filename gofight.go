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
