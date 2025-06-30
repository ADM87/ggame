package game

import (
	"image/color"

	"github.com/ADM87/ggame/resources"
	"github.com/ADM87/ggame/src/camera"
	"github.com/ADM87/ggame/src/components"
	"github.com/ADM87/ggame/src/keyboard"
	"github.com/ADM87/ggame/src/objects"
	"github.com/ADM87/ggame/src/sys"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640 * 0.6
	ScreenHeight = 480 * 0.6
)

var (
	backgroundColor = color.RGBA{100, 149, 237, 255}
)

type Game interface {
	Start() error
}

type gameshell struct {
	gameCamera camera.Camera

	player objects.Actor
	actors []objects.Actor
}

func createTestActor(r components.SpriteRenderer, x, y float64) objects.Actor {
	actor := objects.NewActor()
	actor.SetRenderer(r)
	actor.Transform().SetPosition(x, y)
	actor.Transform().SetOrigin(12, 24)
	return actor
}

func NewGame() Game {
	return &gameshell{
		gameCamera: camera.NewCamera(0, 0, ScreenWidth, ScreenHeight),
	}
}

func (g *gameshell) Start() error {
	testImage, err := resources.LoadImage("tile_0000")
	if err != nil {
		panic(err)
	}
	testRenderer := components.NewSpriteRenderer()
	testRenderer.SetImage(testImage)

	g.player = createTestActor(testRenderer, float64(ScreenWidth)/2, float64(ScreenHeight)/2)
	g.actors = make([]objects.Actor, 10)

	for i := 0; i < len(g.actors); i++ {
		x := float64(i%5) * 32
		y := float64(i/5) * 32
		g.actors[i] = createTestActor(testRenderer, x, y)
	}

	keyboard.RegisterKey(ebiten.KeyEscape, keyboard.KeyPhaseDown, func() error {
		sys.ShutdownWith(0)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowUp, keyboard.KeyPhaseHeld, func() error {
		x, y := g.player.Transform().Position()
		g.player.Transform().SetPosition(x, y-1)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowDown, keyboard.KeyPhaseHeld, func() error {
		x, y := g.player.Transform().Position()
		g.player.Transform().SetPosition(x, y+1)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowLeft, keyboard.KeyPhaseHeld, func() error {
		x, y := g.player.Transform().Position()
		g.player.Transform().SetPosition(x-1, y)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowRight, keyboard.KeyPhaseHeld, func() error {
		x, y := g.player.Transform().Position()
		g.player.Transform().SetPosition(x+1, y)
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
