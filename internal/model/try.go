package model

type TryType struct {
	KnockedDownPins int
	foul            bool
	hasValue        bool
}

type Try interface {
	TryElement
	GetKnockedDownPins() int
	IsFoul() bool
}

func NewTry(knockedDownPins int) *TryType {
	newTry := &TryType{
		KnockedDownPins: knockedDownPins,
		hasValue:        true,
	}
	if knockedDownPins == -1 {
		newTry.foul = true
		newTry.KnockedDownPins = 0
	}
	return newTry
}

func (try *TryType) Accept(tryVisitor TryVisitor) {
	tryVisitor.VisitTry(try)
}

func (try *TryType) GetKnockedDownPins() int {
	return try.KnockedDownPins
}

func (try *TryType) IsFoul() bool {
	return try.foul
}
