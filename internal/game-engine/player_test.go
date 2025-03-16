package gameengine

import (
	"testing"
)

type MockDice struct {
	value int
}

func (d MockDice) Roll() int {
	return d.value
}

func TestPlayer_Play(t *testing.T) {
	testCases := []struct {
		name     string
		diceVal  int
		expected int
	}{{
		"Roll 1", 1, 1,
	}, {
		"Roll 6", 6, 6,
	}, {
		"Roll 3", 3, 3,
	}}

	player := Player{
		Name: "Abc",
	}

	for _, testCase := range testCases {
		result := player.Play(MockDice{value: testCase.diceVal})
		if result != testCase.expected {
			t.Errorf("Expected %v Got %v", testCase.expected, result)
		}
	}
}

func TestPlayer_AddWins(t *testing.T) {
	player := Player{
		Name: "Abc",
	}

	if player.wins != 0 {
		t.Errorf("Expected %v Got %v", 0, player.wins)
	}

	player.AddWins()

	if player.wins != 1 {
		t.Errorf("Expected %v Got %v", 1, player.wins)
	}
}

func TestPlayer_AddScore(t *testing.T) {
	player := Player{
		Name: "Abc",
	}

	if player.score != 0 {
		t.Errorf("Expected %v Got %v", 0, player.score)
	}

	player.AddScore(50)

	if player.score != 50 {
		t.Errorf("Expected %v Got %v", 50, player.score)
	}
}

func TestPlayer_ContinuePlayer(t *testing.T) {
	player := Player{
		Name:      "Abc",
		HoldScore: 20,
	}

	if player.ContinuePlaying(30) != false {
		t.Errorf("Expected %v Got %v", false, true)
	}

	if player.ContinuePlaying(19) != true {
		t.Errorf("Expected %v Got %v", true, false)
	}
}
