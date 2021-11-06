package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTryWithoutTryTwo(t *testing.T) {
	// given
	mockPlayer := new(MockPlayer)
	lastFrame := NewLastFrameWithTryOne(1, NewTry(5))

	// when
	result := lastFrame.addTry(5, mockPlayer)

	// then
	assert.True(t, result)
	assert.Nil(t, lastFrame.TryThree)
	assert.Equal(t, 5, lastFrame.TryTwo.GetKnockedDownPins())
}

func TestAddTryWitTryTwo(t *testing.T) {
	// given
	mockPlayer := new(MockPlayer)
	lastFrame := NewLastFrameWithTryOne(1, NewTry(5))
	lastFrame.TryTwo = NewTry(8)

	// when
	result := lastFrame.addTry(5, mockPlayer)

	// then
	assert.True(t, result)
	assert.Equal(t, 8, lastFrame.TryTwo.GetKnockedDownPins())
	assert.Equal(t, 5, lastFrame.TryThree.GetKnockedDownPins())
}

func TestCalculateScoreWithOneTryAndDontHaveTryThree(t *testing.T) {
	// given
	mockPlayer := new(MockPlayer)

	lastFrame := NewLastFrameWithTryOne(1, NewTry(5))
	lastFrame.TryTwo = NewTry(4)

	// when
	score := lastFrame.calculateScore(mockPlayer, 10)

	// then
	assert.Equal(t, 19, score)
}

func TestCalculateScoreWithTwoTryAndDontHaveTryThree(t *testing.T) {
	// given
	mockPlayer := new(MockPlayer)
	lastFrame := NewLastFrameWithTryOne(1, NewTry(5))
	lastFrame.TryTwo = NewTry(2)

	// when
	score := lastFrame.calculateScore(mockPlayer, 10)

	// then
	assert.Equal(t, 17, score)
}

func TestCalculateScoreWithTwoTryAndHaveTryThree(t *testing.T) {
	// given
	mockPlayer := new(MockPlayer)
	lastFrame := NewLastFrameWithTryOne(1, NewTry(5))
	lastFrame.TryTwo = NewTry(2)
	lastFrame.TryThree = NewTry(2)

	// when
	score := lastFrame.calculateScore(mockPlayer, 10)

	// then
	assert.Equal(t, 19, score)
}

func TestCalculateStrikeScoreBonusForLastFrame(t *testing.T) {
	// given
	lastFrame := NewLastFrameWithTryOne(5, NewTry(5))
	lastFrame.TryTwo = NewTry(2)
	mockPlayer := new(MockPlayer)

	// when
	result := lastFrame.calculateStrikeScoreBonus(mockPlayer)

	// then
	assert.Equal(t, 7, result)
	mockPlayer.AssertNotCalled(t, "calculateStrikeScoreBonus")
}

func TestIsCompleteWithTwoTries(t *testing.T) {
	// given
	lastFrame := NewLastFrameWithTryOne(5, NewTry(5))
	lastFrame.TryTwo = NewTry(2)

	// when
	result := lastFrame.isComplete()

	// then
	assert.Equal(t, true, result)
}

func TestIsCompleteWithOneTries(t *testing.T) {
	// given
	lastFrame := NewLastFrameWithTryOne(5, NewTry(5))

	// when
	result := lastFrame.isComplete()

	// then
	assert.Equal(t, false, result)
}

func TestIsCompleteWithThreeTriesGettinMoreThanTen(t *testing.T) {
	// given
	lastFrame := NewLastFrameWithTryOne(5, NewTry(5))
	lastFrame.TryTwo = NewTry(5)
	lastFrame.TryThree = NewTry(3)

	// when
	result := lastFrame.isComplete()

	// then
	assert.Equal(t, true, result)
}

func TestIsCompleteWithTwoTriesGettinMoreThanTen(t *testing.T) {
	// given
	lastFrame := NewLastFrameWithTryOne(5, NewTry(5))
	lastFrame.TryTwo = NewTry(5)

	// when
	result := lastFrame.isComplete()

	// then
	assert.Equal(t, false, result)
}

func TestAcceptLastFrame(t *testing.T) {
	// given
	lastFrame := NewLastFrame(5)
	mockFrameVisitor := new(MockFrameVisitor)
	mockFrameVisitor.On("VisitLastFrame", lastFrame)

	// when
	lastFrame.Accept(mockFrameVisitor)

	// then
	mockFrameVisitor.AssertCalled(t, "VisitLastFrame", lastFrame)
}
