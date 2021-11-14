package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTryWithFramesBefore(t *testing.T) {
	// given
	player := NewPlayer("player test")
	strikeFrame := NewStrikeFrame(1)
	player.frames = append(player.frames, strikeFrame)
	mockFrame := new(MockFrame)
	expectedFrame := NewNormalFrameWithTry(2, NewTry(5))
	mockFrame.On("addTry", 5, player).Return(expectedFrame)

	// when
	result := player.addTry(5)

	// then
	assert.True(t, result)
	assert.True(t, len(player.frames) == 2)
	assert.Equal(t, expectedFrame, player.frames[1])
}

func TestAddTryWithNoFramesBefore(t *testing.T) {
	// given
	player := NewPlayer("player test")
	expectedFrame := NewNormalFrameWithTry(1, NewTry(5))

	// when
	result := player.addTry(5)

	// then
	assert.True(t, result)
	assert.True(t, len(player.frames) == 1)
	assert.Equal(t, expectedFrame, player.frames[0])
}

func TestCalculateScore(t *testing.T) {
	// given
	player := NewPlayer("player test")
	mockFrame1 := new(MockFrame)
	mockFrame1.On("calculateScore", player, 0).Return(100)
	mockFrame2 := new(MockFrame)
	mockFrame2.On("calculateScore", player, 100).Return(150)

	player.frames = append(player.frames, mockFrame1, mockFrame2)

	// when
	result := player.calculateScore()

	// then
	assert.Equal(t, 150, result)
	mockFrame1.AssertExpectations(t)
	mockFrame2.AssertExpectations(t)
}

func TestCreateLastFrame(t *testing.T) {
	// given
	player := NewPlayer("player test")
	expectedFrame := NewLastFrameWithTryOne(10, NewTry(5))

	// when
	player.createFrame(10, 5)

	// then
	assert.Equal(t, expectedFrame, player.frames[0])
}

func TestCreateStrikeFrame(t *testing.T) {
	// given
	player := NewPlayer("player test")
	expectedFrame := NewStrikeFrame(1)

	// when
	player.createFrame(1, 10)

	// then
	assert.Equal(t, expectedFrame, player.frames[0])
}

func TestCreateFrameWithTry(t *testing.T) {
	// given
	player := NewPlayer("player test")
	expectedFrame := NewNormalFrameWithTry(1, NewTry(5))

	// when
	player.createFrame(1, 5)

	// then
	assert.Equal(t, expectedFrame, player.frames[0])
}

func TestUpdateFrame(t *testing.T) {
	// given
	player := NewPlayer("player test")
	mockFrame := new(MockFrame)
	player.frames = append(player.frames, mockFrame)
	newFrame := NewNormalFrame(1)

	// when
	player.updateFrame(newFrame)

	// then
	assert.Equal(t, newFrame, player.frames[0])
}

func TestCalculateStrikeScoreBonus(t *testing.T) {
	// given
	player := NewPlayer("player test")
	mockFrame := new(MockFrame)
	mockFrame.On("GetNumber").Return(0)
	mockFrame.On("calculateStrikeScoreBonus", player).Return(4)
	player.frames = append(player.frames, mockFrame)

	// when
	result := player.calculateStrikeScoreBonus(mockFrame)

	// then
	assert.Equal(t, 4, result)
	mockFrame.AssertExpectations(t)
}

func TestCalCaculateSpareScoreBonus(t *testing.T) {
	// given
	player := NewPlayer("player test")
	mockFrame := new(MockFrame)
	mockFrame.On("GetNumber").Return(0)
	mockFrame.On("GetTryOne").Return(NewTry(5))
	player.frames = append(player.frames, mockFrame)

	// when
	result := player.calculateSpareScoreBonus(mockFrame)

	// then
	assert.Equal(t, 5, result)
	mockFrame.AssertExpectations(t)
}

func TestGetName(t *testing.T) {
	// given
	player := NewPlayer("player test")

	// when
	result := player.GetName()

	// then
	assert.Equal(t, "player test", result)
}

func TestGetFrames(t *testing.T) {
	// given
	player := NewPlayer("player test")

	// when
	result := player.GetFrames()

	// then
	assert.Equal(t, 0, len(result))
}

func TestValidateFramesWithLessThan10(t *testing.T) {
	// given
	player := NewPlayer("player test")

	// when
	result := player.validateFrames()

	// then
	assert.Equal(t, false, result)
}

func TestValidateFramesWith10Frames(t *testing.T) {
	// given
	player := NewPlayer("player test")
	mockFrame := new(MockFrame)
	mockFrame.On("isComplete").Return(true)

	for i := 0; i <= 9; i++ {
		lastFrame := NewLastFrameWithTryOne(i, NewTry(3))
		lastFrame.TryTwo = NewTry(3)
		player.frames = append(player.frames, lastFrame)
	}

	player.frames = append(player.frames, mockFrame)

	// when
	result := player.validateFrames()

	// then
	assert.Equal(t, true, result)
}

func TestAcceptPlayerVisitor(t *testing.T) {
	// given
	player := NewPlayer("player")
	mockPlayerVisitor := new(MockPlayerVisitor)
	mockPlayerVisitor.On("VisitPlayer", player)

	// when
	player.AcceptVisitor(mockPlayerVisitor)

	// then
	mockPlayerVisitor.AssertCalled(t, "VisitPlayer", player)
}

func TestGetScore(t *testing.T) {
	// given
	player := NewPlayer("player")
	player.score = 10

	// when
	score := player.getScore()

	// then
	assert.Equal(t, 10, score)
}
