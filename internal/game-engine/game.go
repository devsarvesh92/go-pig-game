package gameengine

type Game struct {
	NumberOfGames int
	GameOverScore int
	Players       []Player
	DiceRoller    DiceRoller
	WinningScore  int
}

type PlayerWiseStats struct {
	Name         string
	Wins         int
	Losses       int
	HoldingScore int
}

func (g Game) Play() []PlayerWiseStats {

	for i := 1; i <= g.NumberOfGames; i++ {
		gameOver := false
		for j := range g.Players {
			player := &g.Players[j]
			player.ResetScore()
		}

		for !gameOver {
			for j := range g.Players {
				player := &g.Players[j]

				runningScore := 0
				for {
					diceVal := player.Play(g.DiceRoller)
					if diceVal == g.GameOverScore {
						break
					}
					runningScore += diceVal

					if !player.ContinuePlaying(runningScore) {
						// Add score
						player.AddScore(runningScore)
						break
					}
				}

				// Check if player is winner already
				if player.score >= g.WinningScore {
					player.AddWins()
					gameOver = true
					break
				}
			}
		}

	}

	playerWiseStats := make([]PlayerWiseStats, 0, len(g.Players))
	for p := range g.Players {
		player := &g.Players[p]
		playerWiseStats = append(playerWiseStats, PlayerWiseStats{
			Name:         player.Name,
			Wins:         player.wins,
			Losses:       g.NumberOfGames - player.wins,
			HoldingScore: player.HoldScore,
		})
	}
	return playerWiseStats
}
