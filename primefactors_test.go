package tdd_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPrimeFactorSuite(t *testing.T) {
	suite.Run(t, new(PrimeFactorSuite))
}

type PrimeFactorSuite struct {
	suite.Suite
}

func (suite *PrimeFactorSuite) SetupTest() {
}

func (suite *PrimeFactorSuite) TestFactors() {
	assert.Equal(suite.T(), 0, len(factorsOf(1)))
	assert.EqualValues(suite.T(), []int{2}, factorsOf(2))
	assert.EqualValues(suite.T(), []int{3}, factorsOf(3))
	assert.EqualValues(suite.T(), []int{2, 2}, factorsOf(4))
	assert.EqualValues(suite.T(), []int{5}, factorsOf(5))
	assert.EqualValues(suite.T(), []int{2, 3}, factorsOf(6))
	assert.EqualValues(suite.T(), []int{7}, factorsOf(7))
	assert.EqualValues(suite.T(), []int{2, 2, 2}, factorsOf(8))
	assert.EqualValues(suite.T(), []int{3, 3}, factorsOf(9))
	assert.EqualValues(suite.T(), []int{2, 2, 3, 3, 3, 5, 7, 7}, factorsOf(7*7*5*3*3*3*2*2))

}
