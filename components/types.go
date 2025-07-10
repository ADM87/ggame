package components

import (
	"github.com/ADM87/ggame/ecs"
	"github.com/ADM87/ggame/rendering"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ActorTypeID     = ecs.GetTypeID[IActor]()
	CameraTypeID    = ecs.GetTypeID[ICamera]()
	RenderTypeID    = ecs.GetTypeID[IRenderComponent]()
	TransformTypeID = ecs.GetTypeID[ITransformComponent]()
)

type IActor interface {
	ecs.IComponent

	Transform() ITransformComponent    // Returns the transform component of the actor
	RenderComponent() IRenderComponent // Returns the render component of the actor
}

type ICamera interface {
	ITransformComponent

	GetZoom() float64     // Returns the zoom level of the camera
	SetZoom(zoom float64) // Sets the zoom level of the camera

	ViewMatrix() ebiten.GeoM // Returns the view matrix of the camera, which is used for rendering
}

type ITransformComponent interface {
	ecs.IComponent

	GetPosition() (x, y float64) // Returns the position of the entity
	SetPosition(x, y float64)    // Sets the position of the entity

	GetRotation() float64        // Returns the rotation of the entity in radians
	SetRotation(degrees float64) // Sets the rotation of the entity in degrees

	GetScale() (x, y float64) // Returns the scale of the entity
	SetScale(x, y float64)    // Sets the scale of the entity

	LocalMatrix() ebiten.GeoM // Returns the local transformation matrix of the entity
	WorldMatrix() ebiten.GeoM // Returns the world transformation matrix of the entity, including parent transforms

	internal_SetDirty(local, world bool) // Internal method to mark the component as dirty, indicating that its transformation needs to be recalculated
	internal_CacheParentTransform()      // Internal method to cache the parent transform component
}

type IRenderComponent interface {
	ecs.IComponent

	SetRenderer(renderer rendering.IRenderer) // Sets the renderer for this component
	GetRenderer() rendering.IRenderer         // Retrieves the current renderer for this component

	IsVisible() bool         // Returns whether the component should be rendered
	SetVisible(visible bool) // Sets the visibility of the component

	Render(screen *ebiten.Image, camera ebiten.GeoM, rootMatrix ebiten.GeoM) // Renders the component to the screen using the provided camera transformation
}
