package systems

import (
	"github.com/ADM87/ggame/components"
	"github.com/ADM87/ggame/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type renderSystem struct {
	ecs.ISystem

	renderables []components.IRenderComponent
	rootMatrix  ebiten.GeoM
}

func NewRenderSystem() IRenderSystem {
	return &renderSystem{
		ISystem:     ecs.NewSystem(ecs.GetTypeID[IRenderSystem](), components.RenderTypeID),
		renderables: make([]components.IRenderComponent, 0),
		rootMatrix:  ebiten.GeoM{},
	}
}

func (s *renderSystem) Render(target *ebiten.Image, view ebiten.GeoM) {
	for _, renderable := range s.renderables {
		renderable.Render(target, view, s.rootMatrix)
	}
}

func (s *renderSystem) Update(entities []ecs.IEntity, dt float64) {
	s.renderables = s.renderables[:0]
	for _, entity := range entities {
		if renderable, ok := entity.GetComponent(components.RenderTypeID).(components.IRenderComponent); ok {
			s.renderables = append(s.renderables, renderable)
		}
	}
}

func (s *renderSystem) Dispose() {
	s.renderables = nil
	s.rootMatrix.Reset()
	s.ISystem.Dispose()
}
