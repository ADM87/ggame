package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Renderer interface {
	Draw(screen *ebiten.Image)
}

type gamerenderer struct {
}

func NewGameRenderer() Renderer {
	return &gamerenderer{}
}

func (g *gamerenderer) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Ggame!")
}
