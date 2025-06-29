package camera

import "github.com/ADM87/ggame/src/components"

type Camera interface {
	components.Transform
}

type cam struct {
	components.Transform

	viewWidth  int // Width of the camera view
	viewHeight int // Height of the camera view
}

// NewCamera creates a new camera component with the specified position
func NewCamera(x, y float64, viewWidth, viewHeight int) Camera {
	return &cam{
		Transform:  components.NewTransform(x, y),
		viewWidth:  viewWidth,
		viewHeight: viewHeight,
	}
}
