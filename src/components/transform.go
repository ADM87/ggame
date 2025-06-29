package components

import "github.com/hajimehoshi/ebiten/v2"

type Transform interface {
	Position() (x, y float64) // Position returns the current position of the transform
	SetPosition(x, y float64) // SetPosition sets the position of the transform

	Matrix() ebiten.GeoM // Matrix returns the transformation matrix for the transform
}

type transform struct {
	x, y    float64     // x and y coordinates of the transform
	matrix  ebiten.GeoM // Transformation matrix for the transform
	isDirty bool        // isDirty indicates if the transform has changed since the last update
}

// NewTransform creates a new transform component with the specified position
func NewTransform(x, y float64) Transform {
	return &transform{
		x:      x,
		y:      y,
		matrix: ebiten.GeoM{},
	}
}

func (t *transform) Position() (x, y float64) {
	return t.x, t.y
}

func (t *transform) SetPosition(x, y float64) {
	t.x = x
	t.y = y
	t.setDirty()
}

func (t *transform) Matrix() ebiten.GeoM {
	if !t.isDirty {
		return t.matrix
	}

	t.matrix.Reset()
	t.matrix.Translate(t.x, t.y)
	t.isDirty = false

	copy := t.matrix
	return copy
}

func (t *transform) setDirty() {
	t.isDirty = true
}
