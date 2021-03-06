package controller

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/martinsaporiti/two-pines-game/pkg/model"
	"github.com/martinsaporiti/two-pines-game/pkg/printer"
	"github.com/martinsaporiti/two-pines-game/pkg/reader"
)

const (
	version string = "v0.0.3"
)

var game model.Game

type ControllerImpl struct {
	reader  reader.Reader
	printer printer.Printer
}

func NewController(reader reader.Reader, printer printer.Printer) *ControllerImpl {
	return &ControllerImpl{
		reader,
		printer,
	}
}

// Play the game, loading the data from file, validating the generated model and
// calculating the score for each player.
func (ctrl *ControllerImpl) Play() string {
	fmt.Println("Running Version:", version)
	game = loadGame(ctrl.reader)
	correct := game.Validate()

	if !correct {
		panic("Error validating Game")
	}

	game.CalculateScores()
	return ctrl.printer.Print(game)
}

// Return the players names in the game.
func (ctrl *ControllerImpl) GetPlayers() []string {
	var playersNames []string
	for _, p := range game.GetPlayers() {
		playersNames = append(playersNames, p.GetName())
	}
	return playersNames
}

// Return the score per a given player.
func (ctrl *ControllerImpl) GetPlayerScore(player string) int {
	return game.GetPlayerScore(player)
}

// Create a game from data inside the file.
func loadGame(reader reader.Reader) model.Game {
	data, _ := processInputData(reader)
	game := model.NewGame()
	for player, knockedDownPinsArr := range data {
		for _, knockedDownPins := range knockedDownPinsArr {
			ok := game.AddTryToPlayer(player, knockedDownPins)
			if !ok {
				log.Panicf("Error processing: %v - %v", player, knockedDownPins)
			}
		}
	}
	return game
}

// Proccess the content of file and returns the knowked down pins for each player.
func processInputData(reader reader.Reader) (map[string][]int, error) {
	contentFile := reader.GetContent()
	if len(contentFile) == 0 {
		log.Panicf("Invalid Input Data. Empty file.")
	}

	var inputData = make(map[string][]int)
	for _, r := range contentFile {
		splitedRow := strings.Split(r, "\t")
		if len(splitedRow) != 2 {
			log.Panicln("Invalid Input Data")
		}
		// get the player...
		player := splitedRow[0]
		// get knoked downpins...
		knockedDownPins := readKnockedDownPins(splitedRow[1])
		// add the knoked down pins to the player...
		inputData[player] = append(inputData[player], knockedDownPins)
	}

	return inputData, nil
}

// Read the value for the knowd down pins field.
func readKnockedDownPins(s string) int {
	if s == "F" {
		return -1
	}
	knoked, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("You have provided incorrect data to play the game. Each pinfall must be between 0 and 10 or be F")
	}

	if isValidKnokedInputNumber(knoked) {
		log.Panicf("You have provided incorrect data to play the game. Each pinfall must be between 0 and 10")
	}
	return knoked
}

// knoked input number should be between 0 and 10.
func isValidKnokedInputNumber(knoked int) bool {
	return knoked < 0 || knoked > 10
}
