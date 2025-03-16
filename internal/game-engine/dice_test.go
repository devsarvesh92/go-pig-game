package gameengine

import "testing"

func TestRoll(t *testing.T) {
	dice := Dice{Sides: 6}
	result := dice.Roll()

	if result > 6 || result < 1 {
		t.Log("Incorrect value")
		t.Fail()
	}
}
