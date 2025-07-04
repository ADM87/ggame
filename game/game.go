package game

import (
	"image/color"

	"github.com/ADM87/ggame/core"
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

	tile0000Renderer core.ISpriteRenderer
	tile0010Renderer core.ISpriteRenderer

	moveX, moveY float64
)

type Game interface {
	Start() error
}

type gameshell struct {
}

func NewGame() Game {
	return &gameshell{}
}

func (g *gameshell) Start() error {
	g.RegisterKeys()

	tile0000Renderer = core.NewSpriteRenderer()
	tile0000Renderer.SetImage(resources.LoadImage("tile_0000"))
	tile0000Renderer.SetAnchor(0.5, 1)

	tile0010Renderer = core.NewSpriteRenderer()
	tile0010Renderer.SetImage(resources.LoadImage("tile_0010"))
	tile0010Renderer.SetAnchor(0.5, 0)

	hw, hh := float64(ScreenWidth)*0.5, float64(ScreenHeight)*0.5

	g.CreateTest(hw, hh)

	return ebiten.RunGame(g)
}

func (g *gameshell) CreateTest(x, y float64) core.IEntity {
	grandchildA := core.NewEntity()
	grandchildA.SetRenderer(tile0000Renderer)
	grandchildA.SetPosition(9, 0)
	grandchildA.SetScale(0.5, 0.5)

	grandchildB := core.NewEntity()
	grandchildB.SetRenderer(tile0000Renderer)
	grandchildB.SetPosition(-9, 0)
	grandchildB.SetScale(0.5, 0.5)

	child := core.NewEntity()
	child.SetRenderer(tile0010Renderer)
	child.SetPosition(0, 10)
	child.AddChild(grandchildA)
	child.AddChild(grandchildB)

	parent := core.NewEntity()
	parent.SetRenderer(tile0000Renderer)
	parent.SetPosition(x, y)
	parent.AddChild(child)

	renderables = append(renderables, parent)

	return parent
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
		moveY = -1.0
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowUp, keyboard.KeyPhaseUp, func() error {
		moveY = 0.0
		return nil
	})

	keyboard.RegisterKey(ebiten.KeyArrowDown, keyboard.KeyPhaseHeld, func() error {
		moveY = 1.0
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowDown, keyboard.KeyPhaseUp, func() error {
		moveY = 0.0
		return nil
	})

	keyboard.RegisterKey(ebiten.KeyArrowLeft, keyboard.KeyPhaseHeld, func() error {
		moveX = -1.0
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowLeft, keyboard.KeyPhaseUp, func() error {
		moveX = 0.0
		return nil
	})

	keyboard.RegisterKey(ebiten.KeyArrowRight, keyboard.KeyPhaseHeld, func() error {
		moveX = 1.0
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyArrowRight, keyboard.KeyPhaseUp, func() error {
		moveX = 0.0
		return nil
	})
}
