package components

import (
	"github.com/ADM87/ggame/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MinZoom = 0.01
	MaxZoom = 3.0
)

type cameraComponent struct {
	ITransformComponent

	viewMatrix ebiten.GeoM
	viewDirty  bool

	zoom float64 // Current zoom level of the camera
}

func NewCameraComponent(screenWidth, screenHeight int) ICamera {
	return NewCameraComponentWithOwner(ecs.NewEntity(), screenWidth, screenHeight)
}

func NewCameraComponentWithOwner(owner ecs.IEntity, screenWidth, screenHeight int) ICamera {
	return &cameraComponent{
		ITransformComponent: &transformComponent{
			IComponent:  ecs.NewComponentWithOwner(CameraTypeID, owner),
			x:           0,
			y:           0,
			ox:          float64(screenWidth) * 0.5,  // Center the camera
			oy:          float64(screenHeight) * 0.5, // Center the camera
			sx:          1,
			sy:          1,
			rot:         0,
			deg:         0,
			localMatrix: ebiten.GeoM{},
			localDirty:  true,
			worldMatrix: ebiten.GeoM{},
			worldDirty:  true,
		},
		viewMatrix: ebiten.GeoM{},
		viewDirty:  true,
		zoom:       1.0,
	}
}

func (c *cameraComponent) GetZoom() float64 {
	return c.zoom
}

func (c *cameraComponent) SetZoom(zoom float64) {
	if zoom < MinZoom {
		zoom = MinZoom
	} else if zoom > MaxZoom {
		zoom = MaxZoom
	}
	if c.zoom != zoom {
		c.zoom = zoom
		c.ITransformComponent.SetScale(zoom, zoom)
		c.internal_SetDirty(true, true)
	}
}

func (c *cameraComponent) ViewMatrix() ebiten.GeoM {
	if c.viewDirty {
		c.viewMatrix.Reset()
		c.viewMatrix.Concat(c.WorldMatrix())
		c.viewMatrix.Invert()
		c.viewDirty = false
	}
	return c.viewMatrix
}

func (c *cameraComponent) SetPosition(x, y float64) {
	c.ITransformComponent.SetPosition(x, y)
	c.internal_SetDirty(true, true)
}

func (c *cameraComponent) SetRotation(degrees float64) {
	c.ITransformComponent.SetRotation(degrees)
	c.internal_SetDirty(true, true)
}

func (c *cameraComponent) SetScale(x, y float64) {
	c.ITransformComponent.SetScale(x, y)
	c.internal_SetDirty(true, true)
}

func (c *cameraComponent) internal_SetDirty(local, world bool) {
	c.ITransformComponent.internal_SetDirty(local, world)
	c.viewDirty = true
}
