package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer interface {
	GetColor() color.Color
	SetColor(c color.Color)

	Draw(renderTarget *ebiten.Image)
}

type renderer struct {
	color color.Color
}

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

func (r *renderer) Draw(renderTarget *ebiten.Image) {
	renderTarget.Fill(r.color)
}
