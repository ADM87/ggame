package camera

import (
	"github.com/ADM87/ggame/src/components"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera interface {
	components.Transform // Transform embeds the Transform interface to provide position and transformation capabilities

	MoveTo(x, y float64)   // MoveTo sets the camera's position to the specified coordinates
	MoveBy(dx, dy float64) // MoveBy moves the camera by the specified deltas in x and y directions

	GetViewMatrix() ebiten.GeoM // GetViewMatrix retrieves the current view matrix for rendering
}

type cam struct {
	components.Transform

	viewOriginX float64
	viewOriginY float64
	viewWidth   int
	viewHeight  int
	viewMatrix  ebiten.GeoM

	isDirty bool
}

func NewCamera(x, y float64, viewWidth, viewHeight int) Camera {
	c := &cam{
		Transform:   components.NewTransform(x, y),
		viewWidth:   viewWidth,
		viewHeight:  viewHeight,
		viewOriginX: float64(viewWidth) / 2,
		viewOriginY: float64(viewHeight) / 2,
	}
	return c
}

func (c *cam) MoveTo(x, y float64) {
	c.SetPosition(x, y)
}

func (c *cam) MoveBy(dx, dy float64) {
	x, y := c.Position()
	c.SetPosition(x+dx, y+dy)
}

func (c *cam) GetViewMatrix() ebiten.GeoM {
	if !c.isDirty && !c.Transform.IsDirty() {
		return c.viewMatrix
	}

	c.SetOrigin(c.viewOriginX, c.viewOriginY)

	view := c.Matrix()
	view.Invert()

	c.viewMatrix = view
	c.isDirty = false

	return c.viewMatrix
}

func (c *cam) SetDirty() {
	c.Transform.SetDirty()
	c.isDirty = true
}

func (c *cam) SetPosition(x, y float64) {
	c.Transform.SetPosition(x, y)
	c.SetDirty()
}

func (c *cam) SetRotation(radians float64) {
	c.Transform.SetRotation(radians)
	c.SetDirty()
}

func (c *cam) SetScale(sx, sy float64) {
	c.Transform.SetScale(sx, sy)
	c.SetDirty()
}
