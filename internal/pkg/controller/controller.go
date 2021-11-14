package controller

type Controller interface {
	Play() string
	GetPlayers() []string
	GetPlayerScore(player string) int
}
