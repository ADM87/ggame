package game

import (
	"image/color"

	"github.com/ADM87/ggame/src/camera"
	"github.com/ADM87/ggame/src/keyboard"
	"github.com/ADM87/ggame/src/sys"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

var (
	backgroundColor = color.RGBA{100, 149, 237, 255}
)

type Game interface {
	Start() error
}

type gameshell struct {
	gameCamera camera.Camera
}

func NewGame() Game {
	return &gameshell{
		gameCamera: camera.NewCamera(0, 0, ScreenWidth, ScreenHeight),
	}
}

func (g *gameshell) Start() error {
	keyboard.RegisterKey(ebiten.KeyEscape, keyboard.KeyPhaseDown, func() error {
		sys.ShutdownWith(0)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowUp, keyboard.KeyPhaseHeld, func() error {
		g.gameCamera.MoveBy(0, -1)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowDown, keyboard.KeyPhaseHeld, func() error {
		g.gameCamera.MoveBy(0, 1)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowLeft, keyboard.KeyPhaseHeld, func() error {
		g.gameCamera.MoveBy(-1, 0)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowRight, keyboard.KeyPhaseHeld, func() error {
		g.gameCamera.MoveBy(1, 0)
		return nil
	})
	return ebiten.RunGame(g)
}

func (g *gameshell) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth, screenHeight = ScreenWidth, ScreenHeight

	if outsideWidth < screenWidth {
		screenWidth = outsideWidth
	}
	if outsideHeight < screenHeight {
		screenHeight = outsideHeight
	}

	return screenWidth, screenHeight
}
