package classes

import (
	"math/rand"
	"time"
)

// Fighter is the main interface of the application which all classes should implement
type Fighter interface {
	Act(Fighter)
	Suffer(dmg int)
	Heal()
}

// Rolld function is the dice of the app
func Rolld(dice int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(dice) + 1
	return r
}
