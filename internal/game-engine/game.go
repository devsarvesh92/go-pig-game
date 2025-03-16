package gameengine

type Game struct {
	numberOfGames int
	gameOverScore int
	players       []Player
	diceRoller    DiceRoller
	winningScore  int
}

type PlayerWiseStats struct {
	name         string
	wins         int
	losses       int
	score        int
	holdingScore int
}

func (g Game) Play() []PlayerWiseStats {

	for i := 1; i <= g.numberOfGames; i++ {

		for j := range g.players {
			player := &g.players[j]

			runningScore := 0
			for {
				diceVal := player.Play(g.diceRoller)
				if diceVal == g.gameOverScore {
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
			if player.score >= g.winningScore {
				player.AddWins()
				break
			}
		}
	}

	// Build Playerwise stats
	playerWiseStats := make([]PlayerWiseStats, 0, len(g.players))

	for p := range g.players {
		player := &g.players[p]
		playerWiseStats = append(playerWiseStats, PlayerWiseStats{
			name:         player.name,
			wins:         player.wins,
			losses:       g.numberOfGames - player.wins,
			score:        player.score,
			holdingScore: player.holdScore,
		})
	}
	return playerWiseStats
}
