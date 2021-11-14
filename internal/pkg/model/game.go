package model

import (
	"log"
	"sync"
)

type GameType struct {
	players map[string]Player
}

type Game interface {
	AddTryToPlayer(player string, knockedDownPins int) bool
	CalculateScores()
	Validate() bool
	GetPlayers() []Player
	GameElement
}

func NewGame() *GameType {
	return &GameType{
		players: make(map[string]Player),
	}
}

func (g *GameType) GetPlayers() []Player {
	players := make([]Player, 0)
	for _, player := range g.players {
		players = append(players, player)
	}
	return players
}

// Add a new try to the player
func (g *GameType) AddTryToPlayer(player string, knockedDownPins int) bool {
	if g.players[player] == nil {
		g.players[player] = NewPlayer(player)
	}
	return g.players[player].addTry(knockedDownPins)
}

// Calculates the score for the players with goroutines.
func (g *GameType) CalculateScores() {
	var goRoutine sync.WaitGroup
	for _, player := range g.players {
		goRoutine.Add(1)
		go func(p Player) {
			defer goRoutine.Done()
			p.calculateScore()
		}(player)
	}
	goRoutine.Wait()
}

// Validates the model generated per the loaded data.
func (g *GameType) Validate() bool {
	var goRoutine sync.WaitGroup
	ok := true
	for _, player := range g.players {
		goRoutine.Add(1)
		ok = ok && player.validateFrames()
	}
	if !ok {
		log.Panicln("Bad input data to play. Please verify your input data")
	}
	return true
}

func (g *GameType) Accept(gameVisitor GameVisitor) {
	gameVisitor.VisitGame(g)
}
