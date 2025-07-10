package game

import (
	"image/color"

	"github.com/ADM87/ggame/components"
	"github.com/ADM87/ggame/ecs"
	"github.com/ADM87/ggame/keyboard"
	"github.com/ADM87/ggame/rendering"
	"github.com/ADM87/ggame/resources"
	"github.com/ADM87/ggame/sys"
	"github.com/ADM87/ggame/systems"
	"github.com/ADM87/ggame/world"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640 * 0.6
	ScreenHeight = 480 * 0.6
)

var (
	backgroundColor = color.RGBA{100, 149, 237, 255}

	tile0000Renderer rendering.ISpriteRenderer
	tile0010Renderer rendering.ISpriteRenderer
)

type Game interface {
	Start() error
}

type gameshell struct {
	systems map[ecs.ECSTypeID]ecs.ISystem

	player components.IActor
	block  components.IActor
	camera components.ICamera
}

func NewGame() Game {
	tile0000Renderer = rendering.NewSpriteRenderer()
	tile0000Renderer.SetImage(resources.LoadImage("tile_0000"))
	tile0000Renderer.SetAnchor(0.5, 1)

	tile0010Renderer = rendering.NewSpriteRenderer()
	tile0010Renderer.SetImage(resources.LoadImage("tile_0010"))
	tile0010Renderer.SetAnchor(0.5, 0.5)
	return &gameshell{
		systems: make(map[ecs.ECSTypeID]ecs.ISystem),

		player: NewActor(),
		block:  NewActor(),
		camera: NewCamera(),
	}
}

func (g *gameshell) Start() error {
	g.RegisterSystems()
	g.RegisterKeys()

	g.player.RenderComponent().SetRenderer(tile0000Renderer)
	g.block.RenderComponent().SetRenderer(tile0010Renderer)

	g.block.Transform().SetPosition(0, 10)
	g.player.Transform().SetScale(2, 2)

	g.player.GetEntity().AddChild(g.block.GetEntity())

	world.AddEntity(g.player.GetEntity())
	world.AddEntity(g.camera.GetEntity())

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

	keyboard.RegisterKey(ebiten.KeyW, keyboard.KeyPhaseHeld, func() error {
		x, y := g.camera.GetPosition()
		g.camera.SetPosition(x, y-1)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyS, keyboard.KeyPhaseHeld, func() error {
		x, y := g.camera.GetPosition()
		g.camera.SetPosition(x, y+1)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyA, keyboard.KeyPhaseHeld, func() error {
		x, y := g.camera.GetPosition()
		g.camera.SetPosition(x-1, y)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyD, keyboard.KeyPhaseHeld, func() error {
		x, y := g.camera.GetPosition()
		g.camera.SetPosition(x+1, y)
		return nil
	})

	keyboard.RegisterKey(ebiten.KeyMinus, keyboard.KeyPhaseHeld, func() error {
		zoom := g.camera.GetZoom() + 0.01
		g.camera.SetZoom(zoom)
		return nil
	})
	keyboard.RegisterKey(ebiten.KeyEqual, keyboard.KeyPhaseHeld, func() error {
		zoom := g.camera.GetZoom() - 0.01
		g.camera.SetZoom(zoom)
		return nil
	})
}

func (g *gameshell) RegisterSystems() {
	g.systems[systems.RenderSystemTypeID] = systems.NewRenderSystem()
}
