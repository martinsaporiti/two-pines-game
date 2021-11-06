package model

type strikeFrame struct {
	frame
}

type StrikeFrame interface {
	Frame
}

// Constructor
func NewStrikeFrame(number int) *strikeFrame {
	strikeFrame := &strikeFrame{
		frame: frame{
			Number: number,
		},
	}
	strikeFrame.TryOne = NewTry(10)
	return strikeFrame
}

func (sf *strikeFrame) GetNumber() int {
	return sf.Number
}

func (sf *strikeFrame) GetTryOne() Try {
	return sf.TryOne
}

func (sf *strikeFrame) GetTryTwo() Try {
	return sf.TryTwo
}

// Add a new try
func (sf *strikeFrame) addTry(knockedDownPins int, player Player) bool {
	player.createFrame(sf.Number+1, knockedDownPins)
	return true
}

func (sf *strikeFrame) calculateScore(player Player, totalScore int) int {
	sf.Score = totalScore + 10 + player.calculateStrikeScoreBonus(sf)
	return sf.Score
}

func (sf *strikeFrame) calculateStrikeScoreBonus(player Player) int {
	return 10 + player.getFirstValueFromNextFrame(sf)
}

func (sf *strikeFrame) Accept(frameVisitor FrameVisitor) {
	frameVisitor.VisitStrikeFrame(sf)
}

func (sf *strikeFrame) GetScore() int {
	return sf.Score
}
