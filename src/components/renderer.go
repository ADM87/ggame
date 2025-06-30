package components

import (
	"image/color"

	"github.com/ADM87/ggame/src/sys/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Renderer defines the interface for basic rendering components
type Renderer interface {
	types.Disposable // Disposable interface for resource management

	GetColor() color.Color  // GetColor retrieves the current color used for rendering
	SetColor(c color.Color) // SetColor sets the color used for rendering
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

func (r *renderer) Dispose() error {
	return nil
}

// =======================================================================
// Sprite Renderer
// =======================================================================

type SpriteRenderer interface {
	Renderer // SpriteRenderer embeds the Renderer interface to provide rendering capabilities

	GetImage() *ebiten.Image    // GetImage retrieves the image used for rendering
	SetImage(img *ebiten.Image) // SetImage sets the image used for rendering
}

type spriteRenderer struct {
	Renderer // Renderer embeds the Renderer interface to provide rendering capabilities

	image *ebiten.Image // Image used for rendering
}

func NewSpriteRenderer() SpriteRenderer {
	return &spriteRenderer{
		Renderer: NewRenderer(),
		image:    nil,
	}
}

func (sr *spriteRenderer) GetImage() *ebiten.Image {
	return sr.image
}

func (sr *spriteRenderer) SetImage(img *ebiten.Image) {
	sr.image = img
}

func (sr *spriteRenderer) Dispose() error {
	sr.image = nil
	return nil
}
