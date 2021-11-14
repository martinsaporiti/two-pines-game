package printer

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/martinsaporiti/two-pines-game/pkg/model"
)

var (
	printer         strings.Builder
	scoresPrinter   strings.Builder
	pinFallsPrinter strings.Builder
	FRAME           string = "Frame"
	PINFALLS        string = "Pinfalls"
	SCORE           string = "Score"
	DELIMITER       string = "\t"
	LONG_DELIMITER  string = "\t\t"
	SPARE_CHAR      string = "/"
	STRIKE_CHAR     string = "X"
	NEW_LINE        string = "\n"
)

type PrinterGame struct {
	model.FrameVisitor
}

func NewPrinterGame() *PrinterGame {
	return &PrinterGame{}
}

func (pr *PrinterGame) Print(game model.Game) string {
	game.Accept(pr)
	fmt.Print(printer.String())
	return printer.String()
}

func (pr *PrinterGame) VisitGame(game model.Game) {
	printHeader()
	for _, p := range game.GetPlayers() {
		p.AcceptVisitor(pr)
	}
}

func printHeader() {
	printer.WriteString(FRAME + LONG_DELIMITER)
	for i := 1; i <= 9; i++ {
		printer.WriteString(strconv.Itoa(i))
		printer.WriteString(LONG_DELIMITER)
	}
	printer.WriteString("10" + NEW_LINE)
}

func (pr *PrinterGame) VisitPlayer(player model.Player) {
	printer.WriteString(player.GetName() + NEW_LINE)
	printFrames(player.GetFrames(), pr)
	printer.WriteString(NEW_LINE)
}

func printFrames(frames []model.Frame, terminalPrinter *PrinterGame) {
	pinFallsPrinter.Reset()
	scoresPrinter.Reset()
	pinFallsPrinter.WriteString(PINFALLS + DELIMITER)
	scoresPrinter.WriteString(SCORE + LONG_DELIMITER)
	for _, f := range frames {
		f.Accept(terminalPrinter)
	}
	printer.WriteString(pinFallsPrinter.String())
	printer.WriteString(NEW_LINE)
	printer.WriteString(scoresPrinter.String())
}

func (pr *PrinterGame) VisitTry(try model.Try) {
	var tryStr string
	if try.IsFoul() {
		tryStr = "F"
	} else if try.GetKnockedDownPins() == 10 {
		tryStr = STRIKE_CHAR
	} else {
		tryStr = strconv.Itoa(try.GetKnockedDownPins())
	}
	pinFallsPrinter.WriteString(tryStr + DELIMITER)
}

func (pr *PrinterGame) VisitNormalFrame(frame model.NormalFrame) {
	frame.GetTryOne().Accept(pr)
	frame.GetTryTwo().Accept(pr)
	score := strconv.Itoa(frame.GetScore())
	scoresPrinter.WriteString(score + LONG_DELIMITER)
}

func (pr *PrinterGame) VisitSpeareFrame(frame model.SpareFrame) {
	frame.GetTryOne().Accept(pr)
	pinFallsPrinter.WriteString(SPARE_CHAR + DELIMITER)
	score := strconv.Itoa(frame.GetScore())
	scoresPrinter.WriteString(score + LONG_DELIMITER)
}

func (pr *PrinterGame) VisitStrikeFrame(frame model.StrikeFrame) {
	pinFallsPrinter.WriteString(DELIMITER + STRIKE_CHAR + DELIMITER)
	score := strconv.Itoa(frame.GetScore())
	scoresPrinter.WriteString(score + LONG_DELIMITER)
}

func (pr *PrinterGame) VisitLastFrame(frame model.LastFrame) {
	frame.GetTryOne().Accept(pr)
	frame.GetTryTwo().Accept(pr)
	if frame.GetTryThree() != nil {
		frame.GetTryThree().Accept(pr)
	}
	score := strconv.Itoa(frame.GetScore())
	scoresPrinter.WriteString(score + LONG_DELIMITER)
}
