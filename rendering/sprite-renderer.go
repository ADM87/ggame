package rendering

import (
	"github.com/ADM87/ggame/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

type spriteRenderer struct {
	IRenderer

	ax, ay float64
	ox, oy float64

	image       *ebiten.Image
	drawOptions *ebiten.DrawImageOptions
}

func NewSpriteRenderer() ISpriteRenderer {
	sr := &spriteRenderer{
		image:       resources.DefaultImage(),
		drawOptions: &ebiten.DrawImageOptions{},
	}
	return sr
}

func (s *spriteRenderer) Render(screen *ebiten.Image, camera ebiten.GeoM, matrix ebiten.GeoM) {
	if s.image == nil {
		return // No image to render
	}

	// Apply camera transformation
	s.drawOptions.GeoM.Reset()
	s.drawOptions.GeoM.Translate(-s.ox, -s.oy)
	s.drawOptions.GeoM.Concat(matrix)
	s.drawOptions.GeoM.Concat(camera)

	// Draw the image with the current options
	screen.DrawImage(s.image, s.drawOptions)
}

func (s *spriteRenderer) GetAnchor() (float64, float64) {
	return s.ax, s.ay
}

func (s *spriteRenderer) SetAnchor(x, y float64) {
	if s.ax != x || s.ay != y {
		s.ax = x
		s.ay = y

		w, t := s.GetImage().Bounds().Dx(), s.GetImage().Bounds().Dy()

		s.ox = float64(w) * s.ax
		s.oy = float64(t) * s.ay
	}
}

func (s *spriteRenderer) SetImage(image *ebiten.Image) {
	if image == nil {
		image = resources.DefaultImage() // Use default image if nil
	}
	s.image = image
}

func (s *spriteRenderer) GetImage() *ebiten.Image {
	return s.image
}

func (s *spriteRenderer) SetBlendMode(mode ebiten.Blend) {
	s.drawOptions.Blend = mode
}

func (s *spriteRenderer) SetColorScale(cs ebiten.ColorScale) {
	s.drawOptions.ColorScale = cs
}

func (s *spriteRenderer) SetFilter(filter ebiten.Filter) {
	s.drawOptions.Filter = filter
}

func (s *spriteRenderer) SetMipMaps(toggle bool) {
	s.drawOptions.DisableMipmaps = !toggle
}

func (s *spriteRenderer) Dispose() {
	s.image = nil
	s.drawOptions = nil
}
