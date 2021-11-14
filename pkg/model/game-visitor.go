package model

type GameVisitor interface {
	VisitGame(game Game)
}

type GameElement interface {
	Accept(gameVisitor GameVisitor)
}
