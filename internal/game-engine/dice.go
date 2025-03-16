package gameengine

import "math/rand"

type DiceRoller interface {
	Roll() int
}

type Dice struct {
	Sides int
}

func (d Dice) Roll() int {
	return rand.Intn(d.Sides) + 1
}
