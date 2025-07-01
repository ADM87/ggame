package components

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ========================================================================
// Rendering Interfaces
// ========================================================================

// IRender defines a basic interface for a rendering components
type IRender interface {
	Render(target *ebiten.Image, view ebiten.GeoM, matrix ebiten.GeoM) // Render draws the component onto the target image
}

// ISpriteRenderer defines the interface for a sprite rendering component
type ISpriteRenderer interface {
	IRender

	GetImage() *ebiten.Image    // GetImage retrieves the image used for rendering
	SetImage(img *ebiten.Image) // SetImage sets the image used for rendering

	SetBlendMode(mode ebiten.Blend)     // SetBlendMode sets the blend mode for rendering
	SetColorScale(cs ebiten.ColorScale) // SetColorScale sets the color scale for rendering
	SetFilter(filter ebiten.Filter)     // SetFilter sets the filter for rendering
	SetMipMaps(toggle bool)             // UseMipMaps enables or disables mipmaps for the renderer
}

// =======================================================================
// Transformation Interfaces
// =======================================================================

// IMovable defines the interface for a component that can be moved in the game world.
type IMovable interface {
	Position() (x, y float64) // Position retrieves the current position of the movable component
	SetPosition(x, y float64) // SetPosition sets the position to the specified coordinates
}

// IOrigin
type IOrigin interface {
	Origin() (ox, oy float64) // Origin retrieves the current origin point
	SetOrigin(ox, oy float64) // SetOrigin sets the origin point
}

// IRotatable defines the interface for a component that can be rotated in the game world.
type IRotatable interface {
	Rotation() float64           // Rotation retrieves the current rotation in radians
	SetRotation(radians float64) // SetRotation sets the rotation in radians
}

// IScalable defines the interface for a component that can be scaled in the game world.
type IScalable interface {
	Scale() (sx, sy float64) // Scale retrieves the current scale factors in x
	SetScale(sx, sy float64) // SetScale sets the scale factors in x and y directions
}

// ITransform defines the interface for a transform component that can be to manipulate position, rotation, scale, and origin.
type ITransform interface {
	IMovable   // Embedding Movable interface for position manipulation
	IOrigin    // Embedding Origin interface for origin manipulation
	IRotatable // Embedding Rotatable interface for rotation manipulation
	IScalable  // Embedding Scalable interface for scaling manipulation

	Matrix() ebiten.GeoM // Matrix returns the transformation matrix for the transform

	IsDirty() bool // IsDirty checks if the transform is dirty, meaning it has changed since the last update
	SetDirty()     // SetDirty marks the transform as dirty, indicating it has changed
}
