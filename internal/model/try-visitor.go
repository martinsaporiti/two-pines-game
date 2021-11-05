package model

type TryVisitor interface {
	VisitTry(try Try)
}

type TryElement interface {
	Accept(tryVisitor TryVisitor)
}
