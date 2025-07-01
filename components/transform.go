package components

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	minScale = 0.0001 // Minimum scale to prevent division by zero
)

type transform struct {
	x, y    float64     // x and y coordinates of the transform
	ox, oy  float64     // Origin x and y coordinates for the transform
	sx, sy  float64     // Scale factors for the transform
	radians float64     // Rotation in radians for the transform
	degrees float64     // Rotation in degrees for the transform
	matrix  ebiten.GeoM // Transformation matrix for the transform
	isDirty bool        // isDirty indicates if the transform has changed since the last update
}

// NewTransform constructs a new transform component
func NewTransform() ITransform {
	return &transform{
		x:       0,
		y:       0,
		ox:      0,
		oy:      0,
		sx:      1,
		sy:      1,
		radians: 0,
		degrees: 0,
		isDirty: true,
		matrix:  ebiten.GeoM{},
	}
}

// =======================================================================
// ITransform Interface Implementation
// =======================================================================

func (t *transform) Matrix() ebiten.GeoM {
	if t.isDirty {
		t.matrix.Reset()
		t.matrix.Rotate(t.radians)   // Apply rotation
		t.matrix.Scale(t.sx, t.sy)   // Apply scale
		t.matrix.Translate(t.x, t.y) // Translate to position first
		t.isDirty = false
	}
	return t.matrix
}

func (t *transform) SetDirty() {
	t.isDirty = true
}

func (t *transform) IsDirty() bool {
	return t.isDirty
}

// =======================================================================
// IMovable Interface Implementation
// =======================================================================

func (t *transform) Position() (x, y float64) {
	return t.x, t.y
}

func (t *transform) SetPosition(x, y float64) {
	if t.x == x && t.y == y {
		return
	}
	t.x, t.y = x, y
	t.SetDirty()
}

// =======================================================================
// IOrigin Interface Implementation
// =======================================================================

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

// =======================================================================
// IRotatable Interface Implementation
// =======================================================================

func (t *transform) Rotation() float64 {
	return t.degrees
}

func (t *transform) SetRotation(degrees float64) {
	wrapped := math.Mod(degrees, 360.0)
	if wrapped < 0 {
		wrapped += 360.0
	}
	if t.degrees == wrapped {
		return
	}
	t.degrees = wrapped
	t.radians = wrapped * 0.0174532925199 // Convert degrees to radians
	t.SetDirty()
}

// =======================================================================
// IScalable Interface Implementation
// =======================================================================

func (t *transform) Scale() (sx, sy float64) {
	return t.sx, t.sy
}

func (t *transform) SetScale(sx, sy float64) {
	if sx < minScale {
		sx = minScale
	}
	if sy < minScale {
		sy = minScale
	}
	if t.sx == sx && t.sy == sy {
		return
	}
	t.sx, t.sy = sx, sy
	t.SetDirty()
}
