package model

type spareFrame struct {
	frame
}

type SpareFrame interface {
	Frame
}

func NewSpareFrame(number int) *spareFrame {
	return &spareFrame{
		frame: frame{
			Number: number,
		},
	}
}

func NewSpareFrameWithTries(number int, tryOne, tryTwo Try) *spareFrame {
	spareFrame := &spareFrame{
		frame: frame{
			Number: number,
		},
	}
	spareFrame.TryOne = tryOne
	spareFrame.TryTwo = tryTwo
	return spareFrame
}

func (sf *spareFrame) GetTryOne() Try {
	return sf.TryOne
}

func (sf *spareFrame) GetTryTwo() Try {
	return sf.TryTwo
}

func (sf *spareFrame) GetNumber() int {
	return sf.Number
}

func (sf *spareFrame) addTry(knockedDownPins int, player Player) bool {
	player.createFrame(sf.Number+1, knockedDownPins)
	return true
}

func (sf *spareFrame) calculateScore(player Player, totalScore int) int {
	sf.Score = totalScore + 10 + player.calculateSpareScoreBonus(sf)
	return sf.Score
}

func (sf *spareFrame) calculateStrikeScoreBonus(player Player) int {
	return 10
}

func (sf *spareFrame) Accept(frameVisitor FrameVisitor) {
	frameVisitor.VisitSpeareFrame(sf)
}

func (sf *spareFrame) GetScore() int {
	return sf.Score
}
