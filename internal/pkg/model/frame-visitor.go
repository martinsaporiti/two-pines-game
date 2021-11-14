package model

type FrameVisitor interface {
	VisitNormalFrame(frame NormalFrame)
	VisitLastFrame(frame LastFrame)
	VisitSpeareFrame(frame SpareFrame)
	VisitStrikeFrame(frame StrikeFrame)
}

type FrameElement interface {
	Accept(frameVisitor FrameVisitor)
}
