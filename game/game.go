package game

import (
	"image/color"

	"github.com/ADM87/ggame/camera"
	"github.com/ADM87/ggame/components"
	"github.com/ADM87/ggame/keyboard"
	"github.com/ADM87/ggame/objects"
	"github.com/ADM87/ggame/resources"
	"github.com/ADM87/ggame/sys"
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

	player   objects.Entity
	entities []objects.Entity
}

func createTestActor(r components.SpriteRenderer, x, y float64) objects.Entity {
	actor := objects.NewActor()
	actor.SetRenderer(r)
	actor.SetPosition(x, y)
	actor.SetOrigin(12, 24)
	return actor
}

func NewGame() Game {
	return &gameshell{
		gameCamera: camera.NewCamera(0, 0, ScreenWidth, ScreenHeight),
	}
}

func (g *gameshell) Start() error {
	tile0000Image, err := resources.LoadImage("tile_0000")
	if err != nil {
		panic(err)
	}
	tile0000Renderer := components.NewSpriteRenderer()
	tile0000Renderer.SetImage(tile0000Image)

	tile0010Image, err := resources.LoadImage("tile_0010")
	if err != nil {
		panic(err)
	}
	tile0010Renderer := components.NewSpriteRenderer()
	tile0010Renderer.SetImage(tile0010Image)

	playerChild := createTestActor(tile0000Renderer, 10, 10)
	playerChild.SetPosition(10, 10)

	g.player = createTestActor(tile0000Renderer, float64(ScreenWidth)/2, float64(ScreenHeight)/2)
	g.player.AddChild(playerChild)

	g.entities = make([]objects.Entity, 0)
	g.entities = append(g.entities, g.player)
	g.entities = append(g.entities, createTestActor(tile0010Renderer, 100, 100))

	g.RegisterKeys()

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

func (g *gameshell) RegisterKeys() {
	keyboard.RegisterKey(ebiten.KeyEscape, keyboard.KeyPhaseDown, func() error {
		sys.ShutdownWith(0)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowUp, keyboard.KeyPhaseHeld, func() error {
		x, y := g.player.Position()
		g.player.SetPosition(x, y-1)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowDown, keyboard.KeyPhaseHeld, func() error {
		x, y := g.player.Position()
		g.player.SetPosition(x, y+1)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowLeft, keyboard.KeyPhaseHeld, func() error {
		x, y := g.player.Position()
		g.player.SetPosition(x-1, y)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowRight, keyboard.KeyPhaseHeld, func() error {
		x, y := g.player.Position()
		g.player.SetPosition(x+1, y)
		return nil
	})
}
