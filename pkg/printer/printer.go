package printer

import "github.com/martinsaporiti/two-pines-game/pkg/model"

type Printer interface {
	Print(game model.Game) string
}
