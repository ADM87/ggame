package game

import (
	"github.com/ADM87/ggame/src/game/loop"
	"github.com/ADM87/ggame/src/game/renderer"
	"github.com/ADM87/ggame/src/game/window"
	"github.com/ADM87/ggame/src/keyboard"
	"github.com/ADM87/ggame/src/sys"
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
	loop.Loop
	renderer.Renderer
	window.Window
}

func NewGame(windowWidth, windowHeight int) Game {
	return &gameshell{
		Loop:     loop.NewGameLoop(),
		Renderer: renderer.NewGameRenderer(GameScreenWidth, GameScreenHeight),
		Window:   window.NewGameWindow(GameScreenWidth, GameScreenHeight),
	}
}

func (g *gameshell) Start() error {
	keyboard.RegisterKey(ebiten.KeyEscape, keyboard.KeyPhaseDown, func() error {
		sys.Shutdown()
		return nil
	})
	return ebiten.RunGame(g)
}
