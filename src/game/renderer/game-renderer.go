package renderer

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer interface {
	Draw(screen *ebiten.Image)
	SetClearColor(c color.Color)
}

type gamerenderer struct {
	clearColor color.Color

	screenWidth  int
	screenHeight int
}

func NewGameRenderer(width, height int) Renderer {
	return &gamerenderer{
		clearColor:   color.RGBA{100, 149, 237, 255},
		screenWidth:  width,
		screenHeight: height,
	}
}

func (g *gamerenderer) Draw(screen *ebiten.Image) {
	screen.Fill(g.clearColor)
}

func (g *gamerenderer) SetClearColor(c color.Color) {
	g.clearColor = c
}
