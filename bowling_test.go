package tdd_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestBowlingTestSuite(t *testing.T) {
	suite.Run(t, new(BowlingTestSuite))
}

type BowlingTestSuite struct {
	suite.Suite
	game *game
}

func (suite *BowlingTestSuite) SetupTest() {
	suite.game = NewGame()
}

func (suite *BowlingTestSuite) rollMany(n int, pins int) {
	for i := 0; i < n; i++ {
		suite.game.Roll(pins)
	}
}

func (suite *BowlingTestSuite) rollSpare() {
	suite.rollMany(2, 5)
}

func (suite *BowlingTestSuite) rollStrike() {
	suite.game.Roll(10)
}

func (suite *BowlingTestSuite) TestGutterGame() {
	suite.rollMany(20, 0)
	assert.Equal(suite.T(), 0, suite.game.Score())
}

func (suite *BowlingTestSuite) TestAllOnes() {
	suite.rollMany(20, 1)
	assert.Equal(suite.T(), 20, suite.game.Score())
}

func (suite *BowlingTestSuite) TestOneSpare() {
	suite.rollSpare()
	suite.game.Roll(7)
	suite.rollMany(17, 0)
	assert.Equal(suite.T(), 24, suite.game.Score())
}

func (suite *BowlingTestSuite) TestOneStrike() {
	suite.rollStrike()
	suite.game.Roll(2)
	suite.game.Roll(3)
	assert.Equal(suite.T(), 20, suite.game.Score())
}

func (suite *BowlingTestSuite) TestPerfectGame() {
	suite.rollMany(12, 10)
	assert.Equal(suite.T(), 300, suite.game.Score())
}
