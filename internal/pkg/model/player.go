package model

type player struct {
	name   string
	score  int
	frames []Frame
}

type Player interface {
	createFrame(number int, knockedDownPins int)
	updateFrame(frame Frame)
	calculateSpareScoreBonus(frame SpareFrame) int
	calculateStrikeScoreBonus(frame StrikeFrame) int
	getFirstValueFromNextFrame(frame Frame) int
	addTry(knockedDownPins int) bool
	calculateScore() int
	validateFrames() bool
	getScore() int
	GetName() string
	GetFrames() []Frame
	PlayerElement
}

// Constructor...
func NewPlayer(name string) *player {
	return &player{name: name}
}

func (p *player) GetName() string {
	return p.name
}

func (p *player) GetFrames() []Frame {
	return p.frames
}

// Add a new Try to the player.
func (p *player) addTry(knockedDownPins int) bool {
	if len(p.frames) == 0 {
		p.createFrame(1, knockedDownPins)
		return true
	} else {
		lastFrame := p.frames[len(p.frames)-1]
		return lastFrame.addTry(knockedDownPins, p)
	}
}

// Return the score for the player.
// Set score field with the calculated score.
func (p *player) calculateScore() int {
	totalScore := 0
	for i := 0; i < len(p.frames); i++ {
		totalScore = p.frames[i].calculateScore(p, totalScore)

	}
	p.score = totalScore
	return totalScore
}

// Create a new frame for the player.
func (p *player) createFrame(number int, knockedDownPins int) {
	var frame Frame
	if number == 10 {
		frame = NewLastFrameWithTryOne(number, NewTry(knockedDownPins))
	} else if knockedDownPins == 10 {
		frame = NewStrikeFrame(number)
	} else {
		frame = NewNormalFrameWithTry(number, NewTry(knockedDownPins))
	}
	p.frames = append(p.frames, frame)
}

func (p *player) updateFrame(frame Frame) {
	p.frames[frame.GetNumber()-1] = frame
}

func (p *player) calculateSpareScoreBonus(frame SpareFrame) int {
	return p.getFirstValueFromNextFrame(frame)
}

func (p *player) calculateStrikeScoreBonus(frame StrikeFrame) int {
	return p.frames[frame.GetNumber()].calculateStrikeScoreBonus(p)
}

func (p *player) getFirstValueFromNextFrame(frame Frame) int {
	return p.frames[frame.GetNumber()].GetTryOne().GetKnockedDownPins()
}

func (p *player) validateFrames() bool {
	if len(p.frames) < 9 {
		return false
	}
	return p.frames[9].(LastFrame).isComplete()
}

func (p *player) getScore() int {
	return p.score
}

func (p *player) AcceptVisitor(playerVisitor PlayerVisitor) {
	playerVisitor.VisitPlayer(p)
}
