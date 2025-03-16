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
		players:       []Player{{Name: "Abc", score: 98, HoldScore: 2, wins: 0}, {Name: "Def", score: 10, HoldScore: 15, wins: 0}},
		numberOfGames: 1,
		gameOverScore: 1,
		diceRoller:    MockDice{value: 2},
		winningScore:  100,
		expectedResult: []PlayerWiseStats{
			{
				Name:         "Abc",
				Wins:         1,
				Losses:       0,
				HoldingScore: 2,
			},
			{
				Name:         "Def",
				Wins:         0,
				Losses:       1,
				HoldingScore: 15,
			},
		},
	}}

	for _, testCase := range test_cases {
		g := Game{
			NumberOfGames: testCase.numberOfGames,
			GameOverScore: testCase.gameOverScore,
			Players:       testCase.players,
			DiceRoller:    testCase.diceRoller,
			WinningScore:  testCase.winningScore,
		}
		result := g.Play()

		if !reflect.DeepEqual(testCase.expectedResult, result) {
			t.Errorf("Expected %v Got %v", testCase.expectedResult, result)
		}
	}
}
