package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTryForNormalFrameWithoutTryTwoCreatingNewTry(t *testing.T) {
	// given
	normalFrame := NewNormalFrameWithTry(1, NewTry(5))
	expectedFrame := NewNormalFrameWithTry(1, NewTry(5))
	expectedFrame.TryTwo = NewTry(2)

	mockPlayer := new(MockPlayer)
	mockPlayer.On("updateFrame", expectedFrame)

	// when
	result := normalFrame.addTry(2, mockPlayer)

	// then
	assert.True(t, result)
	assert.Equal(t, expectedFrame, normalFrame)
}

func TestAddTryForNormalFrameWithoutTryTwoCreatingSpeareFrame(t *testing.T) {
	// given
	normalFrame := NewNormalFrameWithTry(1, NewTry(5))
	expectedFrame := NewSpareFrameWithTries(1, NewTry(5), NewTry(5))
	mockPlayer := new(MockPlayer)
	mockPlayer.On("updateFrame", expectedFrame)

	// when
	result := normalFrame.addTry(5, mockPlayer)

	// then
	assert.True(t, result)
	mockPlayer.AssertExpectations(t)
}

func TestAddTryForNormalFrameWithTryTwo(t *testing.T) {
	// given
	normalFrame := NewNormalFrameWithTry(9, NewTry(5))
	normalFrame.TryTwo = NewTry(1)
	mockPlayer := new(MockPlayer)
	mockPlayer.On("createFrame", 10, 5)

	// when
	result := normalFrame.addTry(5, mockPlayer)

	// then
	assert.True(t, result)
	assert.NotNil(t, normalFrame.TryTwo)
	mockPlayer.AssertExpectations(t)
}

func TestCalculateScoreForNormalFrame(t *testing.T) {
	// given
	normalFrame := NewNormalFrameWithTry(1, NewTry(5))
	normalFrame.TryTwo = NewTry(2)
	mockPlayer := new(MockPlayer)

	// when
	result := normalFrame.calculateScore(mockPlayer, 10)

	// then
	assert.Equal(t, 17, normalFrame.Score)
	assert.Equal(t, 17, result)
	mockPlayer.AssertNotCalled(t, "calculateStrikeScoreBonus")
	mockPlayer.AssertNotCalled(t, "calculateSpareScoreBonus")
}

func TestCalculateStrikeScoreBonusForNormalFrame(t *testing.T) {
	// given
	normalFrame := NewNormalFrameWithTry(1, NewTry(5))
	normalFrame.TryTwo = NewTry(2)
	mockPlayer := new(MockPlayer)

	// when
	result := normalFrame.calculateStrikeScoreBonus(mockPlayer)

	// then
	assert.Equal(t, 7, result)
	mockPlayer.AssertNotCalled(t, "getFirstValueFromNextFrame")
	mockPlayer.AssertNotCalled(t, "calculateStrikeScoreBonus")
	mockPlayer.AssertNotCalled(t, "calculateSpareScoreBonus")
}

func TestAcceptNormalFrame(t *testing.T) {
	// given
	normalFrame := NewNormalFrame(5)
	mockFrameVisitor := new(MockFrameVisitor)
	mockFrameVisitor.On("VisitNormalFrame", normalFrame)

	// when
	normalFrame.Accept(mockFrameVisitor)

	// then
	mockFrameVisitor.AssertCalled(t, "VisitNormalFrame", normalFrame)
}
