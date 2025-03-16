package gameengine

import (
	"reflect"
	"testing"
)

func TestGamePlay(t *testing.T) {
	test_cases := []struct {
		players        []Player
		numberOfGames  int
		gameOverScore  int
		diceRoller     DiceRoller
		winningScore   int
		expectedResult []PlayerWiseStats
	}{{
		players:       []Player{{name: "Abc", score: 98, holdScore: 2, wins: 0}, {name: "Def", score: 10, holdScore: 15, wins: 0}},
		numberOfGames: 1,
		gameOverScore: 1,
		diceRoller:    MockDice{value: 2},
		winningScore:  100,
		expectedResult: []PlayerWiseStats{
			{
				name:         "Abc",
				wins:         1,
				losses:       0,
				score:        100,
				holdingScore: 2,
			},
			{
				name:         "Def",
				wins:         0,
				losses:       1,
				score:        10,
				holdingScore: 15,
			},
		},
	}}

	for _, testCase := range test_cases {
		g := Game{
			numberOfGames: testCase.numberOfGames,
			gameOverScore: testCase.gameOverScore,
			players:       testCase.players,
			diceRoller:    testCase.diceRoller,
			winningScore:  testCase.winningScore,
		}
		result := g.Play()

		if !reflect.DeepEqual(testCase.expectedResult, result) {
			t.Errorf("Expected %v Got %v", testCase.expectedResult, result)
		}
	}
}
