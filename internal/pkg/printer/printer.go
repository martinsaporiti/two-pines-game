package printer

import "github.com/martinsaporiti/two-pines-game/internal/pkg/model"

type Printer interface {
	Print(game model.Game) string
}
