package gameengine

type Player struct {
	name      string
	score     int
	holdScore int
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
	return p.holdScore > runningScore
}
