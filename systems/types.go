package systems

import (
	"github.com/ADM87/ggame/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	RenderSystemTypeID = ecs.GetTypeID[IRenderSystem]()
)

type IRenderSystem interface {
	ecs.ISystem

	Render(target *ebiten.Image, view ebiten.GeoM)
}
