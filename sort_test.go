package tdd_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSortTestSuite(t *testing.T) {
	suite.Run(t, new(SortTestSuite))
}

type SortTestSuite struct {
	suite.Suite
}

func (suite *SortTestSuite) SetupTest() {
}

func (suite *SortTestSuite) TestSort() {
	assert.EqualValues(suite.T(), []int{}, Sort([]int{}))
	assert.EqualValues(suite.T(), []int{1}, Sort([]int{1}))
	assert.EqualValues(suite.T(), []int{1, 2}, Sort([]int{1, 2}))
	assert.EqualValues(suite.T(), []int{1, 2}, Sort([]int{2, 1}))
	assert.EqualValues(suite.T(), []int{1, 2, 3}, Sort([]int{1, 2, 3}))
	assert.EqualValues(suite.T(), []int{1, 2, 3}, Sort([]int{2, 1, 3}))
	assert.EqualValues(suite.T(), []int{1, 2, 3}, Sort([]int{3, 1, 2}))
	assert.EqualValues(suite.T(), []int{1, 2, 3}, Sort([]int{3, 2, 1}))
	assert.EqualValues(suite.T(), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9},
		Sort([]int{9, 4, 8, 0, 7, 1, 9, 2, 3, 6, 5}))
}
