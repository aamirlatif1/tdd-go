package tdd_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestStackTestSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}

type StackTestSuite struct {
	suite.Suite
	stack *stack
}

func (suite *StackTestSuite) SetupTest() {
	suite.stack = NewStack()
}

func (suite *StackTestSuite) TestCanCreateStack() {
	assert.True(suite.T(), suite.stack.IsEmpty())
}

func (suite *StackTestSuite) TestAfterOnePush_isNotEmpty() {
	suite.stack.Push(0)
	assert.False(suite.T(), suite.stack.IsEmpty())
	assert.Equal(suite.T(), 1, suite.stack.Size())

}

func (suite *StackTestSuite) TestAfterOnePushAndOnePop_isEmpty() {
	suite.stack.Push(0)
	_, _ = suite.stack.Pop()
	assert.True(suite.T(), suite.stack.IsEmpty())
	assert.Equal(suite.T(), 0, suite.stack.Size())

}

func (suite *StackTestSuite) TestAfterTwoPushes_sizeIsTwo() {
	suite.stack.Push(1)
	suite.stack.Push(2)
	assert.Equal(suite.T(), 2, suite.stack.Size())
}

func (suite *StackTestSuite) TestPoppingEmptyStack_generateError() {
	_, err := suite.stack.Pop()
	assert.EqualError(suite.T(), err, "stack underflow")
}

func (suite *StackTestSuite) TestAfterPushX_willPopX() {
	suite.stack.Push(99)
	poppedValue, err := suite.stack.Pop()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 99, poppedValue)

	suite.stack.Push(88)
	poppedValue, err = suite.stack.Pop()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 88, poppedValue)
}

func (suite *StackTestSuite) TestAfterPushingXThenY_popYThenX() {
	suite.stack.Push(99)
	suite.stack.Push(88)

	xvalue, _ := suite.stack.Pop()
	assert.Equal(suite.T(), 88, xvalue)
	yValue, _ := suite.stack.Pop()
	assert.Equal(suite.T(), 99, yValue)
}
