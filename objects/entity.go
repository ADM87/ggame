package objects

import (
	"github.com/ADM87/ggame/components"
	"github.com/ADM87/ggame/sys/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Entity defines an interface for a hierarchical entity in the game.
//
// Disposing of the entity will also dispose its children. Any further use of the entity after disposal will result in undefined behavior.
//
// An entity provides a rendering interface that allows it to render any potential renderable children it may have.
type Entity interface {
	types.Disposable // Entity extends the Disposable interface for resource management
	components.Transform

	AddChild(Entity)    // AddChild adds a child entity to the current entity
	RemoveChild(Entity) // RemoveChild removes a child entity from the current entity

	NumChildren() int         // NumChildren returns the number of child entities
	ChildAt(index int) Entity // ChildAt returns the child entity at the specified index

	Render(target *ebiten.Image, viewMatrix ebiten.GeoM, transform ebiten.GeoM, op *ebiten.DrawImageOptions)
}

type entity struct {
	components.Transform
	children []Entity
}

func NewEntity() Entity {
	return &entity{
		Transform: components.NewTransform(0, 0),
		children:  []Entity{},
	}
}

func (e *entity) AddChild(child Entity) {
	if child == nil {
		return // Do not add nil children
	}
	e.children = append(e.children, child)
}

func (e *entity) RemoveChild(child Entity) {
}

func (e *entity) NumChildren() int {
	return len(e.children)
}

func (e *entity) ChildAt(index int) Entity {
	if index < 0 || index >= len(e.children) {
		return nil // Return nil if index is out of bounds
	}
	return e.children[index]
}

func (e *entity) Render(target *ebiten.Image, viewMatrix ebiten.GeoM, transform ebiten.GeoM, op *ebiten.DrawImageOptions) {

}

func (e *entity) Dispose() {
	for _, child := range e.children {
		child.Dispose()
	}
	e.children = nil
}
