package components

import (
	"github.com/ADM87/ggame/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

type spriteRenderer struct {
	image       *ebiten.Image
	drawOptions *ebiten.DrawImageOptions
}

func NewSpriteRenderer() ISpriteRenderer {
	return &spriteRenderer{
		image:       resources.DefaultImage(),
		drawOptions: &ebiten.DrawImageOptions{},
	}
}

// =======================================================================
// ISpriteRenderer Implementation
// =======================================================================

func (sr *spriteRenderer) GetImage() *ebiten.Image {
	return sr.image
}

func (sr *spriteRenderer) SetImage(img *ebiten.Image) {
	sr.image = img
}

func (sr *spriteRenderer) SetBlendMode(mode ebiten.Blend) {
	sr.drawOptions.Blend = mode
}

func (sr *spriteRenderer) SetColorScale(cs ebiten.ColorScale) {
	sr.drawOptions.ColorScale = cs
}

func (sr *spriteRenderer) SetFilter(filter ebiten.Filter) {
	sr.drawOptions.Filter = filter
}

func (sr *spriteRenderer) SetMipMaps(toggle bool) {
	sr.drawOptions.DisableMipmaps = !toggle
}

func (sr *spriteRenderer) Render(target *ebiten.Image, finalTransform ebiten.GeoM) {
	if sr.image == nil {
		return // No image to render
	}

	sr.drawOptions.GeoM = finalTransform
	target.DrawImage(sr.image, sr.drawOptions)
}

func (sr *spriteRenderer) Dispose() {
	sr.image = nil
}
