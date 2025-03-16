package main

import (
	"fmt"
	gameengine "gopig/internal/game-engine"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pig [player1HoldScore] [player2HoldScore]",
	Short: "A simulation of the pig dice game",
	Long: `Pig is a dice game simulation where two players compete using different strategies.
Each player has a "hold score" strategy that determines when they stop rolling.`,
	Args: cobra.ExactArgs(2), // Enforce exactly 2 arguments
	Run: func(cmd *cobra.Command, args []string) {
		p1HoldScore, err := strconv.Atoi(args[0])
		if err != nil || p1HoldScore < 1 || p1HoldScore > 100 {
			fmt.Printf("Invalid hold score %v for player 1: must be between 1-100\n", args[0])
			os.Exit(1)
		}

		p2HoldScore, err := strconv.Atoi(args[1])
		if err != nil || p2HoldScore < 1 || p2HoldScore > 100 {
			fmt.Printf("Invalid hold score %v for player 2: must be between 1-100\n", args[1])
			os.Exit(1)
		}

		startGame(p1HoldScore, p2HoldScore)
	},
}

func startGame(p1HoldScore, p2HoldScore int) {
	// Create game with appropriate configuration
	g := gameengine.Game{
		NumberOfGames: 10,
		GameOverScore: 1,
		Players: []gameengine.Player{
			{Name: "Player 1", HoldScore: p1HoldScore},
			{Name: "Player 2", HoldScore: p2HoldScore},
		},
		DiceRoller:   gameengine.Dice{Sides: 6},
		WinningScore: 100,
	}

	// Run the game and get results
	results := g.Play()

	print(p1HoldScore, p2HoldScore, results)
}

func print(p1HoldScore, p2HoldScore int, results []gameengine.PlayerWiseStats) {
	// Extract data for both players
	player1 := results[0]

	totalGames := player1.Wins + player1.Losses
	winPercentage := float64(player1.Wins) / float64(totalGames) * 100
	lossPercentage := float64(player1.Losses) / float64(totalGames) * 100

	fmt.Printf(
		"Result: Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
		p1HoldScore,
		p2HoldScore,
		player1.Wins,
		totalGames,
		winPercentage,
		player1.Losses,
		totalGames,
		lossPercentage,
	)
}

func init() {
	rootCmd.SetUsageTemplate(
		`pig - A command line tool to simulate a game of pig. It is a two-player game played with a 6-sided die.

Usage:
	pig [strategy1] [strategy2]

Args:
	strategy   The number between 1 to 100

Example usage:
	$ pig 10 15
	Result: Holding at  10 vs Holding at 15: wins: 3/10 (30.0%), losses: 7/10 (70.0%)
`,
	)
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
