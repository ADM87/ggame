package game

import (
	"image/color"

	"github.com/ADM87/ggame/components"
	"github.com/ADM87/ggame/entities"
	"github.com/ADM87/ggame/keyboard"
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
	entities []entities.IEntity
	entity   entities.IEntity // For testing purposes
	child    entities.IEntity // For testing purposes
}

func NewGame() Game {
	return &gameshell{
		entities: make([]entities.IEntity, 0),
		entity:   nil, // Will be set in Start()
	}
}

func (g *gameshell) Start() error {
	g.RegisterKeys()

	tile0000Image := resources.LoadImage("tile_0000")
	tile0010Image := resources.LoadImage("tile_0010")

	tile0000Renderer := components.NewSpriteRenderer()
	tile0000Renderer.SetImage(tile0000Image)

	tile0010Renderer := components.NewSpriteRenderer()
	tile0010Renderer.SetImage(tile0010Image)

	g.child = entities.NewEntity()
	g.child.SetRenderer(tile0010Renderer)
	g.child.SetScale(0.5, 0.5)
	g.child.SetPosition(0, 0)
	g.child.SetOrigin(12, 12)

	g.entity = entities.NewEntity()
	g.entity.SetRenderer(tile0000Renderer)
	g.entity.SetPosition(float64(ScreenWidth/2), float64(ScreenHeight/2))
	g.entity.SetOrigin(12, 12)
	g.entity.SetScale(2, 2)
	g.entity.AddChild(g.child)
	g.entities = append(g.entities, g.entity)

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
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowDown, keyboard.KeyPhaseHeld, func() error {
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowLeft, keyboard.KeyPhaseHeld, func() error {
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowRight, keyboard.KeyPhaseHeld, func() error {
		return nil
	})
}
