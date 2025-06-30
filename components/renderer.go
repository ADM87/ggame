package components

import (
	"github.com/ADM87/ggame/sys/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Renderer defines the interface for basic rendering components
type Renderer interface {
	types.Disposable // Disposable extends the Disposable interface for resource management

	Render(target *ebiten.Image, viewMatrix ebiten.GeoM, transformMatrix ebiten.GeoM, op *ebiten.DrawImageOptions) // Render draws the component onto the target using the provided matrices and options
}

type SpriteRenderer interface {
	Renderer // SpriteRenderer extends the Renderer interface

	GetImage() *ebiten.Image    // GetImage retrieves the image used for rendering
	SetImage(img *ebiten.Image) // SetImage sets the image used for rendering
}

// =======================================================================
// Renderer Implementation
// =======================================================================

type renderer struct {
}

// NewRenderer creates a basic renderer with a default color.
func NewRenderer() Renderer {
	return &renderer{}
}

func (r *renderer) Render(target *ebiten.Image, viewMatrix ebiten.GeoM, transformMatrix ebiten.GeoM, op *ebiten.DrawImageOptions) {
	// Default implementation does nothing
}

func (r *renderer) Dispose() {
	// Default implementation does nothing
}

// =======================================================================
// Sprite Renderer
// =======================================================================

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

func (sr *spriteRenderer) Render(target *ebiten.Image, viewMatrix ebiten.GeoM, transformMatrix ebiten.GeoM, op *ebiten.DrawImageOptions) {
	if sr.image == nil {
		return
	}

	op.GeoM = transformMatrix
	op.GeoM.Concat(viewMatrix)

	target.DrawImage(sr.image, op)
}

func (sr *spriteRenderer) Dispose() {
	sr.Renderer.Dispose()
	sr.image = nil
}
