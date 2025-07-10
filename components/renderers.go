package components

import (
	"github.com/ADM87/ggame/ecs"
	"github.com/ADM87/ggame/rendering"
	"github.com/hajimehoshi/ebiten/v2"
)

type spriteRendererComponent struct {
	ecs.IComponent

	renderer rendering.IRenderer
	visible  bool
}

func NewSpriteRenderComponent() IRenderComponent {
	return NewSpriteRenderComponentWithOwner(ecs.NewEntity())
}

func NewSpriteRenderComponentWithOwner(owner ecs.IEntity) IRenderComponent {
	return &spriteRendererComponent{
		IComponent: ecs.NewComponentWithOwner(RenderTypeID, owner),
		visible:    true,
	}
}

func (s *spriteRendererComponent) SetRenderer(renderer rendering.IRenderer) {
	s.renderer = renderer
}

func (s *spriteRendererComponent) GetRenderer() rendering.IRenderer {
	return s.renderer
}

func (s *spriteRendererComponent) IsVisible() bool {
	return s.visible
}

func (s *spriteRendererComponent) SetVisible(visible bool) {
	s.visible = visible
}

func (s *spriteRendererComponent) Render(screen *ebiten.Image, camera ebiten.GeoM, rootMatrix ebiten.GeoM) {
	if !s.visible {
		return
	}
	s.renderHierarchy(s.GetEntity(), screen, camera, rootMatrix)
}

func (s *spriteRendererComponent) renderHierarchy(entity ecs.IEntity, screen *ebiten.Image, camera ebiten.GeoM, rootMatrix ebiten.GeoM) {
	if transform, ok := entity.GetComponent(TransformTypeID).(ITransformComponent); ok {
		matrix := transform.LocalMatrix()
		matrix.Concat(rootMatrix)

		s.renderer.Render(screen, camera, matrix)

		for _, child := range entity.Children() {
			if renderer, ok := child.GetComponent(RenderTypeID).(IRenderComponent); ok {
				renderer.Render(screen, camera, matrix)
			}
		}
	}
}
