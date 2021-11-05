package main

import (
	"log"
	"os"
	"time"

	"github.com/martinsaporiti/two-pines-game/internal/controller"
	"github.com/martinsaporiti/two-pines-game/internal/printer"
	"github.com/martinsaporiti/two-pines-game/internal/reader"
)

func main() {
	defer handleErrors()
	arg := os.Args[1]
	fr := reader.NewFileReader(arg)
	pr := printer.NewPrinterGame()
	ctrl := controller.NewController(fr, pr)
	fileContent := ctrl.Play()
	createFile(fileContent)
}

func handleErrors() {
	if r := recover(); r != nil {
		log.Println("=> You can not play the game :( Recovered: ", r)
	}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createFile(fileContent string) {
	t := time.Now()
	filePrefix := "./" + t.Format("20060102150405")
	f, err := os.Create(filePrefix + "_result.txt")
	check(err)
	defer f.Close()
	_, err = f.WriteString(fileContent)
	check(err)
}