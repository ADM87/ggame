package components

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ========================================================================
// Logic Interfaces
// ========================================================================

type UpdateFunc[T IUpdatable] func(dt float64) // UpdateFunc defines a function type for updating components

type IUpdatable interface {
	Update(dt float64) // Update performs the logic update for the component
}

type IUpdater[T IUpdatable] interface {
	IUpdatable

	Add(updateFunc UpdateFunc[T])    // Add adds a new component to the updater
	Remove(updateFunc UpdateFunc[T]) // Remove removes a component from the updater
}

// ========================================================================
// Rendering Interfaces
// ========================================================================

// IRenderer defines a basic interface for a rendering components
type IRenderer interface {
	Render(target *ebiten.Image, view ebiten.GeoM, matrix ebiten.GeoM) // Render draws the component onto the target image
}

// ISpriteRenderer defines the interface for a sprite rendering component
type ISpriteRenderer interface {
	IAnchor
	IRenderer

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

type IAnchor interface {
	Anchor() (ax, ay float64) // Anchor retrieves the current anchor point of the component
	SetAnchor(ax, ay float64) // SetAnchor sets the anchor point to the specified coordinates
}

// IMovable defines the interface for a component that can be moved in the game world.
type IMovable interface {
	Position() (x, y float64) // Position retrieves the current position of the movable component
	SetPosition(x, y float64) // SetPosition sets the position to the specified coordinates
}

// IRotatable defines the interface for a component that can be rotated in the game world.
type IRotatable interface {
	Rotation() float64           // Rotation retrieves the current rotation in degrees
	SetRotation(degrees float64) // SetRotation sets the rotation in degrees
}

// IScalable defines the interface for a component that can be scaled in the game world.
type IScalable interface {
	Scale() (sx, sy float64) // Scale retrieves the current scale factors in x
	SetScale(sx, sy float64) // SetScale sets the scale factors in x and y directions
}

// ITransform defines the interface for a transform component that can be to manipulate position, rotation, scale, and origin.
type ITransform interface {
	IMovable   // Embedding Movable interface for position manipulation
	IRotatable // Embedding Rotatable interface for rotation manipulation
	IScalable  // Embedding Scalable interface for scaling manipulation

	Matrix() ebiten.GeoM // Matrix returns the transformation matrix for the transform

	IsDirty() bool // IsDirty checks if the transform is dirty, meaning it has changed since the last update
	SetDirty()     // SetDirty marks the transform as dirty, indicating it has changed
}
