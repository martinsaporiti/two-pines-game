package printer

import "github.com/martinsaporiti/two-pines-game/internal/model"

type Printer interface {
	Print(game model.Game) string
}
