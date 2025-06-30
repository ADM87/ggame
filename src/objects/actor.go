package objects

import (
	"github.com/ADM87/ggame/src/components"
	"github.com/ADM87/ggame/src/sys/types"
)

// Actor defines an interface renderable entity.
//
// Renderers assigned to the actor are not disposed of when the actor is disposed, as they are not the owners of the renderer but are just assigned to it.
type Actor interface {
	types.Disposable // Actor extends the Disposable interface for resource management
	Entity           // Actor extends the Entity interface

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

func (a *actor) Dispose() error {
	a.renderer = nil
	return nil
}
