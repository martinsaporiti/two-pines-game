package model

type lastFrame struct {
	frame
	TryThree Try
}

type LastFrame interface {
	Frame
	GetTryThree() Try
	HasTryThree() bool
	isComplete() bool
}

// Constructor
func NewLastFrame(number int) *lastFrame {
	return &lastFrame{
		frame: frame{
			Number: number,
		},
	}
}

// Constructor
func NewLastFrameWithTryOne(number int, tryOne Try) *lastFrame {
	return &lastFrame{
		frame: frame{
			Number: number,
			TryOne: tryOne,
		},
	}
}

func (lf *lastFrame) GetTryOne() Try {
	return lf.TryOne
}

func (lf *lastFrame) GetTryTwo() Try {
	return lf.TryTwo
}

func (lf *lastFrame) GetNumber() int {
	return lf.Number
}

func (lf *lastFrame) GetScore() int {
	return lf.Score
}

func (lf *lastFrame) HasTryThree() bool {
	return lf.TryThree != nil
}

func (lf *lastFrame) GetTryThree() Try {
	return lf.TryThree
}

// Add a new Try.
func (lf *lastFrame) addTry(knockedDownPins int, player Player) bool {
	if lf.TryTwo == nil {
		lf.TryTwo = NewTry(knockedDownPins)
		return true
	} else if lf.TryThree == nil && (lf.TryOne.GetKnockedDownPins()+lf.TryTwo.GetKnockedDownPins()) >= 10 {
		lf.TryThree = NewTry(knockedDownPins)
		return true
	}
	return false
}

func (lf *lastFrame) calculateScore(player Player, totalScore int) int {
	score := lf.TryOne.GetKnockedDownPins() + lf.TryTwo.GetKnockedDownPins()
	if lf.TryThree != nil {
		score += lf.TryThree.GetKnockedDownPins()
	}
	lf.Score = score + totalScore
	return lf.Score
}

func (lf *lastFrame) calculateStrikeScoreBonus(player Player) int {
	return lf.TryOne.GetKnockedDownPins() + lf.TryTwo.GetKnockedDownPins()
}

func (lf *lastFrame) isComplete() bool {
	if lf.TryTwo == nil {
		return false
	}
	if lf.TryOne.GetKnockedDownPins()+lf.TryTwo.GetKnockedDownPins() >= 10 {
		return lf.TryThree != nil
	}
	return true
}

func (lf *lastFrame) Accept(frameVisitor FrameVisitor) {
	frameVisitor.VisitLastFrame(lf)
}
