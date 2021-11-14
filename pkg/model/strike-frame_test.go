package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTryForStrikeFrame(t *testing.T) {
	// given
	strikeFrame := NewStrikeFrame(9)

	mockPlayer := new(MockPlayer)
	mockPlayer.On("createFrame", 10, 5)

	// when
	result := strikeFrame.addTry(5, mockPlayer)

	// then
	mockPlayer.AssertExpectations(t)
	assert.True(t, result)
}

func TestCalculateScoreForStrikeFrame(t *testing.T) {
	// given
	strikeFrame := NewStrikeFrame(1)
	mockPlayer := new(MockPlayer)
	mockPlayer.On("calculateStrikeScoreBonus").Return(1)

	// when
	result := strikeFrame.calculateScore(mockPlayer, 100)

	// then
	mockPlayer.AssertExpectations(t)
	assert.Equal(t, 111, strikeFrame.Score)
	assert.Equal(t, 111, result)
}

func TestCalculateStrikeScoreBonusForStrikeFrame(t *testing.T) {
	// given
	strikeFrame := NewStrikeFrame(5)
	mockPlayer := new(MockPlayer)
	mockPlayer.On("getFirstValueFromNextFrame").Return(1)

	// when
	result := strikeFrame.calculateStrikeScoreBonus(mockPlayer)

	// then
	mockPlayer.AssertExpectations(t)
	assert.Equal(t, 11, result)
}

func TestAcceptStrikeFrame(t *testing.T) {
	// given
	strikeFrame := NewStrikeFrame(5)
	mockFrameVisitor := new(MockFrameVisitor)
	mockFrameVisitor.On("VisitStrikeFrame", strikeFrame)

	// when
	strikeFrame.Accept(mockFrameVisitor)

	// then
	mockFrameVisitor.AssertCalled(t, "VisitStrikeFrame", strikeFrame)
}
