package gameengine

import "math/rand"

type Dice struct {
}

func (d Dice) Roll() int {
	return rand.Intn(6) + 1
}
