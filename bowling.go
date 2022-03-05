package tdd_go

type game struct {
	rolls       [21]int
	currentRoll int
}

func NewGame() *game {
	return &game{}
}

func (g *game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g *game) Score() int {
	score := 0
	frameIndex := 0
	for frame := 0; frame < 10; frame++ {
		if g.isStrike(frameIndex) {
			score += 10 + g.strikeBonus(frameIndex)
			frameIndex++
		} else if g.isSpare(frameIndex) {
			score += 10 + g.spareBonus(frameIndex)
			frameIndex += 2
		} else {
			score += g.towBallInFrame(frameIndex)
			frameIndex += 2
		}
	}
	return score
}

func (g *game) towBallInFrame(frameIndex int) int {
	return g.rolls[frameIndex] + g.rolls[frameIndex+1]
}

func (g *game) spareBonus(frameIndex int) int {
	return g.rolls[frameIndex+2]
}

func (g *game) strikeBonus(frameIndex int) int {
	return g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
}

func (g *game) isStrike(frameIndex int) bool {
	return g.rolls[frameIndex] == 10
}

func (g *game) isSpare(frameIndex int) bool {
	return g.rolls[frameIndex]+g.rolls[frameIndex+1] == 10
}
