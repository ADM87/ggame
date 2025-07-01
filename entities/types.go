package entities

import (
	"github.com/ADM87/ggame/components"
	"github.com/ADM87/ggame/sys/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// ========================================================================
// IEntity Interface
// ========================================================================

// IEntity defines the interface for a game entity that can be transformed, rendered, and managed in a scene graph.
type IEntity interface {
	components.IMovable
	components.IOrigin
	components.IRender
	components.IRotatable
	components.IScalable

	types.IDisposable

	AddChild(child IEntity)    // AddChild adds a child entity to the current entity
	RemoveChild(child IEntity) // RemoveChild removes a child entity from the current entity
	Children() []IEntity       // Children retrieves the list of child entities
	Parent() IEntity           // Parent retrieves the parent entity of the current entity

	LocalMatrix() ebiten.GeoM // LocalMatrix retrieves the local transformation matrix of the entity
	WorldMatrix() ebiten.GeoM // WorldMatrix retrieves the world transformation matrix of the entity

	Renderer() components.IRender            // Renderer retrieves the renderer component of the entity
	SetRenderer(renderer components.IRender) // SetRenderer sets the renderer for the entity

	// internalSetParent sets the parent entity of the current entity
	//
	// This method is intended for internal use and should not be called directly.
	internalSetParent(parent IEntity)
	// internalSetDirty marks the entity as dirty, indicating that its transformation matrix needs to be recalculated
	//
	// This method is intended for internal use and should not be called directly.
	internalSetDirty()
}
