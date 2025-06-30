package objects

import (
	"github.com/ADM87/ggame/src/components"
	"github.com/ADM87/ggame/src/sys/types"
)

// Actor defines an interface for a visual entity in the game with an update cycle
type Actor interface {
	types.Disposable // Actor extends the Disposable interface for resource management
	Entity           // Actor extends the Entity interface

	Renderer() components.SpriteRenderer            // Renderer returns the SpriteRenderer interface for rendering capabilities
	SetRenderer(renderer components.SpriteRenderer) // SetRenderer sets the SpriteRenderer for the actor
}

type actor struct {
	Entity

	renderer components.SpriteRenderer
}

func NewActor() Actor {
	return &actor{
		Entity:   NewEntity(),
		renderer: nil,
	}
}

func (a *actor) Renderer() components.SpriteRenderer {
	return a.renderer
}

func (a *actor) SetRenderer(renderer components.SpriteRenderer) {
	a.renderer = renderer
}

func (a *actor) Dispose() error {
	if a.renderer == nil {
		return nil
	}
	return a.renderer.Dispose()
}
