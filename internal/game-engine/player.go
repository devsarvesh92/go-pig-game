package gameengine

type Player struct {
	Name      string
	HoldScore int
	score     int
	wins      int
}

func (p Player) Play(dice DiceRoller) int {
	return dice.Roll()
}

func (p *Player) AddWins() {
	p.wins += 1
}

func (p *Player) AddScore(score int) {
	p.score += score
}

func (p *Player) ContinuePlaying(runningScore int) bool {
	return p.HoldScore > runningScore
}

func (p *Player) ResetScore() {
	p.score = 0
}
