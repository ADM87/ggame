package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Renderer interface {
	Draw(screen *ebiten.Image)
}

type gamerenderer struct {
	screenWidth  int
	screenHeight int
}

func NewGameRenderer(width, height int) Renderer {
	return &gamerenderer{
		screenWidth:  width,
		screenHeight: height,
	}
}

func (g *gamerenderer) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Ggame!")
}
