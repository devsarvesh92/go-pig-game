# Pig Game Simulator

A command-line tool to simulate the classic Pig dice game with configurable player strategies.

## About the Game

Pig is a simple dice game where players take turns rolling a six-sided die. On each turn, a player can roll the die repeatedly until either:

1. They choose to "hold" and add their accumulated points to their total score
2. They roll a 1, which ends their turn with no points gained

The first player to reach 100 points wins the game.

## Features

- Simulate games between two players with different "hold" strategies
- Players can use different holding strategies (1-100)
- Run multiple game simulations to analyze win rates
- Statistical analysis of different strategies

## Installation

Clone the repository and build the executable:

```bash
git clone https://github.com/yourusername/go-pig-game.git
cd go-pig-game
make build
```

This will create an executable named `pig` in your project directory.

## Usage

Run the game simulation by providing two holding strategies (integers between 1-100):

```bash
./pig 10 15
```

This example simulates 10 games where Player 1 holds at 10 points and Player 2 holds at 15 points.

Example output:
```
Result: Holding at 10 vs Holding at 15: wins: 4/10 (40.0%), losses: 6/10 (60.0%)
```

## Development

### Prerequisites

- Go 1.18 or higher

### Running Tests

Run all tests:
```bash
make test
```

Run tests with verbose output:
```bash
make test-verbose
```

Generate test coverage report:
```bash
make test-coverage
```

Generate HTML coverage report:
```bash
make test-coverage-html
```

## Project Structure

```
go-pig-game/
├── cmd/
│   └── main.go         # Main application entry point
├── internal/
│   └── game-engine/    # Game logic implementation
│       ├── game.go
│       ├── player.go
│       └── dice.go
├── Makefile            # Build and test commands
└── README.md           # This file
```

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.