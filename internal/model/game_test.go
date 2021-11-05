package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetPlayers(t *testing.T) {
	// given
	mockPlayer := new(MockPlayer)
	game := NewGame()
	game.players["test player"] = mockPlayer

	// when
	players := game.GetPlayers()

	// then
	assert.True(t, len(players) > 0)
}

func TestAddTryToPlayer(t *testing.T) {
	// given
	mockPlayer := new(MockPlayer)
	mockPlayer.On("addTry", 5).Return(true)
	game := NewGame()
	game.players["test player"] = mockPlayer

	// when
	game.AddTryToPlayer("test player", 5)

	// then
	mockPlayer.AssertExpectations(t)
}

func TestAddTryToPlayerFirstAttemp(t *testing.T) {
	// given
	game := NewGame()

	// when
	game.AddTryToPlayer("test player", 5)

	// then
	assert.Equal(t, 1, len(game.players))
}

func TestCalculateScores(t *testing.T) {
	// given
	mockPlayer1 := new(MockPlayer)
	mockPlayer2 := new(MockPlayer)
	mockPlayer1.On("calculateScore").Return(100)
	mockPlayer2.On("calculateScore").Return(100)

	game := NewGame()
	game.players["test player 1"] = mockPlayer1
	game.players["test player 2"] = mockPlayer2

	// when
	game.CalculateScores()

	// then
	mockPlayer1.AssertExpectations(t)
	mockPlayer2.AssertExpectations(t)
}

func TestValidateOnePlayerFails(t *testing.T) {
	// given
	mockPlayer1 := new(MockPlayer)
	mockPlayer2 := new(MockPlayer)
	mockPlayer1.On("validateFrames").Return(true)
	mockPlayer2.On("validateFrames").Return(false)

	game := NewGame()
	game.players["test player 1"] = mockPlayer1
	game.players["test player 2"] = mockPlayer2

	defer func() { recover() }()

	// when
	result := game.Validate()

	// then
	mockPlayer1.AssertExpectations(t)
	mockPlayer2.AssertExpectations(t)
	assert.False(t, result)
}

func TestValidateTwoPlayerSuccess(t *testing.T) {
	// given
	mockPlayer1 := new(MockPlayer)
	mockPlayer2 := new(MockPlayer)
	mockPlayer1.On("validateFrames").Return(true)
	mockPlayer2.On("validateFrames").Return(true)

	game := NewGame()
	game.players["test player 1"] = mockPlayer1
	game.players["test player 2"] = mockPlayer2

	// when
	result := game.Validate()

	// then
	mockPlayer1.AssertExpectations(t)
	mockPlayer2.AssertExpectations(t)
	assert.True(t, result)
}

func TestAcceptGameVisitor(t *testing.T) {
	// given
	game := NewGame()
	mockGameVisitor := new(MockGameVisitor)
	mockGameVisitor.On("VisitGame", game)

	// when
	game.Accept(mockGameVisitor)

	// then
	mockGameVisitor.AssertCalled(t, "VisitGame", game)
}
