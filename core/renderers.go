package core

import (
	"github.com/ADM87/ggame/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

type spriteRenderer struct {
	ax, ay      float64
	ox, oy      float64
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
// IAnchor Implementation
// =======================================================================

func (sr *spriteRenderer) Anchor() (ax, ay float64) {
	return sr.ax, sr.ay
}

func (sr *spriteRenderer) SetAnchor(ax, ay float64) {
	sr.ax = ax
	sr.ay = ay

	w, h := sr.image.Bounds().Dx(), sr.image.Bounds().Dy()

	sr.ox = float64(w) * ax
	sr.oy = float64(h) * ay
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

// =======================================================================
// IRenderer Implementation
// =======================================================================

func (sr *spriteRenderer) Render(target *ebiten.Image, view ebiten.GeoM, matrix ebiten.GeoM) {
	if sr.image == nil {
		return
	}

	sr.drawOptions.GeoM.Reset()
	sr.drawOptions.GeoM.Translate(-sr.ox, -sr.oy)
	sr.drawOptions.GeoM.Concat(matrix)
	sr.drawOptions.GeoM.Concat(view)

	target.DrawImage(sr.image, sr.drawOptions)
}

func (sr *spriteRenderer) Dispose() {
	sr.image = nil
}
