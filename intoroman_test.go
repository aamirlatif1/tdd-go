package tdd_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestIntToRomanTestSuite(t *testing.T) {
	suite.Run(t, new(IntToRomanTestSuite))
}

type IntToRomanTestSuite struct {
	suite.Suite
}

func (suite *IntToRomanTestSuite) SetupTest() {
}

func (suite *IntToRomanTestSuite) TestIntToRoman() {
	assert.Equal(suite.T(), "I", intToRoman(1))
	assert.Equal(suite.T(), "II", intToRoman(2))
	assert.Equal(suite.T(), "III", intToRoman(3))
	assert.Equal(suite.T(), "IV", intToRoman(4))
	assert.Equal(suite.T(), "V", intToRoman(5))
	assert.Equal(suite.T(), "VII", intToRoman(7))
	assert.Equal(suite.T(), "IX", intToRoman(9))
	assert.Equal(suite.T(), "X", intToRoman(10))
	assert.Equal(suite.T(), "XXXIX", intToRoman(39))
	assert.Equal(suite.T(), "XL", intToRoman(40))
	assert.Equal(suite.T(), "XLIX", intToRoman(49))
	assert.Equal(suite.T(), "L", intToRoman(50))
	assert.Equal(suite.T(), "LXXXIX", intToRoman(89))
	assert.Equal(suite.T(), "XC", intToRoman(90))
	assert.Equal(suite.T(), "XCIX", intToRoman(99))
	assert.Equal(suite.T(), "C", intToRoman(100))
	assert.Equal(suite.T(), "MCMXCIV", intToRoman(1994))
}
