package gameengine

import "math/rand"

type DiceRoller interface {
	Roll() int
}

type Dice struct {
	sides int
}

func (d Dice) Roll() int {
	return rand.Intn(d.sides) + 1
}
