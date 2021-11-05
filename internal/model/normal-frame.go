package model

type normalFrame struct {
	frame
}

type NormalFrame interface {
	Frame
	FrameElement
}

func NewNormalFrame(number int) *normalFrame {
	return &normalFrame{
		frame: frame{
			Number: number,
		},
	}
}

func NewNormalFrameWithTry(number int, tryOne Try) *normalFrame {
	return &normalFrame{
		frame: frame{
			Number: number,
			TryOne: tryOne,
		},
	}
}

func (nf *normalFrame) GetTryOne() Try {
	return nf.TryOne
}

func (nf *normalFrame) GetTryTwo() Try {
	return nf.TryTwo
}

func (nf *normalFrame) GetNumber() int {
	return nf.Number
}

// Add a new try to the frame...
func (nf *normalFrame) addTry(knockedDownPins int, player Player) bool {
	if nf.TryTwo == nil {
		if nf.TryOne.GetKnockedDownPins()+knockedDownPins == 10 {
			newFrame := NewSpareFrameWithTries(nf.GetNumber(), nf.TryOne, NewTry(knockedDownPins))
			player.updateFrame(newFrame)
		} else if nf.TryOne.GetKnockedDownPins()+knockedDownPins < 10 {
			nf.TryTwo = NewTry(knockedDownPins)
		} else {
			return false
		}
	} else {
		player.createFrame(nf.Number+1, knockedDownPins)
	}
	return true
}

func (nf *normalFrame) calculateScore(player Player, totalScore int) int {
	nf.Score = totalScore + nf.TryOne.GetKnockedDownPins() + nf.TryTwo.GetKnockedDownPins()
	return nf.Score
}

func (nf *normalFrame) calculateStrikeScoreBonus(player Player) int {
	return nf.TryOne.GetKnockedDownPins() + nf.TryTwo.GetKnockedDownPins()
}

func (nf *normalFrame) Accept(frameVisitor FrameVisitor) {
	frameVisitor.VisitNormalFrame(nf)
}

func (nf *normalFrame) GetScore() int {
	return nf.Score
}
