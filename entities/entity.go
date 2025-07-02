package entities

import (
	"slices"

	"github.com/ADM87/ggame/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type entity struct {
	components.ITransform
	components.IRenderer
	components.IUpdatable

	children []IEntity // List of child entities
	parent   IEntity   // Parent entity

	worldMatrix      ebiten.GeoM // World transformation matrix
	worldMatrixDirty bool        // Flag to indicate if the world matrix is dirty
}

// NewEntity creates a new IEntity instance with default components and properties.
//
// Disposing of an entity will also dispose of its children. Referening a disposed entity will result in unexpected behavior.
func NewEntity() IEntity {
	return &entity{
		ITransform:       components.NewTransform(),
		IRenderer:        nil, // Renderer can be set later
		children:         make([]IEntity, 0),
		parent:           nil,
		worldMatrix:      ebiten.GeoM{},
		worldMatrixDirty: true,
	}
}

func (e *entity) SetPosition(x, y float64) {
	e.ITransform.SetPosition(x, y)
	e.internalSetDirty()
}

func (e *entity) SetScale(x, y float64) {
	e.ITransform.SetScale(x, y)
	e.internalSetDirty()
}

func (e *entity) SetRotation(degrees float64) {
	e.ITransform.SetRotation(degrees)
	e.internalSetDirty()
}

func (e *entity) SetDirty() {
	e.internalSetDirty()
}

// ========================================================================
// IEntity Implementation
// ========================================================================

func (e *entity) AddChild(child IEntity) {
	if child == nil {
		return
	}

	if i := slices.Index(e.children, child); i >= 0 {
		return
	}

	if parent := child.Parent(); parent != nil {
		if parent == e {
			panic("Child's parent is already set to this entity, but missing from the children list")
		}
		parent.RemoveChild(child)
	}

	child.internalSetParent(e)

	e.children = append(e.children, child)
}

func (e *entity) RemoveChild(child IEntity) {
	if child == nil {
		return
	}

	for i, c := range e.children {
		if c == child {
			e.children = append(e.children[:i], e.children[i+1:]...)
			child.internalSetParent(nil)
			return
		}
	}
}

func (e *entity) Children() []IEntity {
	return e.children
}

func (e *entity) Parent() IEntity {
	return e.parent
}

func (e *entity) LocalMatrix() ebiten.GeoM {
	return e.Matrix()
}

func (e *entity) WorldMatrix() ebiten.GeoM {
	if e.worldMatrixDirty {
		if e.parent != nil {
			e.worldMatrix = e.parent.WorldMatrix()
			e.worldMatrix.Concat(e.Matrix())
		} else {
			e.worldMatrix = e.Matrix()
		}
		e.worldMatrixDirty = false
	}
	return e.worldMatrix
}

func (e *entity) Renderer() components.IRenderer {
	return e.IRenderer
}

func (e *entity) SetRenderer(renderer components.IRenderer) {
	e.IRenderer = renderer
}

func (e *entity) internalSetParent(parent IEntity) {
	e.parent = parent
	e.internalSetDirty()
}

func (e *entity) internalSetDirty() {
	e.ITransform.SetDirty()
	e.worldMatrixDirty = true

	for _, child := range e.Children() {
		child.internalSetDirty()
	}
}

// ========================================================================
// IDisposable Implementation
// ========================================================================

func (e *entity) Dispose() {
	for _, child := range e.Children() {
		child.Dispose()
	}

	e.children = nil
	e.parent = nil

	e.ITransform = nil
	e.IRenderer = nil
}

// ========================================================================
// IRender Implementation
// ========================================================================

func (e *entity) Render(target *ebiten.Image, view ebiten.GeoM, matrix ebiten.GeoM) {
	if len(e.Children()) == 0 && e.IRenderer == nil {
		return
	}

	transformMatrix := e.LocalMatrix()
	transformMatrix.Concat(matrix)

	if e.IRenderer != nil {
		e.IRenderer.Render(target, view, transformMatrix)
	}

	for _, child := range e.Children() {
		child.Render(target, view, transformMatrix)
	}
}

// ========================================================================
// IUpdatable Implementation
// ========================================================================

func (e *entity) Update(dt float64) {

}
