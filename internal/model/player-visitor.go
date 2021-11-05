package model

type PlayerVisitor interface {
	VisitPlayer(player Player)
}

type PlayerElement interface {
	AcceptVisitor(playerVisitor PlayerVisitor)
}
