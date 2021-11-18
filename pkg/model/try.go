package model

type try struct {
	KnockedDownPins int
	foul            bool
	hasValue        bool
}

type Try interface {
	TryElement
	GetKnockedDownPins() int
	IsFoul() bool
}

func NewTry(knockedDownPins int) *try {
	newTry := &try{
		KnockedDownPins: knockedDownPins,
		hasValue:        true,
	}
	if knockedDownPins == -1 {
		newTry.foul = true
		newTry.KnockedDownPins = 0
	}
	return newTry
}

func (t *try) Accept(tryVisitor TryVisitor) {
	tryVisitor.VisitTry(t)
}

func (t *try) GetKnockedDownPins() int {
	return t.KnockedDownPins
}

func (t *try) IsFoul() bool {
	return t.foul
}
