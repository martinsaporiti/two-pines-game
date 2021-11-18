package controller

// Defines the api with the entire game
// You must call controller to interact with the game.
type Controller interface {
	Play() string
	GetPlayers() []string
	GetPlayerScore(player string) int
}
