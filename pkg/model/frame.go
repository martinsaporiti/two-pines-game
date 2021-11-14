package model

type frame struct {
	Score  int
	Number int
	TryOne Try
	TryTwo Try
}

type Getters interface {
	GetTryOne() Try
	GetTryTwo() Try
	GetNumber() int
	GetScore() int
}

type Frame interface {
	Getters
	FrameElement
	addTry(knockedDownPins int, player Player) bool
	calculateScore(player Player, totalScore int) int
	calculateStrikeScoreBonus(player Player) int
}
