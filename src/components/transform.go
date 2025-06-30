package components

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

// Transform defines the interface for a transform component that can be used to manipulate position, rotation, scale, and origin.
type Transform interface {
	Component // Transform embeds the Component interface to provide a unique ID for the transform

	AddChild(child Transform)      // AddChild adds a child transform to the current transform
	RemoveChild(child Transform)   // RemoveChild removes a child transform from the current transform
	Parent() Transform             // Parent returns the parent transform of the current transform
	SetParent(parent Transform)    // SetParent sets the parent transform of the current transform
	Matrix() ebiten.GeoM           // Matrix returns the transformation matrix for the transform
	Position() (x, y float64)      // Position returns the current position of the transform
	SetPosition(x, y float64)      // SetPosition sets the position of the transform
	WorldPosition() (x, y float64) // WorldPosition returns the world position of the transform, taking into account its parent transforms
	Rotation() float64             // Rotation returns the rotation of the transform in radians
	SetRotation(radians float64)   // SetRotation sets the rotation of the transform in radians
	Scale() (sx, sy float64)       // Scale returns the scale of the transform
	SetScale(sx, sy float64)       // SetScale sets the scale of the transform
	Origin() (ox, oy float64)      // Origin returns the origin point of the transform
	SetOrigin(ox, oy float64)      // SetOrigin sets the origin point of the transform
	SetDirty()                     // SetDirty marks the transform as dirty, indicating it has changed
	IsDirty() bool                 // IsDirty checks if the transform is dirty, meaning it has changed since the last update
}

type transform struct {
	Component // Component embeds the Component interface to provide a unique ID for the transform

	x, y     float64     // x and y coordinates of the transform
	ox, oy   float64     // Origin point of the transform
	sx, sy   float64     // Scale factors for the transform
	radians  float64     // Rotation in radians for the transform
	parent   Transform   // Parent transform of the current transform
	children []Transform // List of child transforms
	matrix   ebiten.GeoM // Transformation matrix for the transform
	isDirty  bool        // isDirty indicates if the transform has changed since the last update
}

// NewTransform creates a new transform component with the specified position
func NewTransform(x, y float64) Transform {
	return &transform{
		Component: NewComponent(TransformComponentID), // Assign a unique ID for the transform component
		x:         x,
		y:         y,
		ox:        0,
		oy:        0,
		sx:        1,
		sy:        1,
		radians:   0,
		parent:    nil,
		children:  []Transform{},
		isDirty:   true, // Initially dirty to ensure matrix is recalculated
		matrix:    ebiten.GeoM{},
	}
}

func (t *transform) AddChild(child Transform) {
	if child == t {
		panic("Cannot add a transform as its own child")
	}

	if child == nil {
		panic("Cannot add a nil transform as a child")
	}

	if slices.Index(t.children, child) > -1 {
		return // Child already exists, no need to add again
	}

	t.children = append(t.children, child)
}

func (t *transform) RemoveChild(child Transform) {
	if child == nil {
		return // Nothing to remove
	}

	index := slices.Index(t.children, child)
	if index == -1 {
		return // Child not found, nothing to remove
	}

	t.children = append(t.children[:index], t.children[index+1:]...)
}

func (t *transform) Parent() Transform {
	return t.parent
}

func (t *transform) SetParent(parent Transform) {
	if parent == t {
		panic("Cannot set a transform as its own parent")
	}

	if parent == nil {
		t.parent = nil
		return // No parent to set
	}

	if t.hasAncestryOf(parent) {
		panic("Cannot set a parent that is an ancestor of the transform")
	}

	if t.parent == parent {
		return
	}

	if t.parent != nil {
		t.parent.RemoveChild(t)
	}

	t.parent = parent
	parent.AddChild(t)

	t.SetDirty()
}

func (t *transform) Matrix() ebiten.GeoM {
	if !t.isDirty {
		return t.matrix
	}

	t.matrix.Reset()
	t.matrix.Translate(-t.ox, -t.oy) // Translate to origin
	t.matrix.Scale(t.sx, t.sy)       // Apply scale
	t.matrix.Rotate(t.radians)       // Apply rotation
	t.matrix.Translate(t.x, t.y)     // Translate to position

	if t.parent != nil {
		t.matrix.Concat(t.parent.Matrix())
	}

	t.isDirty = false
	return t.matrix
}

func (t *transform) SetDirty() {
	t.isDirty = true
	for _, child := range t.children {
		child.SetDirty()
	}
}

func (t *transform) IsDirty() bool {
	return t.isDirty
}

func (t *transform) Position() (x, y float64) {
	return t.x, t.y
}

func (t *transform) WorldPosition() (x, y float64) {
	m := t.Matrix()
	return m.Apply(0, 0)
}

func (t *transform) SetPosition(x, y float64) {
	if t.x == x && t.y == y {
		return // No change in position, no need to update
	}
	t.x, t.y = x, y
	t.SetDirty()
}

func (t *transform) Rotation() float64 {
	return t.radians
}

func (t *transform) SetRotation(radians float64) {
	if t.radians == radians {
		return // No change in rotation, no need to update
	}
	t.radians = radians
	t.SetDirty()
}

func (t *transform) Scale() (sx, sy float64) {
	return t.sx, t.sy
}

func (t *transform) SetScale(sx, sy float64) {
	if t.sx == sx && t.sy == sy {
		return // No change in scale, no need to update
	}
	t.sx, t.sy = sx, sy
	t.SetDirty()
}

func (t *transform) Origin() (ox, oy float64) {
	return t.ox, t.oy
}

func (t *transform) SetOrigin(ox, oy float64) {
	if t.ox == ox && t.oy == oy {
		return // No change in origin, no need to update
	}
	t.ox, t.oy = ox, oy
	t.SetDirty()
}

func (t *transform) hasAncestryOf(parent Transform) bool {
	p := t.parent
	for p != nil {
		if p == parent {
			return true
		}
		p = p.Parent()
	}
	return false
}
