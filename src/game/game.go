package game

import (
	"image/color"

	"github.com/ADM87/ggame/src/components"
	"github.com/ADM87/ggame/src/keyboard"
	"github.com/ADM87/ggame/src/sys"
	"github.com/ADM87/ggame/src/sys/window"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GameScreenWidth  = 640
	GameScreenHeight = 480
)

type Game interface {
	Start() error
}

type gameshell struct {
	window.Window

	components.Updater
	components.Renderer
}

func NewGame(windowWidth, windowHeight int) Game {
	return &gameshell{
		Updater:  components.NewUpdater(),
		Renderer: components.NewRenderer(),
		Window:   window.NewWindow(windowWidth, windowHeight),
	}
}

func (g *gameshell) Start() error {
	g.Renderer.SetColor(color.RGBA{100, 149, 237, 255})

	keyboard.RegisterKey(ebiten.KeyEscape, keyboard.KeyPhaseDown, func() error {
		sys.Shutdown()
		return nil
	})

	return ebiten.RunGame(g)
}
