package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Renderer defines the interface for basic rendering components
type Renderer interface {
	GetColor() color.Color  // GetColor retrieves the current color used for rendering
	SetColor(c color.Color) // SetColor sets the color used for rendering

	Renderer(buffer *ebiten.Image) // Renderer renders the component to the provided buffer
}

// =======================================================================
// Renderer Implementation
// =======================================================================

type renderer struct {
	color color.Color
}

// NewRenderer creates a basic renderer with a default color.
func NewRenderer() Renderer {
	return &renderer{
		color: color.RGBA{255, 255, 255, 255},
	}
}

func (r *renderer) GetColor() color.Color {
	return r.color
}

func (r *renderer) SetColor(c color.Color) {
	r.color = c
}

func (r *renderer) Renderer(buffer *ebiten.Image) {

}
