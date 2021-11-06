package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTryForSpeareFrameGettingNewNormalFrameWithTry(t *testing.T) {
	// given
	player := NewPlayer("test player")
	player.addTry(5)
	spareFrame := NewSpareFrame(1)
	expectedFrame := NewNormalFrameWithTry(2, NewTry(5))

	// when
	result := spareFrame.addTry(5, player)

	// then
	assert.True(t, result)
	assert.Nil(t, spareFrame.TryTwo)
	assert.Equal(t, player.frames[1], expectedFrame)

}

func TestAddTryForSpeare(t *testing.T) {
	// given
	speareFrame := NewSpareFrame(9)
	mockPlayer := new(MockPlayer)
	mockPlayer.On("createFrame", 10, 5).Return(2)

	// when
	result := speareFrame.addTry(5, mockPlayer)

	// then
	assert.True(t, result)
	mockPlayer.AssertExpectations(t)
}
func TestCalculateScoreForSpeareFrame(t *testing.T) {
	// given
	speareFrame := NewSpareFrameWithTries(0, NewTry(2), NewTry(8))
	mockPlayer := new(MockPlayer)
	mockPlayer.On("calculateSpareScoreBonus", speareFrame).Return(2)

	// when
	result := speareFrame.calculateScore(mockPlayer, 100)

	// then
	mockPlayer.AssertExpectations(t)
	assert.Equal(t, 112, speareFrame.Score)
	assert.Equal(t, 112, result)
}

func TestCalculateStrikeScoreBonusForSpeareFrame(t *testing.T) {
	// given
	speareFrame := NewSpareFrameWithTries(0, NewTry(2), NewTry(8))
	mockPlayer := new(MockPlayer)

	// when
	result := speareFrame.calculateStrikeScoreBonus(mockPlayer)

	// then
	assert.Equal(t, 10, result)
	mockPlayer.AssertNotCalled(t, "getFirstValueFromNextFrame")
}

func TestAcceptSpeareFrame(t *testing.T) {
	// given
	speareFrame := NewSpareFrame(5)
	mockFrameVisitor := new(MockFrameVisitor)
	mockFrameVisitor.On("VisitSpeareFrame", speareFrame)

	// when
	speareFrame.Accept(mockFrameVisitor)

	// then
	mockFrameVisitor.AssertCalled(t, "VisitSpeareFrame", speareFrame)
}
