package objects

import (
	"github.com/ADM87/ggame/components"
	"github.com/hajimehoshi/ebiten/v2"
)

// Actor defines an interface for a renderable entity.
//
// Renderers assigned to the actor are not disposed of when the actor is disposed, as they are not the owners of the renderer but are just assigned to it.
type Actor interface {
	Entity // Actor extends the Entity interface

	Renderer() components.Renderer            // Renderer returns the assigned Renderer of the actor
	SetRenderer(renderer components.Renderer) // SetRenderer assigns a Renderer to the actor
}

type actor struct {
	Entity

	renderer components.Renderer
}

func NewActor() Actor {
	return &actor{
		Entity:   NewEntity(),
		renderer: nil,
	}
}

func (a *actor) Renderer() components.Renderer {
	return a.renderer
}

func (a *actor) SetRenderer(renderer components.Renderer) {
	a.renderer = renderer
}

func (a *actor) Render(target *ebiten.Image, viewMatrix ebiten.GeoM, transform ebiten.GeoM, op *ebiten.DrawImageOptions) {
	if a.renderer != nil {
		a.renderer.Render(target, viewMatrix, a.Matrix(), op)
	}
}

func (a *actor) Dispose() {
	a.Entity.Dispose()
	a.renderer = nil
}
