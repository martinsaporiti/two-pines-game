package model

import "github.com/stretchr/testify/mock"

// ------------------------------
// Mock Player
// ------------------------------

type MockPlayer struct {
	mock.Mock
}

func (mockPlayer *MockPlayer) GetPLayers() []Player {
	args := mockPlayer.Called()
	result := args.Get(0)
	return result.([]Player)
}

func (mockPlayer *MockPlayer) addTry(knockedDownPins int) bool {
	args := mockPlayer.Called(knockedDownPins)
	result := args.Get(0)
	return result.(bool)
}

func (mockPlayer *MockPlayer) calculateScore() int {
	args := mockPlayer.Called()
	result := args.Get(0)
	return result.(int)
}

func (mockPlayer *MockPlayer) createFrame(number int, knockedDownPins int) {
	mockPlayer.Called(10, 5)
}

func (mockPlayer *MockPlayer) updateFrame(frame Frame) {
	mockPlayer.Called(frame)
}

func (mockPlayer *MockPlayer) getFirstValueFromNextFrame(frame Frame) int {
	args := mockPlayer.Called()
	result := args.Get(0)
	return result.(int)
}

func (mockPlayer *MockPlayer) calculateStrikeScoreBonus(frame StrikeFrame) int {
	args := mockPlayer.Called()
	result := args.Get(0)
	return result.(int)
}

func (mockPlayer *MockPlayer) calculateSpareScoreBonus(frame SpareFrame) int {
	args := mockPlayer.Called(frame)
	result := args.Get(0)
	return result.(int)
}

func (mockPlayer *MockPlayer) validateFrames() bool {
	args := mockPlayer.Called()
	result := args.Get(0)
	return result.(bool)
}

func (mockPlayer *MockPlayer) AcceptVisitor(playerVisitor PlayerVisitor) {
	mockPlayer.Called(playerVisitor)
}

func (mockPlayer *MockPlayer) GetName() string {
	args := mockPlayer.Called()
	result := args.Get(0)
	return result.(string)
}
func (mockPlayer *MockPlayer) GetFrames() []Frame {
	args := mockPlayer.Called()
	result := args.Get(0)
	return result.([]Frame)
}

// ------------------------------
// Mock Frame
// ------------------------------

type MockFrame struct {
	mock.Mock
}

func (mockFrame *MockFrame) GetTryOne() Try {
	args := mockFrame.Called()
	result := args.Get(0)
	return result.(Try)
}

func (mockFrame *MockFrame) GetTryTwo() Try {
	args := mockFrame.Called()
	result := args.Get(0)
	return result.(Try)
}

func (mockFrame *MockFrame) GetNumber() int {
	args := mockFrame.Called()
	result := args.Get(0)
	return result.(int)
}

func (mockFrame *MockFrame) addTry(knockedDownPins int, player Player) bool {
	args := mockFrame.Called(knockedDownPins, player)
	result := args.Get(0)
	return result.(bool)
}

func (mockFrame *MockFrame) calculateScore(player Player, totalScore int) int {
	args := mockFrame.Called(player, totalScore)
	result := args.Get(0)
	return result.(int)
}

func (mockFrame *MockFrame) calculateStrikeScoreBonus(player Player) int {
	args := mockFrame.Called(player)
	result := args.Get(0)
	return result.(int)
}

func (mockFrame *MockFrame) Accept(frameVisitor FrameVisitor) {
	frameVisitor.VisitNormalFrame(mockFrame)
}

func (mockFrame *MockFrame) GetScore() int {
	args := mockFrame.Called()
	result := args.Get(0)
	return result.(int)
}

// ------------------------------
// Mock Frame Visitor
// ------------------------------

type MockFrameVisitor struct {
	mock.Mock
}

func (mfv *MockFrameVisitor) VisitNormalFrame(frame NormalFrame) {
	mfv.Called(frame)
}

func (mfv *MockFrameVisitor) VisitLastFrame(frame LastFrame) {
	mfv.Called(frame)
}

func (mfv *MockFrameVisitor) VisitSpeareFrame(frame SpareFrame) {
	mfv.Called(frame)
}

func (mfv *MockFrameVisitor) VisitStrikeFrame(frame StrikeFrame) {
	mfv.Called(frame)
}

// ------------------------------
// Mock Player Visitor
// ------------------------------

type MockPlayerVisitor struct {
	mock.Mock
}

func (mpv *MockPlayerVisitor) VisitPlayer(player Player) {
	mpv.Called(player)
}

// ------------------------------
// Mock Game Visitor
// ------------------------------

type MockGameVisitor struct {
	mock.Mock
}

func (mgv *MockGameVisitor) VisitGame(game Game) {
	mgv.Called(game)
}
